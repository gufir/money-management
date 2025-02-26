<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Button from 'primevue/button';

interface VerifyEmailResponse {
  is_verified: boolean;
  message?: string;
}

const route = useRoute();
const router = useRouter();

const emailId = ref<string | null>(null);
const secretCode = ref<string | null>(null);
const statusMessage = ref<string>('Verifying email...');
const isLoading = ref<boolean>(true);

onMounted(async () => {
  emailId.value = route.query.email_id as string | null;
  secretCode.value = route.query.secret_code as string | null;

  if (!emailId.value || !secretCode.value) {
    statusMessage.value = 'Invalid verification link.';
    isLoading.value = false;
    return;
  }

  try {
    const response = await fetch(
      `http://localhost:8080/v1/verify_email?email_id=${emailId.value}&secret_code=${secretCode.value}`
    );
    const data: VerifyEmailResponse = await response.json();

    if (response.ok && data.is_verified) {
      statusMessage.value = '✅ Email verified successfully! Redirecting to login...';
      setTimeout(() => {
        router.push('/login-user');
      }, 6000);
    } else {
      statusMessage.value = data.message || '❌ Verification failed. Please check your email or request a new link.';
    }
  } catch (error) {
    statusMessage.value = '⚠️ Error verifying email. Please try again later.';
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <div class="flex items-center justify-center min-h-screen bg-gray">
    <div class="bg-white p-6 rounded-lg shadow-lg max-w-md text-center w-full">
      <h1 class="text-2xl font-semibold text-gray-800 mb-4">Email Verification</h1>

      <div v-if="isLoading" class="flex items-center justify-center">
        <div class="w-6 h-6 border-4 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
      </div>

      <p v-else class="text-gray-700">{{ statusMessage }}</p>

      <Button
        label="Go to Login"
        v-if="!isLoading && statusMessage.includes('successfully')"
        @click="router.push('/login-user')"
        class="p-button-success w-1/2 hover-button mt-4"
      />
    </div>
  </div>
</template>

<style scoped>
</style>
