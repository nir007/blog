<template>
<div class="row">
  <div class="col-lg-8">
    <articles v-bind:tag="tag"></articles>
  </div>
  <div class="col-md-4">
    <div class="card my-4">
      <h5 class="card-header">Search</h5>
      <div class="card-body">
        <div class="input-group">
          <input type="text" class="form-control" placeholder="Search for...">
          <span class="input-group-btn">
          <button class="btn btn-secondary" type="button">Go!</button>
        </span>
        </div>
      </div>
    </div>
    <tags></tags>
  </div>
</div>
</template>

<script>
import Articles from '../parts/Articles.vue'
import Tags from '../parts/Tags.vue'
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
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
    'articles': Articles
  },
  created () {
    this.tag = this.$route.params.tag
  }
}
</script>
