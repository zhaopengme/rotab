import api from "./api";

export default {
  async post(url, data) {
    let response = await api({ method: "post", url, data });
    return response;
  },
  async get(url, params) {
    let response = await api({ method: "get", url, params });
    return response;
  }
};
