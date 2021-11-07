import Vue from 'vue'
import Router from 'vue-router'
import Resource from 'vue-resource'
import Articles from '../components/pages/Articles'
import Article from '../components/pages/Article'
import NewArticle from '../components/pages/NewArticle'
import EditArticle from '../components/pages/EditArticle'
import Person from '../components/pages/Person'
import Persons from '../components/pages/Persons'
import Markdown from '../components/pages/Markdown'
import Founded from '../components/pages/Founded'
import PageNotFound from '../components/pages/PageNotFound'

Vue.use(Resource)
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: Articles
    },
    {
      path: '/articles/:tag?',
      name: 'articles',
      component: Articles
    },
    {
      path: '/a/:id',
      name: 'article',
      component: Article
    },
    {
      path: '/new_article',
      name: 'newArticle',
      component: NewArticle
    },
    {
      path: '/edit/article/:id',
      name: 'editArticle',
      component: EditArticle
    },
    {
      path: '/person/:id?',
      name: 'person',
      props: true,
      component: Person
    },
    {
      path: '/persons',
      name: 'persons',
      component: Persons
    },
    {
      path: '/markdown',
      name: 'markdown',
      component: Markdown
    },
    {
      path: '/founded',
      name: 'founded',
      component: Founded
    },
    {
      path: '*',
      component: PageNotFound
    }
  ]
})
