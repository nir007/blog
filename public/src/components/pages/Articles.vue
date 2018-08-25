<template>
<div class="row">
  <div class="col-lg-8">
    <articles v-bind:tag="tag"></articles>
  </div>
  <div class="col-md-4">
    <search></search>
    <tags></tags>
  </div>
</div>
</template>

<script>
import Articles from '../parts/Articles.vue'
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import Search from '../parts/Search.vue'
import Tags from '../parts/Tags.vue'
export default {
  mixins: [ResponseHandler, AuthHandler],
  data () {
    return {
      tag: '',
      params: {
        'tag': 'string'
      }
    }
  },
  components: {
    'tags': Tags,
    'search': Search,
    'articles': Articles
  },
  created () {
    this.tag = this.$route.params.tag
  }
}
</script>
