<template>
  <div class="root">
    <div v-if="isLogged && !notFound" class="row">
      <div class="col-md-12 no-float">
        <br>
        <p>Update article</p>
        <div class="form-group">
          <input v-model="title" type="text" class="form-control" placeholder="title">
        </div>
        <div class="form-group">
          <input v-model="tags" type="text" class="form-control" placeholder="tags divide by comma">
        </div>
        <br>
        <div class="form-group">
          <div class="row">
            <div class="col-md-6 text-right">
              <p>publish after creating:</p>
            </div>
            <div class="col-md-6">
              <switches style="margin-top: 8px" v-model="published"
                        theme="bootstrap"
                        color="info"
                        type-bold="true"
              ></switches>
            </div>
          </div>
        </div>
        <div class="form-group">
          <b-form-textarea v-model="text"
                           placeholder="Enter"
                           :rows="10"
                           :max-rows="15">
          </b-form-textarea>
          <small>
            <a href="#/markdown" target="_blank">markdown help</a>
          </small>
          <hr>
          <vue-markdown :source="text"></vue-markdown>
        </div>
        <hr>
        <div class="form-group">
          <div @click="updateArticle" class="btn btn-outline-secondary btn-block">
            <i class="fa fa-save"></i> Save changes
          </div>
        </div>
      </div>
    </div>
    <div v-if="notFound" class="text-center">
      <h1 class="text-center">Article not found</h1>
      <img src="/static/assets/img/404-error.jpg" alt="404 not found" >
    </div>
    <div v-if="!isLogged && !notFound" class="text-center"><br><br>
      <a href="javascript:void(0)" @click="$root.$emit('join', {})">
        [ join now! ]
      </a> -
      <a href="javascript:void(0)" @click="$root.$emit('signin', {})">
        [ sign in ]
      </a>
      <h2 class="margin-top20px">to continue</h2>
    </div>
  </div>
</template>

<script>
import Switches from 'vue-switches'
import Markdown from 'vue-markdown'
import Prism from 'prismjs'
export default {
  data () {
    return {
      id: 0,
      title: '',
      tags: '',
      text: '',
      authorId: 0,
      published: false,
      notFound: false,
      isLogged: false,
      urls: {
        updateArticle: '/aj_update_article',
        getArticle: 'aj_get_article',
        isLogged: '/aj_is_logged'
      },
      warnings: []
    }
  },
  components: {
    'switches': Switches,
    'vue-markdown': Markdown
  },
  methods: {
    updateArticle () {
      this.warnings = []
      if (!this.title) {
        this.warnings.push('Type title')
      }
      if (!this.text) {
        this.warnings.push('Type text')
      }
      if (this.warnings.length > 0) {
        this.$root.$emit('warning', this.warnings)
        return
      }

      let mapTags = {}
      let tags = this.tags.split(',')

      for (let i in tags) {
        let key = tags[i].trim()
          .replace(/[#| |$|%|^|&|*|(|)|@|!|?|>|<|/] /g, '_')
        mapTags[key] = tags[i]
      }

      let published = this.published ? 1 : 0

      this.$http.post(this.urls.updateArticle,
        {
          id: this.id,
          author_id: this.authorId,
          title: this.title,
          text: this.text,
          tags: mapTags,
          published: published
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          location.href = '#/a/' + r.data.id
        } else {
          this.$root.$emit('alarm', r.data)
        }
      }, function () {
        this.$root.$emit('alarm', 'Some kind of error happened')
      })
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
        console.log(r)
        if (r.status === 200) {
          this.id = r.data.article.id
          this.authorId = r.data.article.author_id
          this.title = r.data.article.title
          this.text = r.data.article.text
          for (let i in r.data.article.tags) {
            this.tags += r.data.article.tags[i] + ', '
          }
          this.published = r.data.article.published
          if (!r.data.article.is_owner) {
            this.$router.push({path: 'a', params: { id: this.id }})
          }
        } else {
          this.notFound = true
        }
      }, function () {
        this.notFound = true
      })
  },
  mounted () {
    this.$http.post(this.urls.isLogged)
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        this.isLogged = r.data
        this.$root.$emit('nav_top_rebuild', r.data)
      })
  }
}
</script>
