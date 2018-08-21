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
export default {
  data () {
    return {
      urls: {
        isLogged: '/aj_is_logged'
      },
      params: ['tag']
    }
  },
  components: {
    'articles': Articles,
    'tags': Tags
  },
  created () {
    this.tag = typeof this.$route.params.tag !== 'undefined'
      ? this.$route.params.tag : ''
  },
  mounted () {
    this.$http.post(this.urls.isLogged)
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        this.$root.$emit('nav_top_rebuild', r.data)
      })
  }
}
</script>
