<template>
  <div class="margin-top50px">
    <div class="text-center">
      <h1 class="text-center">Page not found</h1>
      <img src="/static/assets/img/404-error.jpg" alt="404 not found" >
    </div>
  </div>
</template>

<script>
import AuthHandler from '../mixins/AuthHandler.vue'
export default {
  name: 'PageNotFound',
  mixins: [AuthHandler]
}
</script>
