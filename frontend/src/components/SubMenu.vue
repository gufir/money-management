<script setup lang="ts">
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const menuItems = ref([
  { label: 'Dashboard', icon: 'pi pi-home', subMenus: [], path: '/dashboard', name: 'Dashboard' },
  {
    label: 'Transactions',
    icon: 'pi pi-wallet',
    path: '/transaction',
    subMenus: ['View Transactions', 'Add Transaction'],
    name: 'Transaction',
  },
  {
    label: 'Reports',
    icon: 'pi pi-chart-line',
    path: '/reports',
    subMenus: ['Monthly Report', 'Yearly Report'],
    name: 'Reports',
  },
  {
    label: 'Settings',
    icon: 'pi pi-cog',
    path: '/settings',
    subMenus: ['Profile', 'Security', 'Notifications'],
    name: 'Settings',
  },
])

const openSubMenu = ref<string[]>([])
const route = useRoute()
const router = useRouter()

const toggleSubMenu = (subMenu: string[]) => {
  if (openSubMenu.value.includes(subMenu[0])) {
    openSubMenu.value = openSubMenu.value.filter((item) => item !== subMenu[0])
  } else {
    openSubMenu.value.push(subMenu[0])
  }
}

const isActive = (path: string, name: string, label: string) => {
  if (!path) {
    return route.name === name || openSubMenu.value.includes(label)
  }
  return route.path === path || (path !== '/' && route.path.startsWith(path))
}

const handleMenuClick = (item: { path: string; label: string }) => {
  if (!item.path) {
    router.push({ name: 'NotFound' })
  } else {
    router.push(item.path)
  }
}
</script>

<template>
  <div class="w-64 h-screen bg-white shadow-md fixed left-5 top-29.5 pt-0 z-40">
    <ul class="space-y-2 p-4">
      <li
        v-for="(item, index) in menuItems"
        :key="index"
        class="flex flex-col gap-2 p-2 hover:bg-gray-200 rounded cursor-pointer"
        :class="[
          'flex flex-col gap-2 p-2 rounded cursor-pointer',
          { 'bg-gray-200': isActive(item.path, item.name, item.label) },
        ]"
      >
        <div @click="handleMenuClick(item)" class="flex items-center gap-2">
          <i :class="item.icon"></i>
          <span>{{ item.label }}</span>
        </div>

        <ul v-if="openSubMenu.includes(item.label)" class="space-y-1 pl-6 mt-2">
          <li
            v-for="(subItem, subIndex) in item.subMenus"
            :key="subIndex"
            class="p-2 hover:bg-gray-100 rounded cursor-pointer"
          >
            {{ subItem }}
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<style scoped>
ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
</style>
