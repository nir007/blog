<template>
  <div>
    <b-modal size="md"
             @ok="create"
             @shown="clearForm"
             centered
             ref="createSeries"
             title="New series"
    >
      <div class="form-group">
        <div class="form-group">
          <label>Title</label>
          <input typeof="text" class="form-control" v-model="title" placeholder="New series"/>
        </div>
        <div class="form-group">
          <label>Description</label>
          <input typeof="text" class="form-control" v-model="description" placeholder="This series is about ..."/>
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
      </div>
    </b-modal>
  </div>
</template>

<script>
import Switches from 'vue-switches'
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'CreateSeries',
  mixins: [ResponseHandler],
  data () {
    return {
      title: '',
      description: '',
      published: false,
      modalShow: false,
      warnings: [],
      url: '/create_series'
    }
  },
  components: {
    'switches': Switches
  },
  methods: {
    create: function () {
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
      this.$refs.createSeries.hide()
    },
    clearForm: function () {
      this.warnings = []
      this.title = ''
      this.description = ''
      this.published = false
    }
  },
  mounted () {
    var self = this
    this.$root.$on('init_create_series', function () {
      self.$refs.createSeries.show()
    })
  }
}
</script>
