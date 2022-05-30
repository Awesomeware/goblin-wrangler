import { AxiosStatic } from "axios";

export type LoginRequest = {
  email: string;
  password: string;
};

export default (axios: AxiosStatic) => ({
  login: async(credentials: LoginRequest): Promise<string> => {
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
