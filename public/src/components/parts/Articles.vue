<template>
  <div>
    <div v-if="!isEmpty" class="card margin-top22px" v-for="(item, index) in articles" :key="item.id">
      <div v-if="!isEmpty || (item.isOwner)" v-bind:class="{hiddenBlock: !item.published}" class="card-body">
        <div class="row">
          <div class="col-md-10">
            <h2 class="post__title">
              <a :href="'#/a/' + item.id" class="post__title_link">
                {{ item.title }}
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
              <strong class="text-info">Edit</strong>
            </a>
            <p v-if="item.is_owner" @click="initRemoveArticle(item.id, index)">
                <strong class="text-danger pointer">Remove</strong>
            </p>
          </div>
        </div>
      </div>
    </div>
    <div v-if="isEmpty" class="text-center">
      <img src="/static/assets/img/no-dead-links.jpg" alt="articles not found" >
      <h1 class="text-center">Articles not found</h1>
    </div>
  </div>
</template>

<script>
import ResponseHandler from '../mixins/ResponseHandler.vue'
export default {
  mixins: [ResponseHandler],
  data () {
    return {
      isEmpty: false,
      articles: [],
      removeArticleId: 0,
      removeArticleIndex: 0,
      urls: {
        get: '/get_articles',
        remove: '/remove_article'
      },
      limit: 5
    }
  },
  props: {
    authorId: 'Number',
    tag: 'String',
    showPublished: 'Number'
  },
  methods: {
    initRemoveArticle: function (id, index) {
      this.removeArticleId = id
      this.removeArticleIndex = index
      this.$root.$emit('confirm', 'You want to delete this article! Are you sure ?')
    },
    remove () {
      this.$http.post(this.urls.remove,
        'id=' + this.removeArticleId,
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
          }
        }
      ).then(function (r) {
        r = JSON.parse(r.bodyText)
        if (r.status === 200) {
          this.articles.splice(this.removeArticleIndex, 1)
        } else {
          this.responseFailHandle(r)
        }
      }, function () {
        this.responseFailHandle({status: 500, data: '500 internal server error'})
      })
    }
  },
  mounted () {
    let self = this
    this.$root.$on('ok', function () {
      self.remove()
    })

    let authorId = typeof this.authorId !== 'undefined'
      ? this.authorId : 0

    let tag = typeof this.tag !== 'undefined'
      ? this.tag : ''

    let showPublished = typeof this.showPublished !== 'undefined'
      ? this.showPublished : 0

    this.$http.post(this.urls.get,
      'author_id=' + authorId +
      '&tag=' + tag +
      '&limit=' + this.limit +
      '&offset=' + this.articles.length +
      '&show_published=' + showPublished,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      }
    ).then(function (r) {
      r = JSON.parse(r.bodyText)
      if (r.status === 200) {
        for (let i in r.data) {
          this.articles.push(r.data[i])
        }
        this.isEmpty = this.articles.length === 0
      } else {
        this.responseFailHandle(r)
      }
    }, function () {
      this.responseFailHandle({status: 500, data: '500 internal server error'})
    })
  }
}
</script>
