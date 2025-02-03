import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CreateUser from '../components/CreateUser.vue'
import LoginUser from '../components/LoginUser.vue'

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
      component: LoginUser,
    },
  ],
})

export default router
