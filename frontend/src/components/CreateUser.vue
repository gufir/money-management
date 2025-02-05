<script setup lang="ts">
import { ref, computed } from 'vue'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'
import InputGroup from 'primevue/inputgroup'
import InputGroupAddon from 'primevue/inputgroupaddon'
import InputText from 'primevue/inputtext'
import FloatLabel from 'primevue/floatlabel'
import Button from 'primevue/button'
import axios from 'axios'
import router from '@/router'

const username = ref('')
const full_name = ref('')
const password = ref('')
const confirmPassword = ref('')
const email = ref('')
const errorMessages = ref('')

const isCreateDisabled = computed(() => {
  return (
    !username.value ||
    !full_name.value ||
    !password.value ||
    !confirmPassword.value ||
    !email.value ||
    password.value !== confirmPassword.value
  )
})

const toast = useToast()

const handleCreateUser = async () => {
  if (password.value !== confirmPassword.value) {
    errorMessages.value = 'Passwords do not match'
    toast.add({
      severity: 'error',
      summary: 'Create User Failed',
      detail: errorMessages.value,
      life: 3000,
    })
    return
  }

  try {
    await axios.post(
      'http://localhost:8080/v1/create_user',
      {
        username: username.value,
        full_name: full_name.value,
        password: password.value,
        email: email.value,
      },
      {
        headers: {
          'Content-Type': 'application/json',
        },
      },
    )

    toast.add({
      severity: 'success',
      summary: `Hello ${username.value}`,
      detail: 'You have successfully created an account',
      life: 3000,
    })

    // Redirect ke halaman login setelah berhasil daftar
    router.push({ name: 'LoginUser' })
  } catch (error: any) {
    errorMessages.value =
      error.response?.data?.message || 'An error occurred. Please try again later.'
    toast.add({
      severity: 'error',
      summary: 'Registration Failed',
      detail: errorMessages.value,
      life: 3000,
    })
  }
}
</script>

<template>
  <div class="flex flex-col justify-center items-center min-h-screen">
    <div class="w-full max-w-sm">
      <div class="text-center mb-6">
        <img src="../assets/money-symbol.png" alt="MoneyWise Logo" class="h-12 mx-auto" />
        <h1 class="text-2xl font-bold text-gray-800">MoneyWise</h1>
      </div>

      <div class="flex flex-col gap-5">
        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-user"></i>
          </InputGroupAddon>
          <FloatLabel>
            <InputText id="username" v-model="username" />
            <label for="username">Username</label>
          </FloatLabel>
        </InputGroup>

        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-user"></i>
          </InputGroupAddon>
          <FloatLabel>
            <InputText id="fullname" v-model="full_name" />
            <label for="fullname">Full Name</label>
          </FloatLabel>
        </InputGroup>

        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-envelope"></i>
          </InputGroupAddon>
          <FloatLabel>
            <InputText id="email" type="email" v-model="email" />
            <label for="email">Email</label>
          </FloatLabel>
        </InputGroup>

        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-lock"></i>
          </InputGroupAddon>
          <FloatLabel>
            <InputText id="password" type="password" v-model="password" />
            <label for="password">Password</label>
          </FloatLabel>
        </InputGroup>

        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-lock"></i>
          </InputGroupAddon>
          <FloatLabel>
            <InputText id="confirm-password" type="password" v-model="confirmPassword" />
            <label for="confirm-password">Confirm Password</label>
          </FloatLabel>
        </InputGroup>

        <Button
          label="Create User"
          class="w-full p-button-success hover-button"
          :disabled="isCreateDisabled"
          @click="handleCreateUser"
        />

        <Button
          label="Go to Login"
          class="w-full p-button-outlined create-user"
          @click="() => router.push({ name: 'LoginUser' })"
        />
      </div>
    </div>
  </div>
</template>

<style>
body {
  margin: 0;
  font-family: Arial, sans-serif;
}

.hover-button {
  background-color: #27fb2d !important;
  border-color: #27fb2d !important;
  transition:
    background-color 0.3s ease,
    border-color 0.3s ease;
}

.hover-button:hover {
  background-color: #23d628 !important;
  border-color: #23d628 !important;
}
</style>
