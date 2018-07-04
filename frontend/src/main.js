import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store/index";
import Vuelidate from "vuelidate";
import VueAnalytics from "vue-analytics";

Vue.use(Vuelidate);
Vue.use(VueAnalytics, { id: "UA-51029217-5", router: router });
Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
