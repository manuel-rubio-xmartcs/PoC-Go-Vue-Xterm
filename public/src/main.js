const { createApp } = Vue;
const { createVuetify } = Vuetify;
const { loadModule } = window['vue3-sfc-loader'];
const { App } = './App.vue';

const vuetify = createVuetify({
    theme: { defaultTheme: 'dark' },
})

const options = {
    moduleCache: {
      vue: Vue,
    },
    getFile(url) {
      return fetch(url).then((resp) =>
        resp.ok ? resp.text() : Promise.reject(resp)
      );
    },
    addStyle(styleStr) {
      const style = document.createElement('style');
      style.textContent = styleStr;
      const ref = document.head.getElementsByTagName('style')[0] || null;
      document.head.insertBefore(style, ref);
    },
    log(type, ...args) {
      console.log(type, ...args);
    },
  };

const app = createApp({
    components: {
        App: Vue.defineAsyncComponent(() =>
            loadModule('./public/src/App.vue', options)
        ),
    }
})

app.use(vuetify).mount('#app')