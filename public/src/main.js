// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import NavTop from './components/parts/NavTop'
import BootstrapVue from 'bootstrap-vue'
import router from './router'
import Alarm from './components/parts/Alarm.vue'
import Warning from './components/modals/Warning.vue'

Vue.use(BootstrapVue)

Vue.config.productionTip = false

Vue.component('nav-top', NavTop)
Vue.component('alarm', Alarm)
Vue.component('m-warning', Warning)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<div><alarm></alarm><m-warning></m-warning><nav-top></nav-top><div class="container"><App/></div></div>'
})
