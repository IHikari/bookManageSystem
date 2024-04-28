import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '../stores'
// createRouter 创建路由实例
// 配置history模式
// 1. history模式:createWebHistory  地址栏不带#
// 2. hash模式：createWebHashHistory 地址栏带#

// import.meta.env.BASE_URL: vite中的环境变量，可在vite.config.js中通过 base配置项配置
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/login', component: () => import('../views/login/LoginPage.vue') }, //登录页
    {
      path: '/',
      component: () => import('../views/layout/LayoutContainer.vue'),
      redirect: '/list',
      children: [
        {
          path: '/list',
          component: () => import('../views/book/BookList.vue')
        },
        {
          path: '/book/manage',
          component: () => import('../views/book/BookManage.vue')
        },
        {
          path: '/book/channel',
          component: () => import('../views/book/BookChannel.vue')
        },
        {
          path: '/book/have',
          component: () => import('../views/book/BookHave.vue')
        },
        {
          path: '/user/profile',
          component: () => import('../views/user/UserProfile.vue')
        },
        {
          path: '/user/avatar',
          component: () => import('../views/user/UserAvatar.vue')
        },
        {
          path: '/user/password',
          component: () => import('../views/user/UserPassword.vue')
        },
        {
          path: '/user/manage',
          component: () => import('../views/user/UserManage.vue')
        },
        {
          path: '/record',
          component: () => import('../views/record/RecordBook.vue')
        }
      ]
    }
  ]
})

// 登录访问拦截 ===> 默认直接放行
// 根据返回值决定，是放行还是拦截
// 返回值:
// 1.undefined/true 直接放行
// 2.false 拦回from的地址页面
// 3.具体路径或路径对象 拦截到对应的地址
//       '/login'    {name:'/login' }
router.beforeEach((to) => {
  // 如果没有token且访问的是非登录页,拦截到登录，其他情况正常放行
  const userStore = useUserStore()
  if (!userStore.token && to.path !== '/login') return '/login'
  return true
})

export default router
