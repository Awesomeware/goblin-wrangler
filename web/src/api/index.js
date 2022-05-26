import makeChat from "@/api/chat";
import makeAuth from "@/api/auth";
import axios from "./axios";

export default {
  chat: makeChat(axios),
  auth: makeAuth(axios),
};
