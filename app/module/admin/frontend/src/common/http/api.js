import axios from "axios"; // 注意先安装哦
import config from "./config.js"; // 倒入默认配置
import qs from "qs"; // 序列化请求数据，视服务端的要求
import { Notice } from "view-design";

// axios.defaults.transformRequest = [data => qs.stringify(data)];

window.$axios = axios;

export default options => {
  return new Promise((resolve, reject) => {
    const instance = axios.create({
      baseURL: config.baseURL(),
      withCredentials: config.withCredentials,
      responseType: config.responseType,
      transformRequest: [data => qs.stringify(data)]
    });

    // request 拦截器
    instance.interceptors.request.use(
      config => {
        // Tip: 1
        // 请求开始的时候可以结合 vuex 开启全屏的 loading 动画

        // Tip: 2
        // 带上 token , 可以结合 vuex 或者重 localStorage
        // if (store.getters.token) {
        //     config.headers['X-Token'] = getToken() // 让每个请求携带token--['X-Token']为自定义key 请根据实际情况自行修改
        // } else {
        //     // 重定向到登录页面
        // }
        const token = localStorage.getItem("token");
        if (token) {
          // config.headers["Content-type"] = "application/x-www-form-urlencoded";
          config.headers["Authorization"] = "Bearer " + token;
        }
        return config;
      },
      error => {
        // 请求错误时做些事(接口错误、超时等)
        // Tip: 4
        // 关闭loadding
        console.log("request:", error);

        //  1.判断请求超时
        if (
          error.code === "ECONNABORTED" &&
          error.message.indexOf("timeout") !== -1
        ) {
          console.log(
            "根据你设置的timeout/真的请求超时 判断请求现在超时了，你可以在这里加入超时的处理方案"
          );
          // return service.request(originalRequest);//例如再重复请求一次
        }
        //  2.需要重定向到错误页面
        const errorInfo = error.response;
        console.log(errorInfo);
        if (errorInfo) {
          // error =errorInfo.data//页面那边catch的时候就能拿到详细的错误信息,看最下边的Promise.reject
          const errorStatus = errorInfo.status; // 404 403 500 ... 等
          console.log(errorStatus);
          // router.push({
          //   path: `/error/${errorStatus}`
          // });
        }
        return Promise.reject(error); // 在调用的那边可以拿到(catch)你想返回的错误信息
      }
    );

    // response 拦截器
    instance.interceptors.response.use(
      response => {
        let data;
        // IE9时response.data是undefined，因此需要使用response.request.response(Stringify后的字符串)
        if (response.data == undefined) {
          data = response.request.response;
        } else {
          data = response.data;
        }
        // 根据返回的code值来做不同的处理（和后端约定）
        // data = data.data && data.data[0] ? data.data[0] : data.data;
        // if (data.code === -1) Message.error(data.msg);
        switch (data.code) {
          case 0:
            data = data.data;
            break;
          case 401:
            Notice.error({
              title: "Notification message",
              desc: data.msg
            });
            location.hash = "#/setup/login";
            break;
          case -1:
            Notice.error({
              title: "Notification message",
              desc: data.msg
            });
            break;
          default:
        }
        // 若不是正确的返回code，且已经登录，就抛出错误
        // const err = new Error(data.description)

        // err.data = data
        // err.response = response

        // throw err
        return data;
      },
      err => {
        if (err && err.response) {
          switch (err.response.status) {
            case 400:
              err.message = "请求错误";
              break;

            case 401:
              err.message = "未授权，请登录";
              break;

            case 403:
              err.message = "拒绝访问";
              break;

            case 404:
              err.message = `请求地址出错: ${err.response.config.url}`;
              break;

            case 408:
              err.message = "请求超时";
              break;

            case 500:
              err.message = "服务器内部错误";
              break;

            case 501:
              err.message = "服务未实现";
              break;

            case 502:
              err.message = "网关错误";
              break;

            case 503:
              err.message = "服务不可用";
              break;

            case 504:
              err.message = "网关超时";
              break;

            case 505:
              err.message = "HTTP版本不受支持";
              break;

            default:
          }
        }
        console.error(err);
        // 此处我使用的是 element UI 的提示组件
        // Message.error(`ERROR: ${err}`);
        return Promise.reject(err); // 返回接口返回的错误信息
      }
    );

    //请求处理
    instance(options)
      .then(res => {
        resolve(res);
        return false;
      })
      .catch(error => {
        reject(error);
      });
  });
};
