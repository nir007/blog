<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top">
      <div class="container">
        <a class="navbar-brand" href="#/">{{logo}}</a>
        <button class="navbar-toggler" type="button"
                data-toggle="collapse"
                data-target="#navbarResponsive"
                aria-controls="navbarResponsive"
                aria-expanded="false"
                aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
          <ul class="navbar-nav ml-auto">
            <li v-for="item in centralItems" :key="item.id" v-bind:class="{active: item.active}" class="nav-item">
              <a class="nav-link" :href="item.path">
                {{item.title}}
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
      centralItems: {
        articles: {
          title: 'articles',
          path: '#/articles',
          active: false
        },
        newArticle: {
          title: 'new article',
          path: '#/new_article',
          active: true
        },
        persons: {
          title: 'persons',
          path: '#/persons',
          active: false
        }
      },
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
    console.log(this.$router.currentRoute.fullPath)

    for (let i in this.centralItems) {
      if (this.centralItems[i]) { }
    }

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
