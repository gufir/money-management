<script setup lang="ts">
import { ref } from 'vue'
import Chart from 'primevue/chart'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import axios from 'axios'

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
}

type Expense = {
  category: string
  amount: string
}

const incomeData = ref<ChartData>({
  labels: ['Salary', 'Investments', 'Freelance', 'Rental', 'Other'],
  datasets: [
    {
      data: [3000, 1200, 800, 500, 300],
      backgroundColor: ['#4CAF50', '#66BB6A', '#81C784', '#A5D6A7', '#C8E6C9'],
    },
  ],
})

const expensesData = ref<ChartData>({
  labels: ['Rent', 'Food', 'Transport', 'Entertainment', 'Other'],
  datasets: [
    {
      data: [1000, 500, 200, 300, 150],
      backgroundColor: ['#FF5733', '#FF7043', '#FF8A65', '#FFAB91', '#FFC1A6'],
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

const newIncome = ref<Income>({ category: '', amount: '' })
const newExpense = ref<Expense>({ category: '', amount: '' })

const addIncome = async () => {
  if (newIncome.value.category && newIncome.value.amount) {
    incomeData.value.labels.push(newIncome.value.category)
    incomeData.value.datasets[0].data.push(parseFloat(newIncome.value.amount))

    try {
      await axios.post('/v1/create_transaction', {
        type: 'income',
        category: newIncome.value.category,
        amount: parseFloat(newIncome.value.amount),
      })
      console.log('Income added successfully')
    } catch (error) {
      console.error('Failed to add income:', error)
    }

    newIncome.value.category = ''
    newIncome.value.amount = ''
  }
}

const addExpense = async () => {
  if (newExpense.value.category && newExpense.value.amount) {
    expensesData.value.labels.push(newExpense.value.category)
    expensesData.value.datasets[0].data.push(parseFloat(newExpense.value.amount))

    // Send data to server
    try {
      await axios.post('/v1/create_transaction', {
        type: 'expense',
        category: newExpense.value.category,
        amount: parseFloat(newExpense.value.amount),
      })
      console.log('Expense added successfully')
    } catch (error) {
      console.error('Failed to add expense:', error)
    }

    newExpense.value.category = ''
    newExpense.value.amount = ''
  }
}
</script>

<template>
  <div class="p-4">
    <!-- Header -->
    <header class="flex justify-between items-center p-4 shadow-md bg-white rounded">
      <div class="flex items-center gap-x-4">
        <img src="../../assets/money-symbol.png" alt="MoneyWise Logo" class="h-12" />
        <h1 class="text-2xl font-bold text-gray-700">MoneyWise</h1>
      </div>
      <Button label="Generate Report" class="p-button-success" />
    </header>

    <!-- Pie Chart Section -->
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

    <!-- Form Section -->
    <div class="flex flex-col md:flex-row justify-center gap-4 p-4">
      <div class="w-full md:max-w-md p-4 shadow-md rounded bg-white mx-auto">
        <h2 class="text-md font-semibold">Add New Income</h2>
        <InputText v-model="newIncome.category" placeholder="Category" class="w-full mb-2" />
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

    <!-- View Report Button -->
    <div class="text-center p-4">
      <Button label="View Detailed Report" class="p-button-success" />
    </div>
  </div>
</template>

<style>
body {
  background-color: #f8f9fa;
}
</style>
