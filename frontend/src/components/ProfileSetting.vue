<script setup lang="ts">
import { ref } from 'vue'
import Header from './Header.vue'
import UserSettings from '@/components/settings/UserSettings.vue'

const showSidebar = ref(false)
</script>

<template>
  <Header :showSidebar="showSidebar" @updateShowSidebar="showSidebar = $event" />
  <div :class="['p-4', { 'ml-64': showSidebar, 'ml-0': !showSidebar }]">
    <div class="flex items-center justify-center mt-4">
      <div class="w-full md:flex-row shadow-md rounded bg-white items-center p-4">
        <h2 class="text-xl font-bold mb-4 justify-center">User Settings</h2>
        <UserSettings />
      </div>
    </div>
  </div>
</template>
