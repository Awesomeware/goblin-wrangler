export default (/** @type import("axios").AxiosStatic */ axios) => ({
  login: async (credentials) => {
    //const response = await axios.post("/login", credentials);
    const response = { credentials, data: { token: "FakeToken123" } }; // TODO: Replace with real API call
    const token = `Bearer ${response.data.token}`;
    axios.defaults.headers.common["Authorization"] = token;
    return token;
  },

  me: async () => {
    //const user = (await axios.get("/me")).data;
    const user = { name: "fake user" };
    return user || null; // TODO: Once /me is done, will need to verify whether we need to explicitly return null here.
  },
});
