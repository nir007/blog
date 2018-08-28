// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import NavTop from './components/parts/NavTop'
import BootstrapVue from 'bootstrap-vue'
import Draggable from 'vuedraggable'
import router from './router'
import Notification from './components/parts/notifications/Notification.vue'
import Confirm from './components/modals/Confim.vue'
import Warning from './components/modals/Warning.vue'

Vue.use(BootstrapVue)

Vue.config.productionTip = false

Vue.component('nav-top', NavTop)
Vue.component('confirm', Confirm)
Vue.component('notification', Notification)
Vue.component('draggable', Draggable)
Vue.component('m-warning', Warning)

Vue.extend()

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<div><confirm/><notification/><m-warning/><nav-top/><div class="container"><App/></div></div>'
})
