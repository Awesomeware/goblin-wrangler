import { defineStore } from "pinia";

type ErrorStore = {
  message: string | null;
  errors: {[key: string]: any};
};

export const useErrorStore = defineStore("error", {
  state: () => ({
    message: null,
    errors: {},
  } as ErrorStore),
});
