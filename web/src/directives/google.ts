export default {
  mounted(el: any, binding: any) {
    // @ts-ignore
    google.accounts.id.initialize({
      client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
      callback: binding.value
    });

    // @ts-ignore
    google.accounts.id.renderButton(
      el,
      { theme: 'outline', size: 'large' }
    );
  }
};