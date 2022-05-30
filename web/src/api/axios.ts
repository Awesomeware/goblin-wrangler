import axios from "axios";
import { useErrorStore } from "@/stores/error";

axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL;
axios.defaults.headers.common["Authorization"] = localStorage.getItem("token") || "";
axios.defaults.withCredentials = true;

axios.interceptors.request.use(
  function (config) {
    useErrorStore().$reset();
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

axios.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    switch (error.response.status) {
      case 401:
        localStorage.removeItem("token");
        window.location.reload();
        break;
      case 403:
      case 404:
        // TODO: Send to error page using router
        break;
      case 422:
        useErrorStore().$state = error.response.data;
        break;
      default:
        console.log(error.response.data);
    }

    return Promise.reject(error);
  }
);

export default axios;
