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
  name: 'notificationAlarm',
  data () {
    return {
      notifications: []
    }
  },
  methods: {
    hide: function () {
      this.show = false
    },
    hideAfter: function (timeout) {
      const self = this
      setTimeout(function () {
        self.notifications.pop()
      }, timeout)
    },
    parse: function (notification) {
      let errText = ''
      const objToString = function (obj) {
        let result = ''
        for (const i in obj) {
          result += '[ ' + i + ': ' + obj[i] + ' ] '
        }
        return result
      }

      if (typeof notification.err === 'object' &&
        Object.keys(notification.err).length > 0
      ) {
        for (const i in notification.err) {
          if (notification.err[i] != null &&
            typeof notification.err[i] !== 'object'
          ) {
            if (!/timeout|variant/.test(i)) {
              errText += '[ ' + i + ': ' + notification.err[i] + ' ] '
            }
          } else if (notification.err[i] != null) {
            errText += objToString(notification.err[i])
          }
        }
      } else {
        errText = notification.err
      }

      const variant = 'variant' in notification ? notification.variant : 'warning'

      if ('timeout' in notification) {
        this.hideAfter(notification.timeout)
      }

      return { variant: variant, text: errText }
    }
  },
  created () {
    const self = this
    this.$root.$on('alarm', function (notification) {
      self.notifications = []
      const result = self.parse(notification)
      if (result.text) {
        self.notifications.push(result)
      }
    })
  }
}
</script>
