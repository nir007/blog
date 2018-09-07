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
export default {
  name: 'ConfirmPhone',
  mixins: [Logout],
  data () {
    return {
      code: '',
      url: 'aj_confirm_phone'
    }
  },
  props: {
    isOwner: 'Boolean',
    phone: 'String'
  },
  methods: {
    confirm () {
      this.$http.post(this.url, 'code=' + this.code,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        console.log(r)
      })
    }
  }
}
</script>
