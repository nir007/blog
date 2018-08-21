<template>
  <div v-if="show" class="text-center">
    <b-alert show dismissible variant="danger">
      <strong>Error: {{err}}</strong>
    </b-alert>
  </div>
</template>

<script>
export default {
  name: 'Alarm',
  data () {
    return {
      err: '...',
      show: false
    }
  },
  methods: {
    hide: function () {
      this.show = false
    },
    hideAfter: function (timeout) {
      let self = this
      setTimeout(function () {
        self.hide()
      }, timeout)
    }
  },
  created () {
    let self = this
    this.$root.$on('alarm', function (err) {
      let errText = ''
      let parseErr = function (err) {
        if (err != null && typeof err === 'object') {
          if ('timeout' in err && err.timeout > 0) {
            self.hideAfter(err.timeout)
          }
          for (let i in err) {
            if (typeof err[i] !== 'object') {
              if (i !== 'timeout') {
                errText += '[ ' + i + ': ' + err[i] + ' ] '
              }
            } else {
              parseErr(err[i])
            }
          }
          self.err = errText
        } else {
          self.err = err
        }
      }

      parseErr(err)
      self.show = true
    })
  }
}
</script>
