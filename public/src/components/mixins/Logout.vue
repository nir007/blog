<template>
  <span></span>
</template>

<script>
export default {
  name: 'Logout',
  methods: {
    logout: function () {
      this.$http.get('/logout')
      setTimeout(function () {
        location.href = '#/articles'
      }, 1000)
    }
  }
}
</script>
