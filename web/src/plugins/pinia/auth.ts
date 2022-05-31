import { useAuthStore } from "@/stores/auth";
import { PiniaPluginContext } from "pinia";

export default function authPlugin(context: PiniaPluginContext) {
  var store = useAuthStore();

  if (store.loggedIn) {
    store.getUser();
  }

  return { $auth: store }
};