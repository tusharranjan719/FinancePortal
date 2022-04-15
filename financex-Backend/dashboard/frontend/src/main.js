import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router';

Vue.config.productionTip = false

new Vue({
  vuetify,
  render: h => h(App),
  router,
  watch:{
    '$route' (to) {
       if(to.currentRoute.meta.reload==true){window.location.reload()}
   }
 }
}).$mount('#app')
