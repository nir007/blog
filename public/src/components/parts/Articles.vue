<template>
  <div>
    <div v-if="articles.length > 0" class="card margin-top22px" v-for="(item) in articles" :key="item.id">
      <div class="card-body">
        <div class="row">
          <div class="col-md-10">
            <h2 class="post__title">
              <a :href="'#/a/' + item.id" class="post__title_link">
                <strong>{{ item.title }}</strong>
              </a>
            </h2>
            <small>
              <span v-for="(tag, key) in item.tags" :key="tag.id">
                <a :href="'#/articles/' + key">
                  {{ tag }}
                </a>
              </span>
            </small>
          </div>
          <div class="col-md-2 text-right">
            <a v-if="item.is_owner" :href="'#/edit/article/' + item.id">
              edit
            </a>
          </div>
        </div>
      </div>
    </div>
    <div v-if="isEmpty" class="text-center">
      <img src="/static/assets/img/no-dead-links.jpg" alt="user not found" >
      <h1 class="text-center">Articles not found</h1>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      isEmpty: false,
      articles: [],
      urls: {
        get_articles: '/get_articles'
      },
      limit: 10
    }
  },
  props: {
    authorId: 'Number',
    tag: 'String'
  },
  updated () {
    console.log(this.articles)
  },
  created () {
    let authorId = typeof this.authorId !== 'undefined'
      ? this.authorId : 0

    let tag = typeof this.tag !== 'undefined'
      ? this.tag : ''

    this.$http.post(this.urls.get_articles,
      'author_id=' + authorId +
      '&tag=' + tag +
      '&limit=' + this.limit +
      '&offset' + this.articles.length,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      }
    ).then(function (r) {
      r = JSON.parse(r.bodyText)
      console.log(r)
      if (r.status === 200) {
        for (let i in r.data) {
          this.articles.push(r.data[i])
        }
        this.isEmpty = this.articles === 0
      }
    }, function (e) {
      if (!e.statusText) {
        e.statusText = 'Some kind of error happened'
      }
      this.$root.$emit('alarm', e.statusText)
    })
  }
}
</script>
