import { defineStore, acceptHMRUpdate } from "pinia";

export type ChatState = {
  messages: string[];
};

export const useChatStore = defineStore({
  id: "chat_messages",
  state: () => ({
    messages: [],
  } as ChatState),
  getters: {
    all: (state) => state.messages,
  },
  actions: {
    say(msg: string) {
      this.messages.push(msg);
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useChatStore, import.meta.hot));
}