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
        phoneIsValid: false,
        urls: {
          signIn: '/aj_sign_in'
        },
        warnings: []
      }
    },
    methods: {
      onInputPhone: function ({ number, isValid, country }) {
        this.phone = number
        this.country = country != null ? country.name : ''
        this.phoneIsValid = isValid
      },
      clearForm: function () {
        this.uuid = ''
      },
      signIn: function () {
        this.warnings = []
        if (!this.uuid) {
          this.warnings.push('Type uuid')
        }

        if (this.warnings.length > 0) {
          this.$root.$emit('warning', this.warnings)
          return
        }

        this.$http.post(this.urls.signIn,
          'uuid=' + this.uuid,
          {
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
            }
          }
        ).then(function (r) {
          r = JSON.parse(r.bodyText)
          if (r.status === 200) {
            location.href = '#/person/'
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
