import Vue from "vue";
import Vuex from "vuex";

import user from "./user";
import option from "./option";
import categories from "./categories";
import navigations from "./navigations";
import comments from "./comments";
import posts from "./posts";

import createLogger from "vuex/dist/logger";
Vue.use(Vuex);
const debug = process.env.NODE_ENV !== "production";
export default new Vuex.Store({
  modules: {
    user,
    option,
    categories,
    navigations,
    comments,
    posts
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
});
