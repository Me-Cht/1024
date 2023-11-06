import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-plus'
const app = createApp(App)

app.use(router).use(ElementUI).mount('#app')
