<template>
  <div>
    <b-modal size="lg"
             @ok="update"
             centered
             ref="editSeries"
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
          <h4 class="text-center">Articles of this series: </h4>
          <draggable v-model="articlesOfSeries" :options="{group:'title'}" @start="drag=true" @end="drag=false">
            <b-list-group v-for="(item, index) in articlesOfSeries" :key="item.id">
              <b-list-group-item>
                <div class="row">
                  <div class="col-md-1 text-left">
                    <strong>
                      <b-badge variant="light" pill>{{index}}</b-badge>
                    </strong>
                  </div>
                  <div class="col-md-9">
                    {{item.title}}
                  </div>
                  <div class="col-md-2 text-right">
                    <strong v-on:click="removeArticleFromSeries(index)" class="pointer text-danger">
                      Remove
                    </strong>
                  </div>
                </div>
              </b-list-group-item>
            </b-list-group>
          </draggable>
        </div>
        <hr class="margin-top20px margin-bottom20px">
        <div v-if="!isNotArticles" class="form-group">
          <h4 class="text-center">Select articles: </h4>
          <b-list-group v-for="(item, index) in articles" :key="item.id">
            <b-list-group-item>
              <div class="row">
                <div class="col-md-1 text-left">
                  <strong>
                    <b-badge variant="light" pill>{{index}}</b-badge>
                  </strong>
                </div>
                <div class="col-md-9">
                  {{item.title}}
                </div>
                <div class="col-md-2 text-right">
                  <strong v-on:click="addArticleToSeries(index)" class="pointer text-success">
                    Add
                  </strong>
                </div>
              </div>
            </b-list-group-item>
          </b-list-group>
        </div>
        <div v-if="isNotArticles" class="form-control text-center">
          <h3 class="text-center">
            <a href="#/new_article">Crate the first article</a>
          </h3>
        </div>
      </div>
    </b-modal>
  </div>
</template>

<script>
import Switches from 'vue-switches'
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'EditSeries',
  mixins: [ResponseHandler],
  data () {
    return {
      seriesId: 0,
      articlesOfSeries: [],
      articles: [],
      isNotArticles: false,
      title: '',
      description: '',
      published: false,
      count: 0,
      modalShow: false,
      limit: 10,
      warnings: [],
      urls: {
        update: '/update_series',
        getOneSeries: '/get_one_series',
        getPublishedArticles: '/get_published_articles'
      }
    }
  },
  components: {
    'switches': Switches
  },
  methods: {
    addArticleToSeries: function (index) {
      this.articlesOfSeries.push(this.articles[index])
      this.articles.splice(index, 1)
    },
    removeArticleFromSeries: function (index) {
      this.articles.push(this.articlesOfSeries[index])
      this.articlesOfSeries.splice(index, 1)
    },
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

      for (let i in this.articlesOfSeries) {
        this.articlesOfSeries[i].text = ''
        this.articlesOfSeries[i].tags = {}
        this.articlesOfSeries[i].order = i
      }

      let dataToSend = {
        id: this.seriesId,
        title: this.title,
        description: this.description,
        published: this.published ? 1 : 0,
        articles: this.articlesOfSeries
      }

      console.log(dataToSend)

      this.$http.post(this.urls.update, dataToSend, {
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          this.$root.$emit('series_updated', dataToSend)
        } else {
          this.responseFailHandle(r)
        }
      }, function (e) {
        this.responseFailHandle({status: 500, data: e, timeout: 10000})
      })
    },
    hideModal: function () {
      this.$refs.create_series.hide()
    }
  },
  mounted () {
    var self = this
    this.$root.$on('init_edit_series', function (el) {
      self.seriesId = el.id
      self.id = 0
      self.published = false
      self.title = ''
      self.sescription = ''
      self.articlesOfSeries = []
      self.articles = []
      self.$refs.editSeries.show()
      self.$http.post(self.urls.getOneSeries,
        'series_id=' + el.id,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          self.id = r.data.id
          self.title = r.data.title
          self.description = r.data.description
          self.published = r.data.published
          self.count = r.data.count
        } else {
          self.responseFailHandle(r)
        }
      })

      self.$http.post(self.urls.getPublishedArticles,
        'limit=' + self.limit +
        '&offset=' + self.articles.length,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          for (let i in r.data) {
            self.articles.push(r.data[i])
          }
          self.isNotArticles = self.articles.length === 0
        } else {
          self.responseFailHandle(r)
        }
      })
    })
  }
}
</script>
