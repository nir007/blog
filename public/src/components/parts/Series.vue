<template>
  <div>
    <div class="card my-4">
      <h5 title="click to create" class="card-header pointer" v-on:click="initCreateSeries">
        Series +
      </h5>
      <div v-if="series.length > 0">
        <b-list-group>
          <b-list-group-item
            v-for="(item, index) in series" :key="item.id"
            v-bind:class="{'green-bg': item.is_updated}"
            class="flex-column align-items-start"
          >
            <div v-bind:title="item.count + ' articles'">
              <div class="row">
                <div class="col-md-8">
                  <h5 class="mb-1">{{item.title}}</h5>
                  <p class="mb-1">
                    {{item.description}}
                  </p>
                </div>
                <div class="col-md-4 text-right">
                  <p v-on:click="initEditSeries(item.id, index)" class="pointer text-info">
                    <strong>Edit</strong>
                  </p>
                  <p v-on:click="initRemoveSeries(item.id, index)" class="pointer text-danger">
                    <strong>Remove</strong>
                  </p>
                </div>
              </div>
            </div>
          </b-list-group-item>
        </b-list-group>
        <b-button
          @click="getSeries"
          variant="outline-dark"
          v-if="lastCountGot === limit"
          block
        >More
        </b-button>
      </div>
      <div v-if="series.length === 0" class="text-center">
        <h3 class="pointer margin-top20px margin-bottom20px" v-on:click="initCreateSeries">
          Create the first series
        </h3>
      </div>
    </div>
    <edit-series></edit-series>
    <create-series></create-series>
  </div>
</template>

<script>
import EditSeries from '../modals/EditSeries.vue'
import CreateSeries from '../modals/CreateSeries.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'Series',
  mixins: [ResponseHandler],
  components: {
    EditSeries,
    CreateSeries
  },
  data () {
    return {
      series: [],
      limit: 2,
      lastCountGot: 0,
      editSeriesId: 0,
      removeSeriesId: 0,
      removeSeriesIndex: 0,
      urls: {
        getUserSeries: '/get_user_series',
        deleteSeries: '/delete_series'
      }
    }
  },
  props: {
    authorId: 'Number'
  },
  methods: {
    initCreateSeries: function () {
      this.$root.$emit('init_create_series')
    },
    initEditSeries: function (id, index) {
      this.$root.$emit('init_edit_series', {id: id, index: index})
    },
    initRemoveSeries: function (id, index) {
      this.removeSeriesId = id
      this.removeSeriesIndex = index
      this.$root.$emit('confirm', 'You want to delete this series! Are you sure ?')
    },
    removeSeries: function () {
      this.$http.post(this.urls.deleteSeries, 'id=' + this.removeSeriesId, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      }).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200 && r.data > 0) {
          this.series.splice(this.removeSeriesIndex, 1)
        } else {
          this.responseFailHandle(r)
        }
      })
    },
    getSeries: function () {
      this.$http.post(this.urls.getUserSeries,
        'author_id=' + this.authorId +
        '&limit=' + this.limit +
        '&offset=' + this.series.length,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          for (let i in r.data) {
            this.series.push(r.data[i])
          }
          this.lastCountGot = r.data != null ? r.data.length : 0
        } else {
          this.responseFailHandle(r)
        }
      })
    }
  },
  mounted () {
    this.getSeries()
    let self = this
    this.$root.$on('series_updated', function (updates) {
      for (let i in self.series) {
        if (self.series[i].id === updates.id) {
          self.series[i].title = updates.title
          self.series[i].description = updates.description
          self.series[i].published = updates.published
          self.series[i].count = updates.count
          self.series[i].is_updated = true
        }
      }
    })
    this.$root.$on('created_series', function (series) {
      self.series.push(series)
    })
    this.$root.$on('ok', function () {
      self.removeSeries()
    })
  }
}
</script>
