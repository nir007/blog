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
          <div class="col-md-9 text-left">
            <p>nickName: {{nickName}}</p>
            <p>about: {{person}}</p>
            <p v-if="isOwner">uuid: {{uuid}}</p>
            <p>
              <small class="text-muted">
                Зареган: {{createdAt}}
              </small>
            </p>
          </div>
        </div>
      </div>
      <articles v-bind:author-id="id"></articles>
    </div>
    <div v-if="userNotFound" class="text-center">
      <img src="/static/assets/img/no-dead-links.jpg" alt="user not found" >
      <h1 class="text-center">User not found</h1>
    </div>
  </div>
  <div class="col-lg-4">
    <div class="card my-4">
      <h5 class="card-header">Search</h5>
      <div class="card-body">
        <div class="input-group">
          <input type="text" class="form-control" placeholder="Search for...">
          <span class="input-group-btn">
            <button class="btn btn-secondary" type="button">Go!</button>
           </span>
        </div>
      </div>
    </div>
    <tags></tags>
  </div>
</div>
</template>

<script>
import Tags from '../parts/Tags.vue'
import Articles from '../parts/Articles.vue'
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'Person',
  mixins: [ResponseHandler, AuthHandler],
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
      avatarPath: '/static/assets/img/',
      urls: {
        getPerson: '/aj_get_person'
      }
    }
  },
  components: {
    'articles': Articles,
    'tags': Tags
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
        } else {
          this.userNotFound = true
          this.responseFailHandle(r)
        }
      }, function () {
        this.userNotFound = true
        this.responseFailHandle({status: 500, data: '500 internal server error'})
      })
  }
}
</script>
