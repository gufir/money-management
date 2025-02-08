<script setup lang="ts">
import LoginForm from '../components/LoginUser.vue'
import Toast from 'primevue/toast'
import { computed } from 'vue'
import Dashboard from '@/components/dashboard/Dashboard.vue'
import store from '../store'
import type { User } from '../types/user'
import { useToast } from 'primevue/usetoast'
import router from '@/router'
import Cookies from 'js-cookie'
import Header from '@/components/Header.vue'
import Transaction from '@/components/Transaction.vue'

const toast = useToast()
const user = computed(() => store.state.user)
const onLogout = (user: User) => {
  toast.add({
    severity: 'success',
    summary: `Goodbye ${user.username}`,
    detail: 'You have successfully logged out',
    life: 3000,
  })
  store.clearUser()
  Cookies.remove('user')
  Cookies.remove('accessToken')
  Cookies.remove('refreshToken')
  router.push('/')
}
</script>

<template>
  <div>
    <Toast />
    <Dashboard v-if="user" :user="user" />
    <LoginForm v-else />
  </div>
</template>
