//src/router/router.js
import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/pages/index/index'
import Setting from '@/pages/setting/index'

Vue.use(Router)

let router = new Router({
  mode: '',
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index,
      meta:{
        title: '主页'
      }
    },
    {
      path: '/setting',
      name: 'Setting',
      component: Setting,
      meta:{
        title: '设置'
      }
    },
    ]
})


export default router

