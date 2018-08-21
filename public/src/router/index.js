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
      component: Articles
    },
    {
      path: '/a/:id',
      name: 'Article',
      component: Article
    },
    {
      path: '/new_article',
      name: 'NewArticle',
      component: NewArticle
    },
    {
      path: '/edit/article/:id',
      name: 'EditArticle',
      component: EditArticle
    },
    {
      path: '/person/:id?',
      name: 'Person',
      props: true,
      component: Person
    },
    {
      path: '/persons',
      name: 'Persons',
      component: Persons
    },
    {
      path: '/markdown',
      name: 'Markdown',
      component: Markdown
    },
    {
      path: '*',
      component: PageNotFound
    }
  ]
})
