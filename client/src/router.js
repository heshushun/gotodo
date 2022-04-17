import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Health from "./views/Health";

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/health',
      name: 'health',
      component: Health
    },
    {
      path: '/about',
      name: 'about',
      component: function () {
        // 这会为此路由生成一个单独的块 (about.[hash].js)
        // 访问路由时延迟加载。
        return import(/* webpackChunkName: "about" */ './views/About.vue')
      }
    }
  ]
})
