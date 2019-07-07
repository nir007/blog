<template>
  <div class="row margin-top20px">
    <div class="col-lg-12">
      <div v-if="isLogged && loggedUser.is_confirmed">
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
            Save
          </div>
        </div>
      </div>
    </div>
    <div class="col-lg-12 margin-top50px" v-if="isLogged && !loggedUser.is_confirmed">
      <confirm-phone
        v-bind:is-owner="1"
        v-bind:phone="loggedUser.phone">
      </confirm-phone>
    </div>
    <div v-if="!isLogged && !loggedUser.is_confirmed" class="text-center margin-top50px">
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
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import ConfirmPhone from '../parts/ConfirmPhone.vue'
export default {
  name: 'NewArticle',
  mixins: [ResponseHandler, AuthHandler],
  data () {
    return {
      title: '',
      tags: '',
      text: '',
      published: false,
      isLogged: false,
      loggedUser: false,
      urls: {
        addArticle: '/aj_add_article'
      },
      warnings: []
    }
  },
  components: {
    'confirm-phone': ConfirmPhone,
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
          alert('saved')
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
    var self = this
    this.$root.$on('check_is_logged', function (user) {
      if (user && 'id' in user && user.id > 0) {
        self.isLogged = true
        self.loggedUser = user
      } else {
        self.isLogged = false
      }
    })
  }
}
</script>
