<template>
  <div>
    <p class="pointer text-info underline-link" v-if="isOwner" v-on:click="logout">
      Logout
    </p>
    Please confirm your phone number, enter code from SMS. <strong>{{phone}}</strong>
    <div class="form-group">
      <div class="row">
        <div class="col-lg-10">
          <input v-model="code" class="form-control" type="text" maxlength="5" placeholder="code">
          <small v-if="isEmptyCode" class="text-danger">enter code</small>
        </div>
        <div class="col-lg-2">
          <button @click="confirm" type="button" class="btn btn-success">
            Confirm
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Logout from '../mixins/Logout.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'ConfirmPhone',
  mixins: [Logout, ResponseHandler],
  data () {
    return {
      code: '',
      isEmptyCode: false,
      url: 'aj_confirm_phone'
    }
  },
  props: {
    isOwner: 'Boolean',
    phone: 'String'
  },
  methods: {
    confirm () {
      if (!this.code) {
        this.isEmptyCode = true
        return
      }

      this.isEmptyCode = false
      this.$http.post(this.url, 'code=' + this.code,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          location.reload()
        } else {
          this.responseFailHandle(r)
        }
      })
    }
  }
}
</script>
