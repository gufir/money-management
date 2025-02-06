<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import Chart from 'primevue/chart'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import axios from 'axios'
import store from '@/store'
import Header from '../Header.vue'
import { useToast } from 'primevue/usetoast'
import Cookies from 'js-cookie'
import Dropdown from 'primevue/dropdown'

const getRandomColor = (): string => {
  const letters = '0123456789ABCDEF'
  let color = '#'
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)]
  }
  return color
}

// Define types
type Dataset = {
  data: number[]
  backgroundColor: string[]
}

type ChartData = {
  labels: string[]
  datasets: Dataset[]
}

const incomeData = ref<ChartData>({ labels: [], datasets: [{ data: [], backgroundColor: [] }] })
const expensesData = ref<ChartData>({ labels: [], datasets: [{ data: [], backgroundColor: [] }] })

const chartOptions = ref({
  plugins: {
    legend: {
      position: 'right',
    },
  },
})

const type = ref<string>('')
const description = ref<string>('')
const amount = ref<string>('')
const categoryName = ref<string>('')
const userId = store.state.user?.user_uuid
const isTransactionDisabled = computed(
  () => !type.value || !description.value || !amount.value || !categoryName.value,
)
const toast = useToast()
const errorMessages = ref<string>('')

const typeOptions = ref([
  { label: 'Expense', value: 'expense' },
  { label: 'Income', value: 'income' },
])

const handleTransaction = async () => {
  try {
    const response = await axios.post(
      'http://localhost:8080/v1/create_transaction',
      {
        amount: amount.value,
        description: description.value,
        type: type.value,
        userId: userId,
        categoryName: categoryName.value,
      },
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${Cookies.get('accessToken')}`,
        },
      },
    )

    toast.add({
      severity: 'success',
      summary: `Transaction ${response.data.id} success`,
      detail: 'Create Transaction Successful',
      life: 3000,
    })

    fetchChartData()
  } catch (error: any) {
    if (error.response && error.response.status === 404) {
      errorMessages.value = error.response.data.message
    } else {
      errorMessages.value = 'An error occurred. Please try again later.'
    }

    toast.add({
      severity: 'error',
      summary: 'Create Transaction Failed',
      detail: errorMessages.value,
      life: 3000,
    })
  }
}

const fetchChartData = async () => {
  try {
    const response = await axios.get('http://localhost:8080/v1/transactions', {
      headers: {
        Authorization: `Bearer ${Cookies.get('accessToken')}`,
      },
    })
    const transactions = response.data
    const income = transactions.filter((t: any) => t.type === 'income')
    const expenses = transactions.filter((t: any) => t.type === 'expense')
    incomeData.value = {
      labels: income.map((t: any) => t.category),
      datasets: [
        {
          data: income.map((t: any) => t.amount),
          backgroundColor: income.map(() => getRandomColor()),
        },
      ],
    }

    expensesData.value = {
      labels: expenses.map((t: any) => t.category),
      datasets: [
        {
          data: expenses.map((t: any) => t.amount),
          backgroundColor: expenses.map(() => getRandomColor()),
        },
      ],
    }
  } catch (error) {
    console.error('Error fetching chart data:', error)
  }
}

onMounted(fetchChartData)
</script>

<template>
  <Header />
  <div class="p-4">
    <div class="flex flex-col md:flex-row justify-between gap-4 p-4">
      <div class="w-full md:w-1/2 p-4 shadow-md rounded bg-white">
        <h2 class="text-md font-semibold">Income Overview</h2>
        <Chart
          type="pie"
          :data="incomeData"
          :options="chartOptions"
          style="max-width: 320px; margin: auto"
        />
      </div>
      <div class="w-full md:w-1/2 p-4 shadow-md rounded bg-white">
        <h2 class="text-md font-semibold">Expenses Overview</h2>
        <Chart
          type="pie"
          :data="expensesData"
          :options="chartOptions"
          style="max-width: 320px; margin: auto"
        />
      </div>
    </div>

    <div class="flex flex-col text-center p-4">
      <h2 class="text-md font-semibold">Add Transaction</h2>
      <InputText v-model="categoryName" placeholder="Category" class="w-full mb-2" />
      <InputText v-model="description" placeholder="Description" class="w-full mb-2" />
      <InputText v-model="amount" placeholder="Amount" type="number" class="w-full mb-2" />
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
        class="p-button-success w-full"
        @click="handleTransaction"
        :disabled="isTransactionDisabled"
      />
    </div>

    <div class="flex flex-col text-center p-4">
      <Button label="View Detailed Report" class="p-button-success" />
    </div>
  </div>
</template>

<style>
body {
  background-color: #f8f9fa;
}
</style>
