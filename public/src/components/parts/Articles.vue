<template>
  <div>
    <div v-if="!isEmpty" class="card margin-top22px" v-for="(item) in articles" :key="item.id">
      <div v-if="!isEmpty || (item.isOwner)" v-bind:class="{hiddenBlock: !item.published}" class="card-body">
        <div class="row">
          <div class="col-md-10">
            <h2 class="post__title">
              <a :href="'#/a/' + item.id" class="post__title_link">
                <strong>{{ item.title }}</strong>
              </a>
            </h2>
            <small>
              <span v-for="(tag, key) in item.tags" :key="tag.id">
                <a :href="'#/articles/' + key">
                  {{ tag }}
                </a>
              </span>
            </small>
          </div>
          <div class="col-md-2 text-right">
            <a v-if="item.is_owner" :href="'#/edit/article/' + item.id">
              <strong class="text-info">Edit</strong>
            </a>
          </div>
        </div>
      </div>
    </div>
    <div v-if="isEmpty" class="text-center">
      <img src="/static/assets/img/no-dead-links.jpg" alt="articles not found" >
      <h1 class="text-center">Articles not found</h1>
    </div>
  </div>
</template>

<script>
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  mixins: [ResponseHandler],
  data () {
    return {
      isEmpty: false,
      articles: [],
      urls: {
        getArticles: '/get_articles'
      },
      limit: 10
    }
  },
  props: {
    authorId: 'Number',
    tag: 'String'
  },
  mounted () {
    let authorId = typeof this.authorId !== 'undefined'
      ? this.authorId : 0

    let tag = typeof this.tag !== 'undefined'
      ? this.tag : ''

    this.$http.post(this.urls.getArticles,
      'author_id=' + authorId +
      '&tag=' + tag +
      '&limit=' + this.limit +
      '&offset=' + this.articles.length,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      }
    ).then(function (r) {
      r = JSON.parse(r.bodyText)
      if (r.status === 200) {
        for (let i in r.data) {
          this.articles.push(r.data[i])
        }
        this.isEmpty = this.articles.length === 0
      } else {
        this.responseFailHandle(r)
      }
    }, function () {
      this.responseFailHandle({status: 500, data: '500 internal server error'})
    })
  }
}
</script>
