import { acceptHMRUpdate, defineStore } from "pinia";
import api from "@/api";
import { AuthLoginRequest } from "@/models/requests";
import User from "@/models/user.model";

export type AuthState = {
  user: User | null;
};

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: null,
  } as AuthState),

  getters: {
    loggedIn(): boolean {
      return this.user != null;
    },
  },

  actions: {
    async login(credentials: AuthLoginRequest) {
      const token = await api.auth.login(credentials);
      localStorage.setItem("token", token);
      await this.getUser();
    },

    async logout() {
      localStorage.removeItem("token");
      this.$reset();
    },

    async getUser() {
      this.user = await api.auth.me();
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAuthStore, import.meta.hot));
}