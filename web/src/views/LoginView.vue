<template>
  <div
    id="g_id_onload"
    :data-client_id="googleClientID"
    :data-login_uri="googleLoginURI"
  />
  <div class="g_id_signin"
    data-type="standard"
    data-size="medium"
    data-theme="outline"
    data-text="continue_with"
    data-shape="rectangular"
  />
</template>

<script setup lang="ts">
import { ref, onBeforeUnmount } from "vue";
import { useErrorStore } from "@/stores/error";

const error = useErrorStore();

onBeforeUnmount(() => error.$reset());

const googleClientID = ref(import.meta.env.VITE_GOOGLE_CLIENT_ID);
const googleLoginURI = ref(import.meta.env.VITE_API_BASE_URL + "/" + import.meta.env.VITE_API_SSO_LOGIN_PATH);
</script>
