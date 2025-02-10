<script setup lang="ts">
import { ref } from 'vue'
import Header from './Header.vue'

const showSidebar = ref(false)
</script>

<template>
  <Header :showSidebar="showSidebar" @updateShowSidebar="showSidebar = $event" />
  <div :class="['p-4', { 'ml-64': showSidebar, 'ml-0': !showSidebar }]">
    <div class="flex items-center justify-center mt-4">
      <div class="w-full md:flex-row shadow-md rounded bg-white items-center p-4"></div>
    </div>
  </div>
</template>
