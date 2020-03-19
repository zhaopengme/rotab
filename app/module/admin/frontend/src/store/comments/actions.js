import http from "../../common/http";

export default {
  async list(state, params) {
    const result = await http.get("/admin/comments/list", params);
    return result;
  },
  async get(state, id) {
    const result = await http.get("/admin/comments/get", {
      id: id
    });
    return result;
  },
  async delete(state, id) {
    const result = await http.get("/admin/comments/delete", {
      id: id
    });
    return result;
  },
  async saveOrUpdate(state, params) {
    const result = await http.post("/admin/comments/save-or-update", params);
    return result;
  }
};
