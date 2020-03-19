import http from "../../common/http";

export default {
  async getByCategory(state, params) {
    const result = await http.get("/admin/option/get-by-category", params);
    return result;
  },
  async updateGeneralOption(state, params) {
    const result = await http.post(
      "/admin/option/update-general-option",
      params
    );
    return result;
  }
};
