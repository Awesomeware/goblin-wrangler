import makeChat from "@/api/chat";
import axios from "axios";

class API {
  constructor() {
    if (API._instance) {
      return API._instance;
    }

    API._instance = this;

    this.axios = axios.create({
      baseURL: import.meta.env.API_BASE_URL,
      timeout: 1000,
    });

    // https://github.com/axios/axios#interceptors might be useful later
    // to catch auth errors / token renewal?
    this.axios.interceptors.response.use(
      () => {},
      (error) => {
        console.log("Error: " + error);
        return Promise.reject(error);
      }
    );

    this.chat = makeChat(this.axios);
  }
}

API._instance = undefined;

export default new API();
