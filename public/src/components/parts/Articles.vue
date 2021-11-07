<template>
  <div>
    <div v-if="!isEmpty" class="card margin-top22px"
         v-for="(item, index) in articles" :key="item.id"
    >
      <div v-if="!isEmpty || (item.isOwner)"
           v-bind:class="{hiddenBlock: !item.published}"
           class="card-body"
      >
        <div class="row">
          <div class="col-md-10">
            <h2 class="post__title underline-link">
              <a :href="'#/a/' + item.id" class="post__title_link">
                {{ item.title }}
              </a>
            </h2>
          </div>
          <div class="col-md-2 text-right">
            <a v-if="item.is_owner" :href="'#/edit/article/' + item.id">
              <strong class="text-info">Edit</strong>
            </a>
            <p v-if="item.is_owner" @click="initRemoveArticle(item.id, index)">
                <strong class="text-danger pointer">Remove</strong>
            </p>
          </div>
        </div>
        <div class="text preview-text">
          <vue-markdown :source="item.text"></vue-markdown>
        </div>
        <hr>
        <div class="text-center">
          <a :href="'#/a/' + item.id">
            Read more
          </a>
        </div>
      </div>
    </div>
    <b-button
      @click="getArticles"
      variant="outline-dark"
      class="margin-top20px"
      v-if="lastCountGot === limit"
      block
    >More
    </b-button>
    <div v-if="isEmpty && !authorId" class="text-center">
      <img class="margin-top50px"
           src="/static/assets/img/no-dead-links.jpg"
           alt="articles not found"
      >
      <h1 class="text-center">Articles not found</h1>
    </div>
    <div v-if="isEmpty && authorId" class="text-center">
      <br><br><br><br>
      <h1 class="text-center">Articles not found</h1>
      <br><br><br><br>
    </div>
  </div>
</template>

<script>
import ResponseHandler from '../mixins/ResponseHandler.vue'
import Markdown from 'vue-markdown'
import Prism from 'prismjs'
export default {
  mixins: [ResponseHandler],
  data () {
    return {
      isEmpty: false,
      articles: [],
      removeArticleId: 0,
      removeArticleIndex: 0,
      lastCountGot: 0,
      urls: {
        get: '/get_articles',
        remove: '/remove_article'
      },
      limit: 5
    }
  },
  props: {
    authorId: 'Number',
    tag: 'String',
    showPublished: 'Number'
  },
  components: {
    'vue-markdown': Markdown
  },
  methods: {
    getArticles () {
      let authorId = typeof this.authorId !== 'undefined'
        ? this.authorId : 0

      let tag = typeof this.tag !== 'undefined'
        ? this.tag : ''

      let showPublished = typeof this.showPublished !== 'undefined'
        ? this.showPublished : 0

      this.$http.post(this.urls.get,
        'author_id=' + authorId +
        '&tag=' + tag +
        '&limit=' + this.limit +
        '&offset=' + this.articles.length +
        '&show_published=' + showPublished,
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
          this.lastCountGot = r.data != null ? r.data.length : 0
        } else {
          this.responseFailHandle(r)
        }
      }, function () {
        this.responseFailHandle({status: 500, data: '500 internal server error'})
      })
    },
    initRemoveArticle (id, index) {
      this.removeArticleId = id
      this.removeArticleIndex = index
      this.$root.$emit('confirm', 'You want to delete this article! Are you sure ?')
    },
    remove () {
      this.$http.post(this.urls.remove,
        'id=' + this.removeArticleId,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          this.articles.splice(this.removeArticleIndex, 1)
        } else {
          this.responseFailHandle(r)
        }
      }, function () {
        this.responseFailHandle({status: 500, data: '500 internal server error'})
      })
    }
  },
  updated () {
    Prism.highlightAll()
  },
  mounted () {
    const self = this
    this.$root.$on('ok', function () {
      self.remove()
    })
    this.getArticles()
  }
}
</script>
