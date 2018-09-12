<template>
  <b-modal centered
           ref="signin"
           title="Enter you uuid"
           @ok="signIn"
           @shown="clearForm">
    <div class="form-group">
      <label>Phone number please</label>
      <vue-tel-input v-model="phone"
                     @onInput="onInputPhone"
                     :preferredCountries="['us', 'gb', 'ua', 'ru']">
      </vue-tel-input>
      <small v-if="showFailNumber" class="text-danger">this number is invalid</small>
    </div>
    <div class="form-group">
      <label>Enter code from sms</label>
      <input type="text" maxlength="10"/>
      <small v-if="showFailNumber" class="text-danger">this number is invalid</small>
    </div>
  </b-modal>
</template>

<script>
export default {
  name: 'SignIn',
  data () {
    return {
      phone: '',
      country: '',
      showFailNumber: false,
      phoneIsExists: false,
      phoneIsValid: false,
      urls: {
        getCode: '/aj_get_code_to_login',
        login: '/aj_login'
      },
      warnings: []
    }
  },
  methods: {
    onInputPhone ({ number, isValid, country }) {
      this.phone = number
      this.country = country != null ? country.name : ''
      this.phoneIsValid = isValid
    },
    clearForm () {
      this.phone = ''
      this.code = ''
    },
    signIn () {
      if (!this.phoneIsExists) {
        alert('number is not exist')
      } else {
        location.href = '#/person/'
      }
    },
    getCode () {
      this.phoneIsExists = false
      this.warnings = []

      if (!this.phoneIsValid) {
        this.warnings.push('Number is not valid')
      }

      if (this.warnings.length > 0) {
        this.$root.$emit('warning', this.warnings)
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
        if (r.status === 200) {
          this.phoneIsExists = true
          this.signIn()
        } else {
          this.$root.$emit('alarm', {err: r.data, timeout: 5000})
        }
      })
    }
  },
  mounted () {
    var self = this
    this.$root.$on('signin', function () {
      self.$refs.signin.show()
    })
  }
}
</script>
