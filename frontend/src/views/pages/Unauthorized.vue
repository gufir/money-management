<script setup lang="ts">
import Button from 'primevue/button'
</script>

<template>
  <div class="flex items-center justify-center min-h-screen overflow-hidden">
    <div class="flex flex-col items-center justify-center">
      <div>
        <div
          class="w-full bg-surface-0 dark:bg-surface-900 py-20 px-8 sm:px-20 flex flex-col items-center"
        >
          <span class="text-[#26fb2d] font-bold text-9xl">403</span>
          <h1 class="text-surface-900 dark:text-surface-0 font-bold text-3xl lg:text-7xl mb-2">
            Access Denied
          </h1>
          <div class="text-surface-600 dark:text-surface-200 mb-8">
            You do not have the necessary permisions. Please contact admins.
          </div>
          <RouterLink to="/" custom v-slot="{ navigate }">
            <Button @click="navigate" label="Go to Home" class="p-button-success hover-button" />
          </RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
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
