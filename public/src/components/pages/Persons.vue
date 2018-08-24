<template>
<div class="row">
  <div class="col-lg-8">
    <br>
    <div v-for="(person) in persons" :key="person.id" class="card p-3 margin-bottom5px">
      <div class="row">
        <div class="col-md-3">
          <a v-if="person.avatar" :href="'#/person/' + person.id">
            <img class="avatar" :src="avatarPath + person.avatar">
          </a>
          <div v-if="!person.avatar" class="default-avatar"></div>
        </div>
        <div class="col-md-9 text-left">
          <p>
            <a :href="'#/person/' + person.id">
              nickName: {{person.nick_name}}
            </a>
          </p>
          <p>about: {{person.person}}</p>
          <p>
            <small class="text-muted">
              Зареган: {{person.created_at}}
            </small>
          </p>
        </div>
      </div>
    </div>
  </div>
  <div class="col-lg-4">
    <div class="card my-4">
      <h5 class="card-header">Search</h5>
      <div class="card-body">
        <div class="input-group">
          <input type="text" class="form-control" placeholder="Search for...">
          <div class="input-group-btn">
            <button class="btn btn-secondary" type="button">Go!</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  name: 'Persons',
  mixins: [ResponseHandler, AuthHandler],
  data () {
    return {
      persons: [],
      avatarPath: '/static/assets/img/',
      urls: {
        getPersons: '/aj_get_persons'
      }
    }
  },
  mounted () {
    this.$http.post(this.urls.getPersons,
      'limit=' + this.limit +
      '&offset=' + this.persons.length,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      }
    )
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          for (let i in r.data) {
            this.persons.push(r.data[i])
          }
        } else {
          this.responseFailHandle(r)
        }
      }, function () {
        this.responseFailHandle({status: 500, data: '500 internal server error'})
      })
  }
}
</script>
