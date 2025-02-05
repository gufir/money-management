import { createApp } from 'vue'
import './assets/main.css'
import App from './App.vue'
import Primvue from 'primevue/config'
import router from './router'
import ToastService from 'primevue/toastservice'

const app = createApp(App)

app.use(router)
app.use(Primvue)
app.use(ToastService)

app.mount('#app')
