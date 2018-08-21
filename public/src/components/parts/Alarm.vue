<template>
  <div v-if="show" class="alert alert-danger fade in alert-dismissible">
    <a href="#" class="close" data-dismiss="alert" aria-label="close" title="close">×</a>
    <strong>Error!</strong> {{text}}
  </div>
</template>

<script>
export default {
  name: 'Alarm',
  data () {
    return {
      text: '',
      show: true,
      timeout: -1
    }
  },
  created () {
    this.$root.$on('alarm', function (text) {
      this.show = true
      this.text = text
    })
  }
}
</script>
