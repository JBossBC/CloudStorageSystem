import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import 'element-plus/dist/index.css'
import Vuex, {createStore} from "vuex";
import "./assets/main.css";
import ElementPlus  from "element-plus"
import axiosApi from "@/router/GlobalAxios";
const app =createApp(App)
app.use(ElementPlus)
app.config.globalProperties.$axios=axiosApi
app.use(router)
app.mount('#app')
const store = createStore({
    state() {
        return {
            userName: "",
        }
    },
    mutations: {
        setUserName(state, username) {
            state.userName = username;
        }
    }
})
app.use(store);
