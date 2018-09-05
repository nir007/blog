<template>
  <span></span>
</template>

<script>
export default {
  name: 'ResponseHandler',
  data () {
    return {
      codes: {
        500: 'danger',
        404: 'warning',
        403: 'warning',
        401: 'info'
      },
      defaultTimeout: 30000
    }
  },
  methods: {
    responseFailHandle: function (err) {
      if (typeof err === 'object') {
        if ('data' in err && err.status in this.codes) {
          let timeout = 'timeout' in err ? err.timeout : this.defaultTimeout
          this.$root.$emit(
            'alarm',
            {
              err: err.data,
              variant: this.codes[err.status],
              timeout: timeout
            }
          )
        }
      }
    }
  }
}
</script>
