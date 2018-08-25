<template>
  <div class="card my-4">
    <h5 class="card-header">Search</h5>
    <div class="card-body">
      <div class="input-group">
        <input
               v-model="string"
               type="text"
               v-bind:class="{'form-control-danger': danger}"
               v-on:input="danger = false"
               class="form-control"
               placeholder="Search for..."
        >
        <div class="input-group-btn">
          <button v-on:click="go" class="btn btn-secondary" type="button">Go!</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Search',
  data () {
    return {
      string: '',
      danger: false
    }
  },
  methods: {
    go: function () {
      if (this.string.length) {
        location.href = '#/founded?s=' + this.string
      } else {
        this.danger = true
      }
    }
  }
}
</script>
