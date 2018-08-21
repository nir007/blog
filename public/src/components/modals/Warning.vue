<template>
  <div>
    <b-modal size="sm" centered ref="warning" hide-footer title="Warning">
      <div class="d-block text-center">
        <ul>
          <li v-for="(item) in warnings" :key="item.id">
            {{item}}
          </li>
        </ul>
        <br>
      </div>
      <div class="modal-footer">
        <b-btn variant="outline-danger" block @click="hideModal">
          Ок
        </b-btn>
      </div>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: 'Warning',
  data () {
    return {
      warnings: [],
      modalShow: false
    }
  },
  methods: {
    hideModal: function () {
      this.$refs.warning.hide()
    }
  },
  mounted () {
    var self = this
    this.$root.$on('warning', function (data) {
      self.warnings = data
      self.$refs.warning.show()
    })
  }
}
</script>
