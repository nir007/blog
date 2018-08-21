<template>
  <b-modal centered
           ref="signin"
           title="Enter you uuid"
           @ok="signIn"
           @shown="clearForm">
    <div class="form-group">
      <div class="row">
        <div class="col-md-2">
          <p class="margin-top5px">uuid:</p>
        </div>
        <div class="col-md-10">
          <input type="text"
                 v-model="uuid"
                 class="form-control"
                 placeholder="000-5ffg...">
        </div>
      </div>
    </div>
  </b-modal>
</template>

<script>
export default {
  name: 'SignIn',
  data () {
    return {
      uuid: '',
      urls: {
        signIn: '/aj_sign_in'
      },
      warnings: []
    }
  },
  methods: {
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
          alert(r.data)
        }
      }, function (e) {
        alert(e.statusText)
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
