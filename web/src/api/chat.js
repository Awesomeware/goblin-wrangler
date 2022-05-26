export default (/** @type import("axios").AxiosStatic */ axios) => ({
  say: function say(msg) {
    return axios.post("/api/chat", msg);
  },

  getAll: function getAll() {
    return axios.get("/api/chat");
  },
});
