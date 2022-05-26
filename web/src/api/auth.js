export default (/** @type import("axios").AxiosStatic */ axios) => ({
  login: async (credentials) => {
    const response = await axios.post("/login", credentials);
    const token = `Bearer ${response.data.token}`;
    axios.defaults.headers.common["Authorization"] = token;
    return token;
  },

  me: async () => {
    const user = (await axios.get("/me")).data;
    return user || null;
  },
});
