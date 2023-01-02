import {createRouter, createWebHistory} from 'vue-router'
import CookieUtil from "@/router/RouteUtil";
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/Login',
            name: 'Login',
            component: () => import('../views/Login.vue')
        },
        {
            path: "/",
            name: "index",
            component: () => import('../views/Login.vue')
        },
        {
            path: '/home',
            name: 'home',
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('../views/Home.vue')
        }
    ]
})

router.beforeEach(function (to, from, next) {
    let cookie = CookieUtil.getCookie("isLogin");
    if (to.path === "/Login") {
        if (cookie) {
            return next("/Home");
        } else {
            return next();
        }
    } else {
        if (cookie) {
            return next();
        } else {
            return next("/Login");
        }
    }
})

export default router