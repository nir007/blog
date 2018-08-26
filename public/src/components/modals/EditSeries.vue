<template>
  <div>
    <b-modal size="lg"
             @ok="create"
             @shown="clearForm"
             centered
             ref="edit_series"
             title="Edit series"
    >
      <div class="form-group">
        <div class="form-group">
          <label>Title</label>
          <input typeof="text" class="form-control" v-model="title"/>
        </div>
        <div class="form-group">
          <label>Description</label>
          <input typeof="text" class="form-control" v-model="description"/>
        </div>
        <div class="form-group">
          <div class="row">
            <div class="col-md-6 text-right">
              <p>publish:</p>
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
          <draggable v-model="myArray" :options="{group:'people'}" @start="drag=true" @end="drag=false">
            <transition-group>
              <b-list-group v-for="element in myArray" :key="element.id">
                <b-list-group-item button>{{element.name}}</b-list-group-item>
              </b-list-group>
            </transition-group>
            <button slot="footer" @click="loadArticles">Add</button>
          </draggable>
        </div>
        <articles-user-list v-if="loadArticles"></articles-user-list>
      </div>
    </b-modal>
  </div>
</template>

<script>
import Switches from 'vue-switches'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import ArticlesUserList from '../parts/ArticlesUserList'
export default {
  name: 'EditSeries',
  mixins: [ResponseHandler],
  data () {
    return {
      myArray: [
        {
          name: 'Курица',
          people: 2
        },
        {
          name: 'Питух',
          people: 1
        }
      ],
      title: '',
      description: '',
      published: false,
      modalShow: false,
      loadArticles: false,
      warnings: [],
      urls: {
        update: '/update_series',
        get: '/get_one_series'
      }
    }
  },
  components: {
    'switches': Switches,
    'articles-user-list': ArticlesUserList
  },
  props: {
    seriesId: 'Number'
  },
  methods: {
    update: function () {
      if (!this.title) {
        this.warnings.push('enter title')
      }
      if (!this.description) {
        this.warnings.push('enter description')
      }
      if (this.warnings.length > 0) {
        this.$root.$emit('warning', this.warnings)
        return
      }

      let dataToSend = {
        title: this.title,
        description: this.description,
        published: this.published ? 1 : 0
      }

      this.$http.post(this.url, dataToSend, {
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(function (r) {
        try {
          r = JSON.parse(r.bodyText)
          if (r.status === 200) {
            dataToSend.author_id = r.data
            this.$root.$emit('created_series', dataToSend)
          } else {
            this.responseFailHandle(r)
          }
        } catch (e) {
          this.$root.$emit('warning', [e.message])
        }
      }, function (e) {
        this.responseFailHandle({status: 500, data: e, timeout: 10000})
      })
    },
    hideModal: function () {
      this.$refs.create_series.hide()
    },
    clearForm: function () {
      this.warnings = []
      this.title = ''
      this.description = ''
      this.published = false
      this.loadArticles = false
    }
  },
  mounted () {
    var self = this
    this.$root.$on('init_edit_series', function () {
      self.$refs.edit_series.show()
    })
  }
}
</script>
