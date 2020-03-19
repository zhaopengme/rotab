import "bulma";
import Vue from "vue";

import ViewUI from "view-design";
import "view-design/dist/styles/iview.css";

import "./assets/global.css";

import App from "./App.vue";
// import "./registerServiceWorker";
import router from "./router";
import store from "./store";

import kit from "./common/kit";
import http from "./common/http";

Vue.config.productionTip = false;

Vue.kit = Vue.prototype.$kit = kit;
Vue.http = Vue.prototype.$http = http;

Vue.use(ViewUI);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
