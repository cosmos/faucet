import Vue from "vue";
import Router from "vue-router";
import Faucet from "./views/Faucet.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/",
      name: "faucet",
      component: Faucet
    }
  ]
});
