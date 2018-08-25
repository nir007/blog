<template>
  <span></span>
</template>

<script>
export default {
  name: 'Logout',
  methods: {
    logout: function () {
      this.$http.get('/logout')
    }
  }
}
</script>
