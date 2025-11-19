<script setup lang="ts">
import InputGroup from 'primevue/inputgroup'
import InputGroupAddon from 'primevue/inputgroupaddon'
import InputText from 'primevue/inputtext'
import FloatLabel from 'primevue/floatlabel'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import { ref, computed } from 'vue'
import axios from 'axios'
import router from '@/router'
import Dialog from 'primevue/dialog'

const email = ref<string>('')
const isForgotDisabled = computed(() => !email.value)
const toast = useToast()

const showError = ref(false)
const errorMessage = ref('')

const handlerForgotPassword = async () => {
  // Implementasi logika untuk mengirim email reset password
  try {
    const response = await axios.post(
      'http://localhost:8080/v1/forgot_password',
      {
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
      summary: 'Reset Link Sent',
      detail: 'Please check your email for the password reset link.',
      life: 3000,
    })

    setTimeout(() => {
      router.push('/')
    }, 5000)
  } catch (error: any) {
    console.error('Error sending forgot password request:', error)
    if (error.response) {
      errorMessage.value =
        error.response.data?.message ||
        (error.response.status === 404
          ? 'The provided email does not exist in our records.'
          : 'An error occurred. Please try again later.')
    } else {
      errorMessage.value = 'Network error. Please check your connection.'
    }

    showError.value = true
  }
}
</script>

<template>
  <div class="flex flex-col justify-center items-center min-h-screen">
    <div class="w-full max-w-sm">
      <div class="text-center mb-6">
        <img src="../assets/money-symbol.png" alt="MoneyWise Logo" class="h-12 mx-auto" />
        <h1 class="text-2xl font-bold text-gray-800">MoneyWise</h1>
        <h2 class="text-md font-bold text-gray-600 mt-2">Account Recovery</h2>
      </div>
      <div class="flex flex-col gap-5">
        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-envelope"></i>
          </InputGroupAddon>
          <FloatLabel>
            <InputText id="email" v-model="email" />
            <label for="email">Email</label>
          </FloatLabel>
        </InputGroup>

        <Button
          label="Send Reset Link"
          class="w-full p-button-success hover-button"
          :disabled="isForgotDisabled"
          @click="handlerForgotPassword"
        />
      </div>
    </div>
    <Dialog v-model:visible="showError" header="Request Failed" modal :closable="true" class="w-80">
      <p class="text-red-600 font-semibold">{{ errorMessage }}</p>
    </Dialog>
  </div>
</template>
