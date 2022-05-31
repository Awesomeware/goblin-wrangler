import { AuthLoginRequest } from "@/models/requests";
import User from "@/models/user.model";
import { AxiosStatic } from "axios";

export default (axios: AxiosStatic) => ({
  login: async(credentials: AuthLoginRequest): Promise<string> => {
    const response = await axios.post("/login", credentials);
    const token = `Bearer ${response.data.token}`;
    axios.defaults.headers.common["Authorization"] = token;
    return token;
  },

  me: async (): Promise<User> => {
    const user = (await axios.get<User>("/me")).data;
    return user || null;
  },
});
