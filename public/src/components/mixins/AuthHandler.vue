<template>
  <span></span>
</template>

<script>
import ResponseHandler from './ResponseHandler.vue'
export default {
  name: 'AuthHandler',
  mixins: [ResponseHandler],
  data () {
    return {
      isLoggedUrl: '/aj_is_logged'
    }
  },
  mounted () {
    this.$http.post(this.isLoggedUrl)
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        this.$root.$emit('check_is_logged', r.data)
      })
  }
}
</script>
