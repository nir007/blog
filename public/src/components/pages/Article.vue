<template>
  <div class="row">
    <div class="col-lg-12">
      <br>
      <div v-if="!notFound && article.published">
        <div class="row">
          <div class="col-md-10">
            <p class="lead">
              by
              <a :href="'#/person/' + author.id">
                {{author.nickName}}
              </a> | <small class="text-muted">{{article.createdAt}}</small>
            </p>
          </div>
          <div class="col-md-2 text-right">
            <p v-if="article.isOwner" class="lead">
              <a :href="'#/edit/article/' + article.id">edit</a>
            </p>
          </div>
        </div>
        <p v-if="Object.keys(article.tags).length > 0">
          tags:
          <span v-for="(tag, key) in article.tags" :key="tag.id">
            &nbsp;
            <a :href="'#/articles/' + key">
             [ {{ tag }} ]
            </a>
          </span>
        </p>
        <h1 class="mt-4">{{article.title}}</h1>
        <hr>
        <div class="text">
          <vue-markdown :source="article.text"></vue-markdown>
        </div>
      </div>
      <div v-if="notPublished" class="text-center">
        <br>
        <h1 class="text-center">This article is not published yet</h1>
      </div>
      <div v-if="notFound" class="text-center">
        <h1 class="text-center">Article not found</h1>
        <img src="/static/assets/img/404-error.jpg" alt="404 not found" >
      </div>
      <div class="row margin-top50px margin-bottom20px">
        <div class="col-md-6 text-left">
          <b-button size="lg" v-if="prev" :href="'#/a/' + prev.f1" variant="outline-success">
            {{prev.f2}}
          </b-button>
        </div>
        <div class="col-md-6 text-right">
          <b-button size="lg" v-if="next" :href="'#/a/' + next.f1" variant="outline-success">
            {{next.f2}}
          </b-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Markdown from 'vue-markdown'
import Prism from 'prismjs'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import AuthHandler from '../mixins/AuthHandler.vue'
import Helper from '../mixins/Helper.vue'
export default {
  name: 'Article',
  mixins: [ResponseHandler, AuthHandler, Helper],
  data () {
    return {
      next: false,
      prev: false,
      article: {
        id: 0,
        title: '',
        text: '',
        tags: {},
        createdAt: '',
        published: false,
        isOwner: false
      },
      author: {
        id: 0,
        avatar: '',
        nickName: ''
      },
      isLogged: false,
      notPublished: false,
      notFound: false,
      urls: {
        getArticle: 'aj_get_article'
      }
    }
  },
  updated () {
    Prism.highlightAll()
  },
  created () {
    this.notFound = false
    this.$http.post(this.urls.getArticle,
      'id=' + this.$route.params.id,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      })
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          if (
            'next_article' in r.data.article &&
            'prev_article' in r.data.article
          ) {
            if (r.data.article.next_article != null &&
              typeof r.data.article.next_article === 'object' &&
              'f1' in r.data.article.next_article
            ) {
              this.next = r.data.article.next_article
            }

            if (r.data.article.prev_article != null &&
              typeof r.data.article.prev_article === 'object' &&
              'f1' in r.data.article.prev_article
            ) {
              this.prev = r.data.article.prev_article
            }
          }

          this.article.id = r.data.article.id
          this.article.title = r.data.article.title
          this.article.text = r.data.article.text
          this.article.tags = r.data.article.tags
          this.article.createdAt = this.helpFormatDateTime(r.data.article.created)
          this.article.published = r.data.article.published
          this.article.isOwner = r.data.article.is_owner
          this.author.id = r.data.author.id
          this.author.avatar = r.data.author.avatar
          this.author.nickName = r.data.author.nick_name
          this.notPublished = !r.data.article.published
        } else {
          this.notFound = true
          this.responseFailHandle(r)
        }
      }, function () {
        this.notFound = true
        this.responseFailHandle({status: 500, data: '500 internal server error'})
      })
  },
  components: {
    'vue-markdown': Markdown
  }
}
</script>
