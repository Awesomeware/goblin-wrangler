import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "@/App.vue";
import router from "@/router";
import auth from "@/plugins/auth";
import vuetify from "@/plugins/vuetify";
import { loadFonts } from "@/plugins/webfontloader";

loadFonts();

createApp(App)
  .use(createPinia())
  .use(auth)
  .use(router)
  .use(vuetify)
  .mount("#app");
