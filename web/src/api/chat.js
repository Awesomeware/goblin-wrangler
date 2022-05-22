// @ts-check

/**
 * @param {import("axios").AxiosInstance} axios
 */
export default (axios) => ({
  async say(msg) {
    return axios.post("/chat", msg);
  },

  async getAll() {
    return this.axios.get("/chat");
  },
});
