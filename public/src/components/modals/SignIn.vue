<template>
  <b-modal centered
           ref="signin"
           title="Sign in"
           hide-footer
           @shown="clearForm">
    <div class="form-group" v-if="!showStep2">
      <label>Phone number please</label>
      <div class="row">
        <div class="col-md-9">
          <vue-tel-input v-model="phone"
                         @onInput="onInputPhone"
                         :preferredCountries="['us', 'gb', 'ua', 'ru']">
          </vue-tel-input>
          <small v-if="showFailNumber" class="text-danger">
            this number is invalid
          </small>
        </div>
        <div class="col-md-3">
          <button type="button" @click="getCode" class="btn btn-primary">
            Get Code
          </button>
        </div>
      </div>
    </div>
    <div class="form-group" v-if="showStep2">
      <label>Enter code from sms</label>
      <div class="row">
        <div class="col-md-9">
          <input type="text"
                 maxlength="10"
                 v-model="code"
                 class="form-control"
          />
          <small v-if="showFailCode" class="text-danger">
            this code is invalid
          </small>
        </div>
        <div class="col-md-3">
          <button type="button" @click="signIn" class="btn btn-primary btn-block">
            Sign In
          </button>
        </div>
      </div>
    </div>
    <div v-if="showStep2">
      <div @click="clearForm" class="btn btn-outline-secondary btn-block">
        Back
      </div>
    </div>
  </b-modal>
</template>

<script>
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'SignIn',
  mixins: [ResponseHandler],
  data () {
    return {
      phone: '',
      code: '',
      country: '',
      showStep2: false,
      showFailNumber: false,
      phoneIsExists: false,
      phoneIsValid: false,
      showFailCode: false,
      urls: {
        getCode: '/aj_get_code_to_login',
        login: '/aj_sign_in'
      },
      warnings: []
    }
  },
  methods: {
    onInputPhone ({ number, isValid, country }) {
      this.phone = number
      this.showFailNumber = false
      this.country = country != null ? country.name : ''
      this.phoneIsValid = isValid
    },
    clearForm () {
      this.phone = ''
      this.code = ''
      this.phoneIsExists = false
      this.showFailNumber = false
      this.showStep2 = false
    },
    signIn () {
      if (!this.phoneIsExists) {
        this.$root.$emit('warning', ['Bad code!'])
      } else if (!this.code) {
        this.showFailCode = true
      } else {
        this.$http.post(this.urls.login,
          'code=' + this.code + '&phone=' + this.phone,
          {
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
            }
          }
        ).then(function (r) {
          r = JSON.parse(r.bodyText)
          if (r.status === 200) {
            location.href = '#/person'
            this.$refs.signin.hide()
          } else {
            this.responseFailHandle({status: r.status, err: r.data})
          }
        })
      }
    },
    getCode () {
      if (!this.phoneIsValid) {
        this.showFailNumber = true
        return
      }

      this.$http.post(this.urls.getCode,
        'phone=' + this.phone,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200 || r.status === 404) {
          this.phoneIsExists = r.status === 200
          this.showStep2 = true
        } else {
          this.$root.$emit('warning', [r.data])
        }
      })
    }
  },
  mounted () {
    const self = this
    this.$root.$on('signin', function () {
      self.$refs.signin.show()
    })
  }
}
</script>
