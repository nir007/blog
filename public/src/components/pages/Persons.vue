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
            nickName:
            <a :href="'#/person/' + person.id">
              {{person.nick_name}}
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
    <search></search>
  </div>
</div>
</template>

<script>
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import Search from '../parts/Search.vue'
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
  components: {
    'search': Search
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
