<template>
  <div>
    <b-navbar toggleable="md" type="dark" fixed="top" class="bg-dark">
      <div class="container">
        <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
        <b-navbar-brand href="#/articles">{{logo}}</b-navbar-brand>
        <b-collapse is-nav id="nav_collapse">
          <ul class="navbar-nav ml-auto">
            <li v-for="item in menu.center" :key="item.id"
                v-if="!item.reqAuth || isLogged"
                v-bind:class="{active: item.active}"
                class="nav-item"
            >
              <a class="nav-link" :href="item.path">
                {{item.title}}
              </a>
            </li>
          </ul>
          <ul class="navbar-nav ml-auto">
            <li v-for="item in menu.right"
                :key="item.id"
                v-bind:class="{active: item.active}"
                v-if="((isLogged && item.reqAuth) ||
                   (!isLogged && item.reqNotAuth))"
                class="nav-item"
            >
              <a v-if="item.path"
                 class="nav-link"
                 :href="item.path + '/' + personId"
              >
                {{item.title}}
              </a>
              <a v-if="(!item.path)"
                 class="nav-link"
                 href="javascript:void(0)"
                 @click="item.click"
              >
                {{item.title}}
              </a>
            </li>
          </ul>
        </b-collapse>
      </div>
    </b-navbar>
    <sing-in></sing-in>
    <join></join>
  </div>
</template>

<script>
import SignIn from '../modals/SignIn.vue'
import Join from '../modals/Join.vue'
export default {
  name: 'NavTop',
  data () {
    return {
      menu: {
        center: {
          articles: {
            title: 'articles',
            path: '#/articles',
            active: false,
            reqAuth: false
          },
          newArticle: {
            title: 'new article',
            path: '#/new_article',
            active: true,
            reqAuth: true
          },
          persons: {
            title: 'persons',
            path: '#/persons',
            active: false,
            reqAuth: false
          }
        },
        right: {
          person: {
            title: 'person',
            path: '#/person',
            click: false,
            active: false,
            reqAuth: true,
            reqNotAuth: false
          },
          join: {
            title: '[ join now! ]',
            path: false,
            click: this.join,
            active: false,
            reqAuth: false,
            reqNotAuth: true
          },
          signIn: {
            title: '[ sign in ]',
            path: false,
            click: this.singIn,
            active: false,
            reqAuth: false,
            reqNotAuth: true
          }
        }
      },
      personId: 0,
      logo: 'rakan-tarakan.com',
      isLogged: false
    }
  },
  components: {
    'sing-in': SignIn,
    'join': Join
  },
  methods: {
    singIn: function () {
      this.$root.$emit('signin', {})
    },
    join: function () {
      this.$root.$emit('join', {})
    },
    lightItem () {
      let name = this.$router.currentRoute.name
      for (let i in this.menu) {
        for (let j in this.menu[i]) {
          this.menu[i][j].active = j === name
        }
      }
    }
  },
  mounted () {
    const self = this
    this.$root.$on('check_is_logged', function (user) {
      self.lightItem()
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
