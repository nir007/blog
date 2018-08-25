<template>
  <div class="row">
    <div class="col-lg-8">
      <br>
      <div class="text-center">
        <img src="/static/assets/img/bad.gif">
        <h1 class="text-center">This page does not work yet</h1>
      </div>
    </div>
    <div class="col-lg-4">
      <search></search>
      <tags></tags>
    </div>
  </div>
</template>

<script>
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import Search from '../parts/Search.vue'
import Tags from '../parts/Tags.vue'
export default {
  name: 'Founded',
  mixins: [ResponseHandler, AuthHandler],
  data () {
    return {}
  },
  components: {
    'tags': Tags,
    'search': Search
  }
}
</script>
