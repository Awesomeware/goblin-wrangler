import { ChatSayRequest } from "@/models/requests";
import { AxiosStatic, AxiosResponse } from "axios";

export default (axios: AxiosStatic) => ({
  say: async(msg: ChatSayRequest): Promise<string> => {
    return axios.post("/api/chat", msg);
  },

  getAll: async(): Promise<string[]> => {
    return axios.get("/api/chat");
  },
});
