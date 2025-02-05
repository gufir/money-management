import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CreateUser from '../components/CreateUser.vue'
import LoginUser from '../components/LoginUser.vue'
import authStore from '@/store'
import Cookies from 'js-cookie'

const isAuthenticated = () => {
  return !!Cookies.get('accessToken')
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView,
    },
    {
      path: '/create-user',
      name: 'CreateUser',
      component: CreateUser,
      meta: { requiresUnauth: true },
    },
    {
      path: '/login-user',
      name: 'LoginUser',
      component: LoginUser,
      meta: { requiresUnauth: true },
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/pages/NotFound.vue'),
    },

    {
      path: '/dashboard',
      name: 'Dashboard',
      meta: { requiresAuth: true },
      component: () => import('@/components/dashboard/Dashboard.vue'),
    },
    {
      path: '/access-denied',
      name: 'AccessDenied',
      component: () => import('@/views/pages/Unauthorized.vue'),
    },
  ],
})

router.beforeEach((to, from, next) => {
  const isAuth = isAuthenticated()
  if (to.meta.requiresUnauth && isAuth) {
    next('/') // or next('/dashboard')
  } else if (to.meta.requiresAuth && !isAuth) {
    next('/access-denied')
  } else {
    next()
  }
  // if (to.meta.requiresAuth && !isAuthenticated()) {
  //   next('/access-denied')
  // } else {
  //   next()
  // }
})

export default router
