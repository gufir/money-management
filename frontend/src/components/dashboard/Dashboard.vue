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
  return `#${Math.floor(Math.random() * 16777215).toString(16)}`
}

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
const categories = ref<{ label: string; value: string }[]>([])
const userId = store.state.user?.user_uuid
const transactions = ref<any[]>([])
const isTransactionDisabled = computed(
  () => !type.value || !description.value || !amount.value || !categoryName.value,
)
const toast = useToast()

const typeOptions = ref([
  { label: 'Expense', value: 'expense' },
  { label: 'Income', value: 'income' },
])

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

const fetchTransactions = async () => {
  try {
    const response = await axios.get('http://localhost:8080/v1/get_transaction', {
      headers: { Authorization: `Bearer ${Cookies.get('accessToken')}` },
    })
    transactions.value = response.data.transaction || []
    updateChartData()
  } catch (error) {
    console.error('Error fetching transactions:', error)
  }
}

const getCategoryName = (categoryId: string): string => {
  const category = categories.value.find((c) => c.value === categoryId)
  return category ? category.label : categoryId // Gunakan ID jika kategori belum tersedia
}

const categoryColors = ref<Record<string, string>>({})

const getCategoryColor = (categoryId: string): string => {
  if (!categoryColors.value[categoryId]) {
    categoryColors.value[categoryId] = getRandomColor()
  }
  return categoryColors.value[categoryId]
}
const updateChartData = () => {
  if (!transactions.value.length) return

  const incomeMap = new Map<string, number>()
  const expensesMap = new Map<string, number>()

  transactions.value.forEach((t) => {
    const categoryName = getCategoryName(t.category_id)
    const amount = parseFloat(t.amount)

    if (categoryName) {
      if (t.type === 'income') {
        incomeMap.set(categoryName, (incomeMap.get(categoryName) || 0) + amount)
      } else if (t.type === 'expense') {
        expensesMap.set(categoryName, (expensesMap.get(categoryName) || 0) + amount)
      }
    }
  })

  const getChartData = (dataMap: Map<string, number>) => {
    const labels = Array.from(dataMap.keys())
    const data = Array.from(dataMap.values())
    const backgroundColor = labels.map((label) => {
      const categoryId = categories.value.find((c) => c.label === label)?.value || label
      return getCategoryColor(categoryId)
    })

    return { labels, datasets: [{ data, backgroundColor }] }
  }

  incomeData.value = getChartData(incomeMap)
  expensesData.value = getChartData(expensesMap)
}

const handleTransaction = async () => {
  try {
    await axios.post(
      'http://localhost:8080/v1/create_transaction',
      {
        amount: parseFloat(amount.value),
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
    fetchTransactions()
  } catch (error: any) {
    const errorMsg = error.response?.data?.message || 'An error occurred. Please try again later.'
    toast.add({ severity: 'error', summary: 'Error', detail: errorMsg, life: 3000 })
  }
}

onMounted(async () => {
  await fetchCategories()
  await fetchTransactions()
})
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

    <div class="mt-6 text-center">
      <h2 class="text-md font-semibold">Add Transaction</h2>
      <Dropdown
        v-model="categoryName"
        :options="categories"
        optionLabel="label"
        optionValue="value"
        placeholder="Select Category"
        class="w-full mb-2"
      />
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
  </div>
</template>

<style>
body {
  background-color: #f8f9fa;
}
</style>
