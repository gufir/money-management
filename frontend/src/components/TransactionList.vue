<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Cookies from 'js-cookie'

interface Transaction {
  id: string
  amount: string
  type: string
  description: string
  category_id: string
  user_id: string
}

const transactions = ref<Transaction[]>([])

const fetchTransactions = async () => {
  try {
    const response = await axios.get('http://localhost:8080/v1/get_transaction', {
      headers: { Authorization: `Bearer ${Cookies.get('accessToken')}` },
    })
    transactions.value = response.data.transaction
  } catch (error) {
    console.error('Failed to fetch transactions:', error)
  }
}

onMounted(fetchTransactions)
</script>

<template>
  <div class="p-4">
    <h2 class="text-xl font-bold mb-4">Transaction List</h2>
    <DataTable :value="transactions" paginator :rows="10" class="p-datatable-sm">
      <Column field="amount" header="Amount"></Column>
      <Column field="type" header="Type"></Column>
      <Column field="description" header="Description"></Column>
      <Column field="category_id" header="Category ID"></Column>
      <Column field="user_id" header="User ID"></Column>
    </DataTable>
  </div>
</template>

<style scoped>
.p-datatable {
  width: 100%;
}
</style>
