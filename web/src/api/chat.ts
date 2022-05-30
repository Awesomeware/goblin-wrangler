import { AxiosStatic, AxiosResponse } from "axios";

type SayRequest = {
  message: string;
};

export default (axios: AxiosStatic) => ({
  say: async(msg: SayRequest): Promise<string> => {
    return axios.post("/api/chat", msg);
  },

  getAll: async(): Promise<string[]> => {
    return axios.get("/api/chat");
  },
});
