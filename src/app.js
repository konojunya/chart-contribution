import Vue from 'vue'
import VueHighcharts from 'vue-highcharts'
import Highcharts from 'highcharts'

import App from './components/app.vue'

Vue.use(VueHighcharts)
Vue.component("app", App)

new Vue({
  el: "#app"
})