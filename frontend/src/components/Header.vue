<script setup lang="ts">
import Cookies from 'js-cookie'
import store from '@/store'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import { ref } from 'vue'

const router = useRouter()
const toast = useToast()
const searchQuery = ref('')

const handlerLogout = () => {
  Cookies.remove('accessToken')
  Cookies.remove('refreshToken')

  store.clearUser()

  toast.add({
    severity: 'success',
    summary: 'Logged Out',
    detail: 'You have been successfully logged out.',
    life: 3000,
  })

  router.push('/')
}
</script>

<template>
  <header class="flex justify-between items-center p-4 shadow-md bg-white rounded">
    <div class="flex items-center gap-x-4">
      <img src="../assets/money-symbol.png" alt="MoneyWise Logo" class="h-12" />
      <h1 class="text-2xl font-bold text-gray-700">MoneyWise</h1>
    </div>

    <div class="flex items-center gap-x-4">
      <InputText v-model="searchQuery" placeholder="Search..." class="p-inputtext-sm" />
      <Button label="Generate Report" class="p-button-success" />
      <Button label="Logout" class="p-button-danger" @click="handlerLogout" />
    </div>
  </header>
</template>
