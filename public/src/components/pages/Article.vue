<template>
  <div class="row">
    <div class="col-lg-8">
      <br>
      <div v-if="!notFound && (article.published > 0 || article.isOwner)">
        <div class="row">
          <div class="col-md-10">
            <p class="lead">
              by
              <a :href="'#/person/' + author.id">
                {{author.nickName}}
              </a> | <small class="text-muted">{{article.createdAt}}</small>
            </p>
          </div>
          <div class="col-md-2 text-right">
            <p v-if="article.isOwner" class="lead">
              <a :href="'#/edit/article/' + article.id">edit</a>
            </p>
          </div>
        </div>
        <p v-if="Object.keys(article.tags).length > 0">
          tags:
          <span v-for="(tag, key) in article.tags" :key="tag.id">
            &nbsp;
            <a :href="'#/articles/' + key">
             [ {{ tag }} ]
            </a>
          </span>
        </p>
        <h1 class="mt-4">{{article.title}}</h1>
        <hr>
        <vue-markdown :source="article.text"></vue-markdown>
        <div class="card my-4">
          <h5 class="card-header">Leave a Comment:</h5>
          <div class="card-body">
          <form>
            <div class="form-group">
            <textarea class="form-control" rows="3"></textarea>
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
          </form>
          </div>
        </div>
        <div class="media mb-4">
          <img class="d-flex mr-3 rounded-circle" src="http://placehold.it/50x50" alt="">
          <div class="media-body">
          <h5 class="mt-0">Commenter Name</h5>
            Cras sit amet nibh libero, in gravida nulla.
          </div>
        </div>
        <div class="media mb-4">
          <img class="d-flex mr-3 rounded-circle" src="http://placehold.it/50x50" alt="">
          <div class="media-body">
            <h5 class="mt-0">Commenter Name</h5>
              Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante sollicitudin. Cras purus odio, vestibulum in vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis in faucibus.
            <div class="media mt-4">
              <img class="d-flex mr-3 rounded-circle" src="http://placehold.it/50x50" alt="">
              <div class="media-body">
                <h5 class="mt-0">Commenter Name</h5>
                  Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante sollicitudin. Cras purus odio, vestibulum in vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis in faucibus.
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="notFound || (article.published <= 0 && !article.isOwner)" class="text-center">
        <h1 class="text-center">Article not found</h1>
        <img src="/static/assets/img/404-error.jpg" alt="404 not found" >
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
    </div>
  </div>
</template>

<script>
import Markdown from 'vue-markdown'
import Prism from 'prismjs'
export default {
  name: 'Article',
  data () {
    return {
      article: {
        id: 0,
        title: '',
        text: '',
        tags: {},
        createdAt: '',
        published: false,
        isOwner: false
      },
      author: {
        id: 0,
        avatar: '',
        nickName: ''
      },
      isLogged: false,
      notFound: false,
      urls: {
        getArticle: 'aj_get_article',
        isLogged: '/aj_is_logged'
      }
    }
  },
  updated () {
    Prism.highlightAll()
  },
  created () {
    this.notFound = false
    this.$http.post(this.urls.getArticle,
      'id=' + this.$route.params.id,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      })
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        console.log(r)
        if (r.status === 200) {
          this.article.id = r.data.article.id
          this.article.title = r.data.article.title
          this.article.text = r.data.article.text
          this.article.tags = r.data.article.tags
          this.article.createdAt = r.data.article.created
          this.article.published = r.data.article.published
          this.article.isOwner = r.data.article.is_owner
          this.author.id = r.data.author.id
          this.author.avatar = r.data.author.avatar
          this.author.nickName = r.data.author.nick_name
        } else {
          this.notFound = true
        }
      }, function () {
        this.notFound = true
      })
  },
  components: {
    'vue-markdown': Markdown
  },
  mounted () {
    this.$http.post(this.urls.isLogged)
      .then(function (r) {
        r = JSON.parse(r.bodyText)
        this.isLogged = r.data
        this.$root.$emit('nav_top_rebuild', r.data)
      })
  }
}
</script>
