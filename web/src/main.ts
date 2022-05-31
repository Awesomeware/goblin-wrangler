import { createPinia } from 'pinia'
import { createApp } from 'vue'
import App from './App.vue'
import router from "@/router";
import vuetify from "@/plugins/vue/vuetify";
import { loadFonts } from "@/plugins/vue/webfontloader";
import auth from '@/plugins/pinia/auth';
import googleSSO from '@/directives/google';

loadFonts();

const pinia = createPinia()
  .use(auth)
 
createApp(App)
  .directive('google', googleSSO)
  .use(pinia)
  .use(router)
  .use(vuetify)
  .mount('#app')