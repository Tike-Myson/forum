import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    authorized: false,
  },
  mutations: {
    toggleAuthState(state) {
      state.authorized = true;
    },
  },
  getters: {
    authState(state) {
      return state.authorized;
    },
  },
  actions: {},
  modules: {},
});
