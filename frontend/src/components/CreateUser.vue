<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'
import InputGroup from 'primevue/inputgroup'
import InputGroupAddon from 'primevue/inputgroupaddon'
import InputText from 'primevue/inputtext'
import FloatLabel from 'primevue/floatlabel'
import Button from 'primevue/button'
import axios from 'axios'
import router from '@/router'
import {
  validateEmail,
  validateFullName,
  validatePassword,
  validateUsername,
} from '@/utils/validator'

const username = ref('')
const full_name = ref('')
const password = ref('')
const confirmPassword = ref('')
const email = ref('')
const errorMessages = ref<string[]>([])

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

const errors = ref({
  username: '',
  full_name: '',
  email: '',
  password: '',
  confirmPassword: '',
})

const toast = useToast()

// 🔍 Watch setiap field dan validasi real-time
watch(username, (val) => {
  errors.value.username = validateUsername(val) || ''
})
watch(full_name, (val) => {
  errors.value.full_name = validateFullName(val) || ''
})
watch(email, (val) => {
  errors.value.email = validateEmail(val) || ''
})
watch(password, (val) => {
  errors.value.password = validatePassword(val) || ''
  if (confirmPassword.value && val !== confirmPassword.value) {
    errors.value.confirmPassword = 'Passwords do not match'
  } else {
    errors.value.confirmPassword = ''
  }
})
watch(confirmPassword, (val) => {
  if (val !== password.value) {
    errors.value.confirmPassword = 'Passwords do not match'
  } else {
    errors.value.confirmPassword = ''
  }
})

const handleCreateUser = async () => {
  // Validasi input sebelum mengirim permintaan
  // errorMessages.value = []
  // const validations = [
  //   validateUsername(username.value),
  //   validateFullName(full_name.value),
  //   validateEmail(email.value),
  //   validatePassword(password.value),
  // ].filter(Boolean) as string[]

  // if (password.value !== confirmPassword.value) {
  //   validations.push('Passwords do not match.')
  // }

  // if (validations.length > 0) {
  //   errorMessages.value = validations
  //   toast.add({
  //     severity: 'error',
  //     summary: 'Validation Failed',
  //     detail: validations.join('\n'),
  //     life: 4000,
  //   })
  //   return
  // }

  if (isCreateDisabled.value) {
    toast.add({
      severity: 'error',
      summary: 'Validation Failed',
      detail: 'Please fix all validation errors before submitting.',
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
        <!-- Username Field -->
        <div>
          <InputGroup>
            <InputGroupAddon>
              <i class="pi pi-user"></i>
            </InputGroupAddon>
            <FloatLabel>
              <InputText id="username" v-model="username" />
              <label for="username">Username</label>
            </FloatLabel>
          </InputGroup>
          <small v-if="errors.username" class="text-red-500">{{ errors.username }}</small>
        </div>

        <!-- Full Name Field-->
        <div>
          <InputGroup>
            <InputGroupAddon>
              <i class="pi pi-user"></i>
            </InputGroupAddon>
            <FloatLabel>
              <InputText id="fullname" v-model="full_name" />
              <label for="fullname">Full Name</label>
            </FloatLabel>
          </InputGroup>
          <small v-if="errors.full_name" class="text-red-500">{{ errors.full_name }}</small>
        </div>

        <!-- Email Field -->
        <div>
          <InputGroup>
            <InputGroupAddon>
              <i class="pi pi-envelope"></i>
            </InputGroupAddon>
            <FloatLabel>
              <InputText id="email" type="email" v-model="email" />
              <label for="email">Email</label>
            </FloatLabel>
          </InputGroup>
          <small v-if="errors.email" class="text-red-500">{{ errors.email }}</small>
        </div>

        <!-- Password Field -->
        <div>
          <InputGroup>
            <InputGroupAddon>
              <i class="pi pi-lock"></i>
            </InputGroupAddon>
            <FloatLabel>
              <InputText id="password" type="password" v-model="password" />
              <label for="password">Password</label>
            </FloatLabel>
          </InputGroup>
          <small v-if="errors.password" class="text-red-500">{{ errors.password }}</small>
        </div>

        <!-- Confirm Password Field -->
        <div>
          <InputGroup>
            <InputGroupAddon>
              <i class="pi pi-lock"></i>
            </InputGroupAddon>
            <FloatLabel>
              <InputText id="confirm-password" type="password" v-model="confirmPassword" />
              <label for="confirm-password">Confirm Password</label>
            </FloatLabel>
          </InputGroup>
          <small v-if="errors.confirmPassword" class="text-red-500">{{
            errors.confirmPassword
          }}</small>
        </div>

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

.text-red-500 {
  color: #ef4444;
  font-size: 0.85rem;
  margin-top: 4px;
  display: block;
}
</style>
