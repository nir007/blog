<template>
  <div class="card my-4">
    <h5 class="card-header">Tags</h5>
    <div class="card-body">
      <ul class="list-unstyled mb-0 list-two-params text-center">
        <li v-for="(val, key) in tags" :key="val.id">
          <a :href="'#/articles/' + key">{{val}}</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Tags',
  data () {
    return {
      tags: {},
      urls: {
        getTags: '/aj_get_tags'
      }
    }
  },
  mounted () {
    this.$http.post(this.urls.getTags, '')
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          this.tags = r.data
        }
      }, function (e) {
        alert(e.statusText)
      })
  }
}
</script>
