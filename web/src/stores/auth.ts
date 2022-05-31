import { acceptHMRUpdate, defineStore } from "pinia";
import api from "@/api";
import { AuthLoginRequest } from "@/models/requests";
import User from "@/models/user.model";
import * as jose from 'jose';

export type AuthState = {
  user: User | null;
  email: string | null;
};

export type GoogleCredentialResponse = {
  clientId: string;
  credential: string;
  select_by: "btn";
};

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: null,
    email: null
  } as AuthState),

  getters: {
    loggedIn(): boolean {
      return this.user != null;
    },
  },

  actions: {
    onGoogleSignin(credentials: GoogleCredentialResponse) {
      try {
        const jwt: any = jose.decodeJwt(credentials.credential);
        console.log('Decoded response:', jwt);
        const now = Math.floor(Date.now()/1000);
        if (jwt.iss == "https://accounts.google.com" &&
            jwt.aud == import.meta.env.VITE_GOOGLE_CLIENT_ID &&
            jwt.email &&
            (jwt.email_verified || jwt.hd) &&
            (jwt.exp && jwt.exp >= now)
        ) {
          console.log('Seems valid!');
          this.email = jwt.email;
          // TODO: Now toss JWT to backend and ask for user details, and possibly
          // get/set some kind of session token to also track session, since Google no longer
          // does that ^^
        } else {
          console.log('Not valid!');
        }
      } catch (err) {
        console.error('Err:', err);
      }
    },

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