import { defineStore } from "pinia";
import api from "@/api";
import { LoginRequest } from "@/api/auth";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: null,
  }),

  getters: {
    loggedIn(): boolean {
      return this.user != null;
    },
  },

  actions: {
    async login(credentials: LoginRequest) {
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
