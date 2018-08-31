<template>
  <span></span>
</template>

<script>
export default {
  name: 'Helper',
  methods: {
    helpFormatDateTime (dateTime) {
      if (dateTime != null && dateTime.length > 20) {
        dateTime = dateTime.split('T')
        if (dateTime.length === 2) {
          let date = dateTime[0].split('-')
          let months = {
            '01': 'Jan.',
            '02': 'Feb.',
            '03': 'Mar.',
            '04': 'Apr.',
            '05': 'May',
            '06': 'June',
            '07': 'July',
            '08': 'Aug.',
            '09': 'Sept.',
            '10': 'Oct.',
            '11': 'Nov.',
            '12': 'Dec.'
          }
          if (date[1] in months) {
            date = date[2] + ' ' + months[date[1]] + ' ' + date[0]
          } else {
            date = (new Date()).getFullYear()
          }

          let time = dateTime[1].split(':')
          time = time[0] + ':' + time[1]

          dateTime = date + ', ' + time
        }
      }

      return dateTime
    }
  }
}
</script>
