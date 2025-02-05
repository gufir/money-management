<script setup lang="ts">
import { ref } from 'vue'
import Chart from 'primevue/chart'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import axios from 'axios'
import store from '@/store'
import Header from '../Header.vue'

const getRandomColor = (): string => {
  const letters = '0123456789ABCDEF'
  let color = '#'
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)]
  }
  return color
}

type Dataset = {
  data: number[]
  backgroundColor: string[]
}

type ChartData = {
  labels: string[]
  datasets: Dataset[]
}

type Income = {
  category: string
  amount: string
  description: string
}

type Expense = {
  category: string
  amount: string
}

const incomeData = ref<ChartData>({
  labels: [],
  datasets: [
    {
      data: [],
      backgroundColor: [],
    },
  ],
})

const expensesData = ref<ChartData>({
  labels: [],
  datasets: [
    {
      data: [],
      backgroundColor: [],
    },
  ],
})

const chartOptions = ref({
  plugins: {
    legend: {
      position: 'right',
    },
  },
})

const newIncome = ref<Income>({ category: '', amount: '', description: '' })
const newExpense = ref<Expense>({ category: '', amount: '' })

const addIncome = async () => {
  if (newIncome.value.category && newIncome.value.amount) {
    const category = newIncome.value.category
    const amount = parseFloat(newIncome.value.amount)
    const description = newIncome.value.description

    // Update local state
    incomeData.value.labels.push(category)
    incomeData.value.datasets[0].data.push(amount)
    incomeData.value.datasets[0].backgroundColor.push(getRandomColor())

    try {
      await axios.post(
        'http://localhost:8080/v1/create_transaction',
        {
          userId: store.state.user?.user_uuid,
          type: 'income',
          categoryName: category,
          amount: amount,
          description: description,
        },
        {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${store.state.accessToken}`,
          },
        },
      )
      console.log('Income added successfully')
    } catch (error) {
      console.error('Failed to add income:', error)
    }

    // Reset form
    newIncome.value.category = ''
    newIncome.value.amount = ''
    newIncome.value.description = ''
  }
}

const addExpense = async () => {
  if (newExpense.value.category && newExpense.value.amount) {
    const category = newExpense.value.category
    const amount = parseFloat(newExpense.value.amount)

    // Update local state
    expensesData.value.labels.push(category)
    expensesData.value.datasets[0].data.push(amount)
    expensesData.value.datasets[0].backgroundColor.push(getRandomColor())

    try {
      await axios.post(
        'http://localhost:8080/v1/create_transaction',
        {
          userId: store.state.user?.user_uuid,
          type: 'expense',
          categoryName: category,
          amount: amount,
        },
        {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${store.state.accessToken}`,
          },
        },
      )
      console.log('Expense added successfully')
    } catch (error) {
      console.error('Failed to add expense:', error)
    }

    // Reset form
    newExpense.value.category = ''
    newExpense.value.amount = ''
  }
}
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

    <div class="flex flex-col md:flex-row justify-center gap-4 p-4">
      <div class="w-full md:max-w-md p-4 shadow-md rounded bg-white mx-auto">
        <h2 class="text-md font-semibold">Add New Income</h2>
        <InputText v-model="newIncome.category" placeholder="Category" class="w-full mb-2" />
        <InputText v-model="newIncome.description" placeholder="Description" class="w-full mb-2" />
        <InputText
          v-model="newIncome.amount"
          placeholder="Amount"
          type="number"
          class="w-full mb-2"
        />
        <Button label="Add Income" class="p-button-success w-full" @click="addIncome" />
      </div>
      <div class="w-full md:max-w-md p-4 shadow-md rounded bg-white mx-auto">
        <h2 class="text-md font-semibold">Add New Expense</h2>
        <InputText v-model="newExpense.category" placeholder="Category" class="w-full mb-2" />
        <InputText
          v-model="newExpense.amount"
          placeholder="Amount"
          type="number"
          class="w-full mb-2"
        />
        <Button label="Add Expense" class="p-button-success w-full" @click="addExpense" />
      </div>
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
