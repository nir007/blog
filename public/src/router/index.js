import Vue from 'vue'
import Router from 'vue-router'
import Resource from 'vue-resource'
import Articles from '@/components/pages/Articles'
import Article from '@/components/pages/Article'
import NewArticle from '@/components/pages/NewArticle'
import EditArticle from '@/components/pages/EditArticle'
import Person from '@/components/pages/Person'
import Persons from '@/components/pages/Persons'
import Markdown from '@/components/pages/Markdown'
import PageNotFound from '@/components/pages/PageNotFound'

Vue.use(Resource)
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Articles
    },
    {
      path: '/articles/:tag?',
      name: 'Articles',
      component: Articles,
      beforeEnter: (to, from, next) => {
        next()
      }
    },
    {
      path: '/a/:id',
      name: 'Article',
      component: Article,
      beforeEnter: (to, from, next) => {
        next()
      }
    },
    {
      path: '/new_article',
      name: 'NewArticle',
      component: NewArticle,
      beforeEnter: (to, from, next) => {
        next()
      }
    },
    {
      path: '/edit/article/:id',
      name: 'EditArticle',
      component: EditArticle,
      beforeEnter: (to, from, next) => {
        next()
      }
    },
    {
      path: '/person/:id?',
      name: 'Person',
      props: true,
      component: Person,
      beforeEnter: (to, from, next) => {
        next()
      }
    },
    {
      path: '/persons',
      name: 'Persons',
      component: Persons,
      beforeEnter: (to, from, next) => {
        next()
      }
    },
    {
      path: '/markdown',
      name: 'Markdown',
      component: Markdown,
      beforeEnter: (to, from, next) => {
        next()
      }
    },
    {
      path: '*',
      component: PageNotFound
    }
  ]
})
