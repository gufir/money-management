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
  category_name: string
  createdAt: string
  userId: string
}

const transactions = ref<Transaction[]>([])

const formatDate = (dateString: string | null | undefined): string => {
  if (!dateString) return 'N/A'
  const date = new Date(dateString) // ISO 8601 support
  return isNaN(date.getTime())
    ? 'Invalid Date'
    : date.toLocaleString('en-GB', {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
      })
}

const fetchTransactions = async () => {
  try {
    const response = await axios.get('http://localhost:8080/v1/get_transaction', {
      headers: { Authorization: `Bearer ${Cookies.get('accessToken')}` },
    })
    console.log(response.data)
    transactions.value = response.data.transaction.map((t: any) => ({
      id: t.id,
      amount: t.amount,
      type: t.type,
      description: t.description,
      category_name: t.category_name,
      created_at: formatDate(t.created_at),
      userId: t.userId,
    }))
  } catch (error) {
    console.error('Failed to fetch transactions:', error)
  }
}

onMounted(fetchTransactions)

defineExpose({ fetchTransactions })
</script>

<template>
  <div class="p-4">
    <h2 class="text-xl font-bold mb-4">Transaction List</h2>
    <DataTable :value="transactions" paginator :rows="10" class="p-datatable-sm">
      <Column field="amount" header="Amount" sortable></Column>
      <Column field="type" header="Type" sortable></Column>
      <Column field="category_name" header="Category Name" sortable></Column>
      <Column field="description" header="Description" sortable></Column>
      <Column field="created_at" header="Date" sortable></Column>
    </DataTable>
  </div>
</template>

<style scoped></style>
