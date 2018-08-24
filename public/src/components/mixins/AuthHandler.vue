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
        if (r.status === 200) {
          this.$root.$emit('nav_top_rebuild', r.data)
        }
      }, function () {
        this.responseFailHandle({status: 500, data: '500 internal server error'})
      })
  }
}
</script>
