import http from "../../common/http";

export default {
  async list(state, params) {
    const result = await http.get("/admin/posts/list", params);
    return result;
  },
  async get(state, id) {
    const result = await http.get("/admin/posts/get", {
      id: id
    });
    return result;
  },
  async delete(state, id) {
    const result = await http.get("/admin/posts/delete", {
      id: id
    });
    return result;
  },
  async saveOrUpdate(state, params) {
    const result = await http.post("/admin/posts/save-or-update", params);
    return result;
  }
};
