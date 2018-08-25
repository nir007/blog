<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top">
      <div class="container">
        <a class="navbar-brand" href="#/">{{logo}}</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
          <ul class="navbar-nav ml-auto">
            <li class="nav-item active">
              <a class="nav-link " href="#/articles">
                articles
              </a>
            </li>
            <li v-if="isLogged" class="nav-item">
              <a class="nav-link" href="#/new_article">
                new article
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#/persons">
                persons
              </a>
            </li>
          </ul>
          <ul class="navbar-nav ml-auto">
            <li v-if="isLogged" class="nav-item">
              <a class="nav-link" :href="'#/person/' + personId">
                profile
              </a>
            </li>
            <li v-if="!isLogged" class="nav-item">
              <a class="nav-link" href="javascript:void(0)" @click="join">
                [ join now! ]
              </a>
            </li>
            <li v-if="!isLogged" class="nav-item">
              <a class="nav-link" href="javascript:void(0)" @click="singIn">
                [ sign in ]
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
    <sing-in></sing-in>
    <join></join>
  </div>
</template>

<script>
import SingIn from '../modals/SignIn.vue'
import Join from '../modals/Join.vue'
export default {
  name: 'NavTop',
  data () {
    return {
      personId: 0,
      logo: 'rakan-tarakan.com',
      isLogged: false
    }
  },
  components: {
    'sing-in': SingIn,
    'join': Join
  },
  methods: {
    singIn: function () {
      this.$root.$emit('signin', {})
    },
    join: function () {
      this.$root.$emit('join', {})
    }
  },
  mounted () {
    var self = this
    this.$root.$on('check_is_logged', function (user) {
      if (user && 'id' in user && user.id > 0) {
        self.isLogged = true
        self.personId = user.id
      } else {
        self.isLogged = false
      }
    })
  }
}
</script>
