<template>
  <div>
    <b-modal size="sm" centered ref="confirm" hide-footer title="Confirm">
      <div class="d-block text-center">
        <h4>{{text}}</h4>
      </div>
      <div class="modal-footer">
        <b-btn variant="outline-danger" @click="handleOK">
          Ok
        </b-btn>
        <b-btn variant="outline-success" @click="handleCancel">
          Cancel
        </b-btn>
      </div>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: 'Confirm',
  data () {
    return {
      text: '',
      modalShow: false
    }
  },
  methods: {
    handleOK: function () {
      this.$root.$emit('ok')
      this.$refs.confirm.hide()
    },
    handleCancel: function () {
      this.$root.$emit('cancel')
      this.$refs.confirm.hide()
    }
  },
  mounted () {
    var self = this
    this.$root.$on('confirm', function (text) {
      self.text = text
      self.$refs.confirm.show()
    })
  }
}
</script>
