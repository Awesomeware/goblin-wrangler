<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { useErrorStore } from "@/stores/error";
import { useAuthStore } from "@/stores/auth";
import { onBeforeUnmount } from "vue";

const error = useErrorStore();
const auth = useAuthStore();

const onCredentialResponse = (credentials: any) => {
  auth.onGoogleSignin(credentials)
};

onBeforeUnmount(() => error.$reset());
</script>

<template>
  <v-app>
    <v-navigation-drawer app theme="dark" permanent>
      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-view-dashboard">
          <RouterLink to="/">Homey</RouterLink>
        </v-list-item>
        <v-list-item prepend-icon="mdi-gavel">
          <RouterLink to="/about">About</RouterLink>
        </v-list-item>
        <v-list-item>
          <div v-google="onCredentialResponse" />
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container fluid>
        <RouterView />
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
@import "./assets/base.css";

a,
.green {
  text-decoration: none;
  color: hsla(160, 100%, 37%, 1);
  transition: 0.4s;
}

@media (hover: hover) {
  a:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

@media (min-width: 1024px) {
  #app {
    padding: 0 2rem;
  }
}
</style>
