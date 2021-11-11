<template>
  <b-modal size="md" id="createGroup" ref="createGroup"  centered hide-footer title="New group">
    <div class="form-group">
      <label>Name</label>
      <b-form-input v-model="name" trim placeholder="Group name"></b-form-input>
    </div>
    <div class="form-group">
      <label>Description</label>
      <b-form-textarea v-model="description"
                       placeholder="Description"
                       :rows="5"
                       :max-rows="10">
      </b-form-textarea>
    </div>
    <div class="modal-footer">
      <b-btn variant="outline-danger" @click="handleCancel">
        Cancel
      </b-btn>
      <b-btn variant="outline-success" @click="create">
        Create
      </b-btn>
    </div>
  </b-modal>
</template>

<script>
export default {
  name: 'createGroup',
  data () {
    return {
      name: '',
      description: '',
      modalShow: false,
      warnings: [],
      url: 'aj_create_group'
    }
  },
  methods: {
    handleCancel: function () {
      this.$refs['createGroup'].hide()
    },

    create: function () {
      this.warnings = []
      if (!this.name) {
        this.warnings.push('enter name')
      }
      if (!this.description) {
        this.warnings.push('enter description')
      }
      if (this.warnings.length > 0) {
        this.$root.$emit('warning', this.warnings)
        return
      }

      let dataToSend = {
        name: this.name,
        description: this.description
      }

      this.$http.post(this.url, dataToSend, {
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(function (r) {
        try {
          r = JSON.parse(r.bodyText)
          if (r.status === 201) {
            dataToSend.id = r.id
            this.$root.$emit('created_group', dataToSend)
          } else {
            this.responseFailHandle(r)
          }
        } catch (e) {
          this.$root.$emit('warning', [e.message])
        }
      }, function (e) {
        this.responseFailHandle({status: 500, data: e, timeout: 10000})
      })
    }
  },
  mounted () {
    const self = this

  }
}
</script>
