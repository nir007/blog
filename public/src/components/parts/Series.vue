<template>
  <div>
    <div class="card my-4">
      <h5 title="click to create" class="card-header pointer" v-on:click="initCreateSeries">
        Series +
      </h5>
      <div v-if="series.length > 0">
        <b-list-group>
          <b-list-group-item v-for="(item) in series" :key="item.id" class="flex-column align-items-start">
            <div v-bind:title="item.count + ' articles'" class="d-flex w-100 justify-content-between">
              <h5 class="mb-1">{{item.title}}</h5>
              <small v-on:click="initEditSeries" class="pointer">
                Edit
              </small>
            </div>
            <p class="mb-1">
              {{item.description}}
            </p>
          </b-list-group-item>
        </b-list-group>
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
export default {
  name: 'Series',
  components: {
    EditSeries,
    CreateSeries
  },
  data () {
    return {
      series: [],
      url: '/get_user_series'
    }
  },
  props: {
    authorId: 'Number'
  },
  methods: {
    initCreateSeries: function () {
      this.$root.$emit('init_create_series')
    },
    initEditSeries: function () {
      alert('SKA')
      this.$root.$emit('init_edit_series')
    },
    getSeries: function () {
      this.$http.post(this.url, 'author_id=' + this.authorId, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      }).then(function (r) {
        r = JSON.parse(r.bodyText)
        console.log(r)
        if (r.status === 200) {
          for (let i in r.data) {
            this.series.push(r.data[i])
          }
        } else {
          this.responseFailHandle(r)
        }
      })
    }
  },
  mounted () {
    this.getSeries()
    let self = this
    this.$root.$on('created_series', function (series) {
      self.series.push(series)
    })
  }
}
</script>
