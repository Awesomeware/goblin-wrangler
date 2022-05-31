import { createPinia } from 'pinia'
import { createApp } from 'vue'
import App from './App.vue'
import router from "@/router";
import vuetify from "@/plugins/vue/vuetify";
import { loadFonts } from "@/plugins/vue/webfontloader";
import auth from '@/plugins/pinia/auth';

loadFonts();

const pinia = createPinia()
  .use(auth)
 
createApp(App)
  .use(pinia)
  .use(router)
  .use(vuetify)
  .mount('#app')