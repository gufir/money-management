<script setup lang="ts">
import InputGroup from 'primevue/inputgroup'
import InputGroupAddon from 'primevue/inputgroupaddon'
import InputText from 'primevue/inputtext'
import FloatLabel from 'primevue/floatlabel'
import { ref, computed } from 'vue'
import Button from 'primevue/button'
import axios from 'axios'
import { useToast } from 'primevue/usetoast'
import router from '@/router'

const username = ref<string>('')
const password = ref<string>('')
const isLoginDisabled = computed(() => !username.value || !password.value)
const toast = useToast()
const errorMessages = ref<string>('')

const handlerLogin = async () => {
  try {
    const response = await axios.post(
      'http://localhost:8080/v1/login_user',
      {
        username: username.value,
        password: password.value,
      },
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'none',
        },
      },
    )

    toast.add({
      severity: 'success',
      summary: `Hello ${response.data.user.username}`,
      detail: 'You have successfully logged in',
      life: 3000,
    })
  } catch (error: any) {
    if (error.response && error.response.status === 404) {
      errorMessages.value = error.response.data.message
    } else {
      errorMessages.value = 'An error occurred. Please try again later.'
    }

    toast.add({
      severity: 'error',
      summary: 'Login Failed',
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
            <i class="pi pi-lock"></i>
          </InputGroupAddon>
          <FloatLabel>
            <InputText id="password" type="password" v-model="password" />
            <label for="password">Password</label>
          </FloatLabel>
        </InputGroup>

        <Button
          label="Login"
          class="w-full p-button-success hover-button"
          :disabled="isLoginDisabled"
          @click="handlerLogin"
        />

        <div class="text-center">
          <a href="#" class="text-sm forgot-password">Forgot Password?</a>
        </div>

        <Button
          label="Create User"
          class="w-full p-button-outlined create-user"
          @click="() => router.push('/create-user')"
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
.forgot-password {
  color: #27fb2d;
  text-decoration: none;
}

.forgot-password:hover {
  text-decoration: underline;
}

.create-user {
  color: #27fb2d !important;
}
</style>
