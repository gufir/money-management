<script setup lang="ts">
import LoginForm from '../components/LoginUser.vue'
import Toast from 'primevue/toast'
import UserInfo from '../components/UserInfo.vue';
import store from '../store';
import { useToast } from 'primevue/usetoast';

const toast = useToast();

const onLogout = (user : User) => {
    toast.add({
        severity: 'success',
        summary: `Goodbye ${user.username}`,
        detail: 'You have successfully logged out',
        life: 3000
    });
    store.clearUser();
}
</script>

<template>
  <div>
    <Toast />
    <h1 class="green">Welcome to Money Management</h1>
    <UserInfo v-if="store.state.user" :user="store.state.user" @logout="onLogout"/>
    <LoginForm v-else />
  </div>
</template>
