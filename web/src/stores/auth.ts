import { acceptHMRUpdate, defineStore } from "pinia";
import api from "@/api";
import { AuthLoginRequest } from "@/models/requests";
import User from "@/models/user.model";
import * as jose from 'jose';

export type AuthState = {
  user: User | null;
  googleID: string | null;
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
    email: null,
    googleID: null
  } as AuthState),

  getters: {
    loggedIn(): boolean {
      return this.user != null;
    },
  },

  actions: {
    async loginWithGoogle(credentials: GoogleCredentialResponse) {
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
          // TODO: Probably acquire our own JWT from backend that we can use to auth against the backend,
          // in addition to the Google one. All the Google one really does is replace username/password,
          // by providing a JWT with a verified Google user with a specific e-mail and Google ID.
          return true;
        } else {
          console.log('Not valid!');
        }
      } catch (err) {
        console.error('Err:', err);
      }
      return false;
    },

    async login(credentials: AuthLoginRequest) {
      const token = await api.auth.login(credentials);
      localStorage.setItem("token", token);
      await this.getUser();
    },

    async logout() {
      localStorage.removeItem("token");
      // @ts-ignore
      google.accounts.id.disableAutoSelect()
      const gID = useAuthStore().googleID;
      if (gID) {
        // @ts-ignore
        google.accounts.id.revoke(gID);
      }
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