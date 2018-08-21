<template>
  <div class="root">
    <div v-if="isLogged" class="row">
      <div class="col-md-12 no-float">
      <br>
      <p>New article</p>
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
        <div class="form-group" style="height: 60%; margin-top: 30px">
          <hr>
          <div @click="addArticle" class="btn btn-outline-secondary btn-block">
              <i class="fa fa-save"></i> Save
          </div>
        </div>
      </div>
    </div>
    <div v-if="!isLogged" class="text-center"><br><br>
      <a href="javascript:void(0)" @click="$root.$emit('join', {})">
        [ join now! ]
      </a> -
      <a href="javascript:void(0)" @click="$root.$emit('signin', {})">
        [ sign in ]
      </a>
      <br><br>
      <h2>to continue</h2>
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
      title: '',
      tags: '',
      published: false,
      text: '',
      isLogged: false,
      urls: {
        addArticle: '/aj_add_article',
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
    addArticle () {
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
        let key = tags[i].trim().replace(/[#| |$|%|^|&|*|(|)|@|!|?|>|<|/] /g, '_')
        mapTags[key] = tags[i]
      }

      let published = this.published ? 1 : 0

      this.$http.post(this.urls.addArticle,
        {
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
          alert(r.data)
        }
      }, function (e) {
        alert(e.statusText)
      })
    }
  },
  updated () {
    Prism.highlightAll()
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
