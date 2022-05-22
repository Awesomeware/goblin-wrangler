import { defineStore, acceptHMRUpdate } from "pinia";
import API from "@/api/api";

export const useChatStore = defineStore({
  id: "chat_messages",
  state: () => ({
    /** @type {{ username: string; message: string; posted_at: Date }[]} */
    messages: [],
  }),
  getters: {
    all: (state) => state.messages,
  },
  actions: {
    /**
     * @param {{ username: string; message: string; posted_at: Date; }} msg
     */
    say(msg) {
      this.messages.push(msg);
      API.chat.say(msg).then(
        (v) => {
          console.log("Got " + v + " back");
        },
        (err) => {
          console.log("Oh no! Didnt work: " + err);
        }
      );
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useChatStore, import.meta.hot));
}
