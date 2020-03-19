import http from "../../common/http";

export default {
  async login({ commit }, params) {
    const result = await http.post("/admin/login", params);
    if (result.token) {
      params.token = result.token;
      params.usernname = result.usernname;
      commit("setUser", params);
    }
    return result;
  },
  async getUserinfo({ commit }) {
    const result = await http.get("/admin/user/info", {});
    commit("setUserinfo", result);
    return result;
  },
  async updateUserinfo(state, params) {
    const result = await http.post("/admin/user/update-info", params);
    return result;
  },
  async updatePassword(state, params) {
    const result = await http.post("/admin/user/update-password", params);
    return result;
  }
};
