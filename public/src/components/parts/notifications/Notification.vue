<template>
  <div>
    <div v-for="(item) in notifications" :key="item.id" class="text-center">
      <b-alert show dismissible :variant="item.variant">
        <strong>Message: {{item.text}}</strong>
      </b-alert>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Alarm',
  data () {
    return {
      notifications: []
    }
  },
  methods: {
    hide: function () {
      this.show = false
    },
    hideAfter: function (timeout, index) {
      let self = this
      setTimeout(function () {
        self.notifications[index] = null
      }, timeout)
    },
    parse: function (obj) {
      console.log(obj)
      let errText = ''
      let variant = 'variant' in obj ? obj.variant : 'warning'
      if ('err' in obj && typeof obj.err === 'object' &&
        Object.keys(obj.err).length > 0) {
        for (let i in obj.err) {
          if (typeof obj.err[i] !== 'object') {
            if (!/timeout|variant/.test(i)) {
              errText += '[ ' + i + ': ' + obj.err[i] + ' ] '
            }
          } else {
            this.parse(obj.err[i])
          }
        }
        this.notifications.push({variant: variant, text: errText})
      } else if (obj.err.length > 0) {
        this.notifications.push({variant: variant, text: obj.err})
      }
    }
  },
  created () {
    let self = this
    this.$root.$on('alarm', function (err) {
      self.parse(err)
    })
  }
}
</script>
