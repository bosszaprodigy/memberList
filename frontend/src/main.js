import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import VueSweetalert2 from 'vue-sweetalert2'
import 'sweetalert2/dist/sweetalert2.min.css'
import '@mdi/font/css/materialdesignicons.css'

const app = createApp(App)
app.use(VueSweetalert2)
app.use(router)
app.use(vuetify)
app.mount('#app')
