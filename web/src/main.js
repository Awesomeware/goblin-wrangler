import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "@/App.vue";
import router from "@/router";
import auth from "@/plugins/auth";

createApp(App).use(createPinia()).use(auth).use(router).mount("#app");
