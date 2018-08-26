<template>
  <div>
    <b-list-group v-if="!isEmpty" v-for="(item) in articles" :key="item.id">
      <b-list-group-item button>{{ item.title }}</b-list-group-item>
    </b-list-group>
    <div v-if="isEmpty" class="text-center">
      <h3 class="text-center">
        <a href="#/new_article">Crate the first article</a>
      </h3>
    </div>
  </div>
</template>

<script>
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'ArticlesUserList',
  mixins: [ResponseHandler],
  data () {
    return {
      isEmpty: false,
      articles: [],
      urls: {
        getArticles: '/get_published_articles'
      },
      limit: 10
    }
  },
  props: {
    authorId: 'Number'
  },
  mounted () {
    let authorId = typeof this.authorId !== 'undefined'
      ? this.authorId : 0

    this.$http.post(this.urls.getArticles,
      'author_id=' + authorId +
      '&limit=' + this.limit +
      '&offset=' + this.articles.length,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      }
    ).then(function (r) {
      r = JSON.parse(r.bodyText)
      console.log(r)
      if (r.status === 200) {
        for (let i in r.data) {
          this.articles.push(r.data[i])
        }
        this.isEmpty = this.articles.length === 0
      } else {
        this.responseFailHandle(r)
      }
    })
  }
}
</script>
