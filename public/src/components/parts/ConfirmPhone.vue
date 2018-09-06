<template>
  <div>
    <p class="pointer text-info underline-link" v-if="isOwner" v-on:click="logout">
      Logout
    </p>
    Please confirm your phone number, enter code from SMS. <strong>{{phone}}</strong>
    <div class="form-group">
      <input class="form-control" type="text" maxlength="5" placeholder="code">
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

    }
  },
  props: {
    isOwner: 'Boolean',
    phone: 'String'
  }
}
</script>
