import axios, { InternalAxiosRequestConfig, AxiosResponse } from "axios";
import { useUserStoreHook } from "@/store/modules/user";
import { ResultEnum } from "@/enums/ResultEnum";
import { TOKEN_KEY } from "@/enums/CacheEnum";
import qs from "qs";
import {
  SetFingerprintToken,
  GetHash,
  GetFingerprint,
} from "@/utils/fingerprint";

// 创建 axios 实例
const service = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API, // 基础请求地址，来自环境变量
  timeout: 50000, // 请求超时时间设置为 50 秒
  headers: { "Content-Type": "application/json;charset=utf-8" }, // 请求头设置
  paramsSerializer: (params) => {
    return qs.stringify(params); // 序列化请求参数
  },
});

// 请求拦截器
service.interceptors.request.use(
  async (config: InternalAxiosRequestConfig) => {
    // 更新指纹
    SetFingerprintToken();
    config.headers["Content-Type"] = "application/json;charset=utf-8";

    // 获取本地存储中的访问令牌
    const accessToken = localStorage.getItem(TOKEN_KEY);
    if (accessToken) {
      config.headers.Authorization = accessToken; // 如果有访问令牌，添加到请求头中
    }

    // 获取指纹并设置到请求头
    const fingerprint = await GetHash();
    if (fingerprint) {
      config.headers.Session = fingerprint;
    }

    // 判断请求的 url 是否是完整链接
    if (/^https?:\/\//.test(config.url || "")) {
      config.baseURL = ""; // 如果是完整链接，则不使用 baseURL
    }

    return config;
  },
  (error: any) => {
    return Promise.reject(error); // 请求错误处理
  }
);

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    // 检查配置的响应类型是否为二进制类型（'blob' 或 'arraybuffer'），如果是，直接返回响应对象
    if (
      response.config.responseType === "blob" ||
      response.config.responseType === "arraybuffer"
    ) {
      return response;
    }

    // 处理响应数据
    const { code, data, msg } = response.data;
    if (code === ResultEnum.SUCCESS || code === ResultEnum.OK) {
      return data; // 如果响应成功，返回数据部分
    }

    // 如果响应错误，弹出错误信息
    ElMessage.error(msg || "系统出错");
    return Promise.reject(new Error(msg || "Error"));
  },
  (error: any) => {
    // 异常处理
    if (error.response.data) {
      const { code, msg } = error.response.data;
      if (code === ResultEnum.TOKEN_INVALID) {
        // 如果令牌无效，通知用户并重新登录
        ElNotification({
          title: "提示",
          message: "您的会话已过期，请重新登录",
          type: "info",
        });
        useUserStoreHook()
          .resetToken()
          .then(() => {
            location.reload(); // 重载页面以重新登录
          });
      } else {
        // 其他错误，弹出错误信息
        ElMessage.error(msg || "系统出错");
      }
    }
    return Promise.reject(error.message); // 返回错误信息
  }
);

// 导出 axios 实例
export default service;
