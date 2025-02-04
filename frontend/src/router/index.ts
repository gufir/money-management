import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CreateUser from '../components/CreateUser.vue'
import LoginUser from '../components/LoginUser.vue'

const isAuthenticated = () => {
  // Replace this with your actual authentication logic (e.g., check a token or session)
  return localStorage.getItem('user') // Example check
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
    },
    {
      path: '/login-user',
      name: 'LoginUser',
      component: LoginUser,
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
      component: () => import('@/components/dashboard/header.vue'),
    },
    {
      path: '/access-denied',
      name: 'AccessDenied',
      component: () => import('@/views/pages/Unauthorized.vue'),
    },
  ],
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !isAuthenticated()) {
    next('/access-denied')
  } else {
    next()
  }
})

export default router
