import http from "../../common/http";

export default {
  async list({ commit }, params) {
    const result = await http.get("/admin/navigations/list", params);
    commit("setList", result);
    return result;
  },
  async all() {
    const result = await http.get("/admin/navigations/all");
    return result;
  },
  async get(state, id) {
    const result = await http.get("/admin/navigations/get", {
      id: id
    });
    return result;
  },
  async delete({ commit }, id) {
    const result = await http.get("/admin/navigations/delete", {
      id: id
    });
    commit("setData", result);
    return result;
  },
  async saveOrUpdate(state, params) {
    const result = await http.post("/admin/navigations/save-or-update", params);
    return result;
  }
};
