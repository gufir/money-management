<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Card from 'primevue/card'
import Avatar from 'primevue/avatar'
import Dropdown from 'primevue/dropdown'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import Password from 'primevue/password'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'
import axios from 'axios'
import Cookies from 'js-cookie'
import Dialog from 'primevue/dialog'
import Toast from 'primevue/toast'
import { useToast } from 'primevue/usetoast'
import store from '@/store'

interface User {
  username: string
  full_name: string
  email: string
}
var user = ref<User | null>(null)
const toast = useToast()
const newFullName = ref<string>('')
const newEmail = ref<string>('')

const fetchUserData = async () => {
  try {
    const response = await axios.get('http://localhost:8080/v1/get_user', {
      headers: { Authorization: `Bearer ${Cookies.get('accessToken')}` },
    })

    if (response.data.user.length > 0) {
      const fetchedUser = response.data.user[0]

      user.value = {
        username: fetchedUser.username,
        full_name: fetchedUser.full_name,
        email: fetchedUser.email,
      }
    }
  } catch (error) {
    console.error('Error fetching User:', error)
  }
}

const showDialog = ref(false)

const handleUpdateUser = async () => {
  const dataToSend =
    newFullName.value !== '' && newEmail.value !== ''
      ? { username: store.state.user.username, full_name: newFullName.value, email: newEmail.value }
      : newFullName.value !== ''
        ? { username: store.state.user.username, full_name: newFullName.value }
        : { username: store.state.user.username, email: newEmail.value }
  try {
    await axios.post('http://localhost:8080/v1/update_user', dataToSend, {
      headers: { Authorization: `Bearer ${Cookies.get('accessToken')}` },
    })

    toast.add({ severity: 'success', summary: 'Success', detail: 'User Changed', life: 3000 })
    const resetForm = () => {
      username.value = ''
      email.value = ''
    }
    fetchUserData()
    showDialog.value = false
  } catch (error: any) {
    const errorMsg = error.response?.data?.message || 'An error occurred. Please try again later.'
    toast.add({ severity: 'error', summary: 'Error', detail: errorMsg, life: 3000 })
  }
}

onMounted(() => {
  fetchUserData()
})
</script>
<template>
  <Toast />
  <Card class="mb-4">
    <template #title>User Info</template>
    <template #content>
      <div class="mb-4" v-if="user">
        <label for="username" class="block text-gray-700">Username</label>
        <InputText id="username" v-model="user.username" class="w-full" :disabled="true" />
      </div>
      <div class="mb-4" v-if="user">
        <label for="full_name" class="block text-gray-700">Full Name</label>
        <InputText id="full_name" v-model="user.full_name" class="w-full" :disabled="true" />
      </div>
      <div class="mb-4" v-if="user">
        <label for="email" class="block text-gray-700">Email</label>
        <InputText id="email" v-model="user.email" class="w-full" :disabled="true" />
      </div>

      <Button
        label="Change User Info"
        class="p-button-success hover-button text-white"
        @click="showDialog = true"
      />
    </template>
  </Card>

  <Card class="mb-4">
    <template #title>Security</template>
    <template #content>
      <div class="mb-4">
        <label for="currentPassword" class="block text-gray-700">Current password</label>
        <Password id="currentPassword" v-model="currentPassword" toggleMask class="w-full" />
      </div>
      <div class="mb-4">
        <label for="newPassword" class="block text-gray-700">New password</label>
        <Password id="newPassword" v-model="newPassword" toggleMask class="w-full" />
      </div>
      <div class="flex space-x-2">
        <Button label="Submit" class="p-button-success hover-button text-white" />
        <Button label="Cancel" class="p-button-danger" type="reset" />
      </div>
    </template>
  </Card>

  <Dialog v-model:visible="showDialog" header="Change User Info" modal :style="{ width: '400px' }">
    <div class="flex flex-col items-center gap-2">
      <InputText v-model="newFullName" placeholder="Full Name" class="w-full mb-4" />
      <InputText v-model="newEmail" placeholder="Email" class="w-full mb-4" />
      <Button
        label="Change User"
        class="p-button-success w-full hover-button"
        @click="handleUpdateUser"
        :disabled="isUpdateUserdisabled"
      />
    </div>
  </Dialog>
</template>

<style>
.body {
  background-color: bg-white;
}
</style>
