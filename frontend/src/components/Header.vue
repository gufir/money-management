<script setup lang="ts">
import { ref } from 'vue'
import Cookies from 'js-cookie'
import store from '@/store'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import SidebarMenu from '@/components/SubMenu.vue'
import { defineProps } from 'vue'

const emit = defineEmits<{
  (e: 'updateShowSidebar', value: boolean): void
}>()

const router = useRouter()
const toast = useToast()
const searchQuery = ref('')
const showSidebar = ref(false)

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

const toggleSidebar = () => {
  showSidebar.value = !showSidebar.value
  emit('updateShowSidebar', showSidebar.value)
}
</script>

<template>
  <header
    class="flex justify-between items-center p-4 shadow-md bg-white fixed top-0 left-0 w-full z-50"
  >
    <div class="flex items-center gap-2">
      <Button icon="pi pi-bars" class="p-button-text" @click="toggleSidebar" />
      <img src="../assets/money-symbol.png" alt="MoneyWise Logo" class="h-12" />
      <h1 class="text-2xl font-bold text-gray-700">MoneyWise</h1>
    </div>

    <div class="flex items-center gap-x-4">
      <InputText v-model="searchQuery" placeholder="Search..." class="p-inputtext-sm" />
      <Button label="Generate Report" class="p-button-success" />
      <Button label="Logout" class="p-button-danger" @click="handlerLogout" />
    </div>
    <SidebarMenu v-if="showSidebar" />
  </header>
</template>
