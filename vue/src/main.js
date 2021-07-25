// import Vue from 'vue'
// import App from './App.vue'
// import ElementPlus from 'element-plus'
// import 'element-plus/lib/theme-chalk/index.css';
//
// Vue.config.productionTip = false
//
// new Vue({
//     render: h => h(App),
// }).$mount("#app").use(ElementPlus)
import {createApp} from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import App from './App.vue';

const app = createApp(App)
app.use(ElementPlus)
app.mount('#app')