export default {
  method: "post",
  // 基础url前缀
  baseURL: () => {
    let baseURL =
      process.env.NODE_ENV === "development"
        ? "http://127.0.0.1:8199"
        : location.origin;
    window.baseURL = baseURL;
    return baseURL;
  },
  // 请求头信息
  headers: {
    "Content-Type": "application/json;charset=UTF-8"
  },
  // 参数
  data: {},
  // 设置超时时间
  timeout: 100000,
  // 携带凭证
  withCredentials: false,
  // 返回数据类型
  responseType: "json"
};
