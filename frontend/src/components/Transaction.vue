<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import axios from 'axios'
import store from '@/store'
import Header from './Header.vue'
import { useToast } from 'primevue/usetoast'
import Cookies from 'js-cookie'
import Dropdown from 'primevue/dropdown'
import InputNumber from 'primevue/inputnumber'
import Toast from 'primevue/toast'
import TransactionList from './TransactionList.vue'

const showSidebar = ref(false)
const showDialog = ref(false)
const transactionListRef = ref<any>(null)

const type = ref<string>('')
const description = ref<string>('')
const amount = ref<number>()
const categoryName = ref<string>('')
const categories = ref<{ label: string; value: string }[]>([])
const userId = store.state.user?.user_uuid
const transactions = ref<any[]>([])
const isTransactionDisabled = computed(
  () => !type.value || !description.value || !amount.value || !categoryName.value,
)
const typeOptions = ref([
  { label: 'Expense', value: 'expense' },
  { label: 'Income', value: 'income' },
])
const toast = useToast()

const fetchCategories = async () => {
  try {
    const response = await axios.get('http://localhost:8080/v1/get_categories', {
      headers: { Authorization: `Bearer ${Cookies.get('accessToken')}` },
    })
    categories.value = response.data.categories.map((category: any) => ({
      label: category.name,
      value: category.id,
    }))
  } catch (error) {
    console.error('Error fetching categories:', error)
  }
}

const handleTransaction = async () => {
  try {
    await axios.post(
      'http://localhost:8080/v1/create_transaction',
      {
        amount: amount.value,
        description: description.value,
        type: type.value,
        userId: userId,
        categoryId: categoryName.value,
      },
      {
        headers: { Authorization: `Bearer ${Cookies.get('accessToken')}` },
      },
    )

    toast.add({ severity: 'success', summary: 'Success', detail: 'Transaction added', life: 3000 })
    resetForm()
    showDialog.value = false
    if (
      transactionListRef.value &&
      typeof transactionListRef.value.fetchTransactions === 'function'
    ) {
      await transactionListRef.value.fetchTransactions()
    }
  } catch (error: any) {
    const errorMsg = error.response?.data?.message || 'An error occurred. Please try again later.'
    toast.add({ severity: 'error', summary: 'Error', detail: errorMsg, life: 3000 })
  }
}

const resetForm = () => {
  type.value = ''
  description.value = ''
  amount.value = undefined
  categoryName.value = ''
}

onMounted(async () => {
  await fetchCategories()
})
</script>

<template>
  <Toast />
  <Header :showSidebar="showSidebar" @updateShowSidebar="showSidebar = $event" />

  <div :class="['p-4', { 'ml-64': showSidebar, 'ml-0': !showSidebar }]">
    <div class="flex items-center justify-center mt-4">
      <div class="w-full md:flex-row shadow-md rounded bg-white items-center p-4">
        <TransactionList ref="transactionListRef" />
        <div class="flex flex-col items-center">
          <Button
            label="Add Transaction"
            class="p-button-success hover-button"
            @click="showDialog = true"
          />
        </div>
      </div>
    </div>
  </div>

  <Dialog v-model:visible="showDialog" header="Add Transaction" modal :style="{ width: '400px' }">
    <div class="flex flex-col items-center gap-2">
      <Dropdown
        v-model="categoryName"
        :options="categories"
        optionLabel="label"
        optionValue="value"
        placeholder="Select Category"
        class="w-full mb-2"
      />
      <InputText v-model="description" placeholder="Description" class="w-full" />
      <InputNumber v-model="amount" placeholder="Amount" class="w-full mb-2" />
      <Dropdown
        v-model="type"
        :options="typeOptions"
        optionLabel="label"
        optionValue="value"
        placeholder="Select Type"
        class="w-full mb-2"
      />
      <Button
        label="Add Transaction"
        class="p-button-success w-full hover-button"
        @click="handleTransaction"
        :disabled="isTransactionDisabled"
      />
    </div>
  </Dialog>
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
