<template>
  <b-modal centered
           size="lg"
           ref="join"
           title="Hello my young friend!"
           @ok="handleOk"
           @shown="clearForm">
    <div class="form-group">
      <p>Before continue tell us about you!</p>
      <div class="form-group">
        <label>Who are you?</label>
        <input type="text" v-model="person" class="form-control" placeholder="I am ...">
      </div>
      <div class="form-group">
        <label>What is your nickname?</label>
        <input type="text" @input="checkNickName" v-model="nickName" class="form-control" placeholder="nir007">
        <small v-if="nickNameExists" class="text-danger">this nickname already used</small>
      </div>
      <div class="form-group">
        <label>Phone number please</label>
        <vue-tel-input v-model="phone"
                       @onInput="onInputPhone"
                      :preferredCountries="['us', 'gb', 'ua', 'ru']">
        </vue-tel-input>
        <small v-if="phoneExists" class="text-danger">this number already used</small>
      </div>
      <div class="form-group">
        <label>Point out you real face:</label>
        <div class="row">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/confused.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/elephant_run.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/funny_elite_machine_gun.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/funny_n_scary.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/headbang.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/headbangers.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/natsu_run.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/polar_bear.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/ryuk.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/shi__happens.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/yui.gif">
          <img @click="setAvatar" class="avatar" src="/static/assets/img/Smokyng.gif">
          <div class="cl"></div>
        </div>
      </div>
    </div>
  </b-modal>
</template>

<script>
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'Join',
  mixins: [ResponseHandler],
  data () {
    return {
      person: '',
      nickName: '',
      avatar: '',
      phone: '',
      country: '',
      phoneIsValid: false,
      urls: {
        join: '/aj_add_user',
        checkNickName: '/aj_get_check_nickname',
        checkPhoneExists: '/aj_get_check_phone'
      },
      nickNameExists: false,
      phoneExists: false,
      warnings: []
    }
  },
  methods: {
    onInputPhone: function ({ number, isValid, country }) {
      this.phone = number
      this.country = country != null ? country.name : ''
      this.phoneIsValid = isValid
      this.checkPhone()
    },
    checkNickName: function () {
      this.$http.post(this.urls.checkNickName,
        'nickname=' + this.nickName,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          this.nickNameExists = r.data
        }
      })
    },
    checkPhone: function () {
      if (this.phoneIsValid) {
        this.$http.post(this.urls.checkPhoneExists,
          'phone=' + this.phone,
          {
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
            }
          }
        ).then(function (r) {
          r = JSON.parse(r.bodyText)
          if (r.status === 200) {
            this.phoneExists = r.data
          }
        })
      } else {
        this.phoneExists = false
      }
    },
    handleOk: function () {
      this.warnings = []
      if (!this.person) {
        this.warnings.push('Enter who are you')
      }

      if (!this.nickName) {
        this.warnings.push('Enter nickname')
      }

      if (this.nickNameExists) {
        this.warnings.push('Enter the other nickname')
      }

      if (!this.phoneIsValid) {
        this.warnings.push('Enter your phone number')
      }

      if (!this.avatar) {
        this.warnings.push('Point out the avatar')
      }

      if (this.warnings.length > 0) {
        this.$root.$emit('warning', this.warnings)
        return
      }

      this.$http.post(this.urls.join,
        {
          person: this.person,
          nick_name: this.nickName,
          avatar: this.avatar,
          country: this.country,
          phone: this.phone
        },
        {
          headers: {
            'Content-Type': 'application/json; charset=UTF-8'
          }
        }
      )
        .then(function (r) {
          r = JSON.parse(r.bodyText)
          if (r.status === 200) {
            location.href = '#/person/'
          } else {
            this.responseFailHandle({status: r.status, err: r.data})
          }
        }, function () {
          this.responseFailHandle({status: 500, err: 'Internal server error'})
        })
    },
    clearForm: function () {
      this.person = ''
      this.nickName = ''
      this.avatar = ''
      this.phone = ''
      this.nickNameExists = false
      this.phoneExists = false
      this.resetAvatar()
    },
    resetAvatar: function () {
      let imgs = document.getElementsByClassName('avatar')
      for (let i = 0; i < imgs.length; i++) {
        imgs[i].style = 'border-color: forestgreen;' +
          ' border-radius: 900px'
      }
    },
    setAvatar: function (e) {
      this.resetAvatar()
      e.target.style = 'border-color: red; border-radius: 0px'
      let parts = e.target.getAttribute('src').split('/')
      this.avatar = parts[parts.length - 1]
    }
  },
  mounted () {
    var self = this
    this.$root.$on('join', function () {
      self.$refs.join.show()
    })
  }
}
</script>
