<template>
<div class="row">
  <div class="col-lg-8">
    <br>
    <div v-if="!userNotFound">
      <div class="card p-3">
        <div class="row">
          <div class="col-md-3">
            <img class="avatar" v-bind:src="avatar">
          </div>
          <div class="col-md-7 text-left">
            <p>nickname: <strong>{{nickName}}</strong></p>
            <p>about: <strong>{{person}}</strong></p>
            <p v-if="isOwner">uuid: {{uuid}}</p>
            <p>
              <small class="text-muted">
                Зареган: {{createdAt}}
              </small>
            </p>
          </div>
          <div class="col-md-2 text-center">
            <span class="pointer" v-if="isOwner" v-on:click="userLogout">
              Logout
            </span>
          </div>
        </div>
      </div>
      <articles v-if="userLoaded" v-bind:author-id="id"></articles>
    </div>
    <div v-if="userNotFound" class="text-center">
      <img src="/static/assets/img/no-dead-links.jpg" alt="user not found" >
      <h1 class="text-center">User not found</h1>
    </div>
  </div>
  <div class="col-lg-4">
    <search></search>
    <series v-if="userLoaded" v-bind:author-id="id"></series>
    <tags></tags>
  </div>
</div>
</template>

<script>
import Tags from '../parts/Tags.vue'
import Articles from '../parts/Articles.vue'
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import Logout from '../mixins/Logout.vue'
import Search from '../parts/Search.vue'
import Series from '../parts/series.vue'
export default {
  name: 'Person',
  mixins: [ResponseHandler, AuthHandler, Logout],
  data () {
    return {
      id: 0,
      uuid: '',
      person: '',
      nickName: '',
      avatar: '',
      createdAt: '',
      isOwner: false,
      userNotFound: false,
      userLoaded: false,
      avatarPath: '/static/assets/img/',
      urls: {
        getPerson: '/aj_get_person'
      }
    }
  },
  components: {
    'articles': Articles,
    'search': Search,
    'tags': Tags,
    'series': Series
  },
  methods: {
    userLogout: function () {
      this.logout()
      setTimeout(function () {
        location.href = '#/articles'
      }, 500)
    }
  },
  created () {
    var id = typeof this.$route.params.id !== 'undefined'
      ? this.$route.params.id : 0
    this.id = id
    this.$http.post(this.urls.getPerson,
      'id=' + id,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      })
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          this.id = r.data.id
          this.uuid = r.data.uuid
          this.person = r.data.person
          this.nickName = r.data.nick_name
          this.avatar = this.avatarPath + r.data.avatar
          this.createdAt = r.data.created_at
          this.isOwner = r.data.is_owner
          this.userNotFound = false
          this.userLoaded = true
        } else {
          this.userNotFound = true
          this.responseFailHandle(r)
        }
      }, function () {
        this.userNotFound = true
        this.responseFailHandle({status: 500, data: 'Internal server error'})
      })
  }
}
</script>
