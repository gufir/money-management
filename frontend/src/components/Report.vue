<script setup lang="ts">
import { ref, computed } from 'vue'
import axios from 'axios'
import { useToast } from 'primevue/usetoast'
import { format } from 'date-fns'
import parseISO from 'date-fns/parseISO'
import Cookies from 'js-cookie'
import Header from './Header.vue'
import Message from 'primevue/message'
import Calendar from 'primevue/calendar'
import Button from 'primevue/button'
import ProgressSpinner from 'primevue/progressspinner'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import store from '@/store'

interface ReportItem {
  id: string
  user_id: string
  amount: number
  type: 'income' | 'expense'
  description: string
  created_at: string
}

const reports = ref<ReportItem[]>([])
const startDate = ref<Date | null>(null)
const endDate = ref<Date | null>(null)
const loading = ref<boolean>(false)
const toast = useToast()
const showSidebar = ref(false)
const fetchTimeout = ref<number | null>(null)

// Validasi tanggal
const isDateValid = computed(() => {
  return !startDate.value || !endDate.value || startDate.value <= endDate.value
})

// Hitung total income dan expense
const totalIncome = computed(() =>
  reports.value
    .filter((item) => item.type === 'income')
    .reduce((sum, item) => sum + item.amount, 0),
)

const totalExpense = computed(() =>
  reports.value
    .filter((item) => item.type === 'expense')
    .reduce((sum, item) => sum + item.amount, 0),
)

// Fetch laporan dari API
const fetchReport = async () => {
  if (!isDateValid.value) {
    toast.add({
      severity: 'error',
      summary: 'Invalid Date Range',
      detail: 'End date cannot be earlier than start date.',
      life: 3000,
    })
    return
  }

  loading.value = true

  try {
    let url = 'http://localhost:8080/v1/get_report_by_date'

    const userId = store.state.user?.user_uuid

    if (startDate.value && endDate.value) {
      const formattedStart = format(startDate.value, 'yyyy-MM-dd')
      const formattedEnd = format(endDate.value, 'yyyy-MM-dd')
      url += `?userId=${userId}&startDate=${formattedStart}&endDate=${formattedEnd}`

      console.log(url)
    }

    const response = await axios.get(url, {
      headers: {
        Authorization: `Bearer ${Cookies.get('accessToken')}`,
      },
    })

    reports.value = response.data.reports || []
  } catch (error: any) {
    toast.add({
      severity: 'error',
      summary: 'Failed to Load Report',
      detail: error.response?.data?.message || 'Make sure you have the correct access.',
      life: 3000,
    })
  } finally {
    loading.value = false
  }
}

const handleFilter = () => {
  if (fetchTimeout.value) {
    clearTimeout(fetchTimeout.value)
  }
  fetchTimeout.value = setTimeout(fetchReport, 500) as unknown as number
}
</script>

<template>
  <Toast />
  <Header :showSidebar="showSidebar" @updateShowSidebar="showSidebar = $event" />
  <div :class="['p-4', { 'ml-64': showSidebar, 'ml-0': !showSidebar }]">
    <div class="flex items-center justify-center mt-4">
      <div class="w-full md:flex-row shadow-md rounded bg-white items-center p-4">
        <h2 class="text-2xl font-semibold mb-4">Transaction Report</h2>

        <div class="mb-4 flex gap-3">
          <Calendar v-model="startDate" placeholder="From" dateFormat="yy-mm-dd" />
          <Calendar v-model="endDate" placeholder="Until" dateFormat="yy-mm-dd" />
          <Button
            label="Filter"
            icon="pi pi-search"
            @click="fetchReport"
            :disabled="!isDateValid"
            class="p-button-success hover-button"
          />
        </div>

        <ProgressSpinner v-if="loading" />

        <div v-else-if="reports.length">
          <DataTable :value="reports" paginator :rows="10" class="p-datatable-sm">
            <Column field="created_at" header="Tanggal">
              <template #body="slotProps">
                {{ format(new Date(slotProps.data.created_at), 'yyyy-MM-dd HH:mm:ss') }}
              </template>
            </Column>
            <Column field="description" header="Deskripsi"></Column>
            <Column field="type" header="Tipe">
              <template #body="slotProps">
                <span :class="slotProps.data.type === 'income' ? 'text-[#27fb2d]' : 'text-red-600'">
                  {{ slotProps.data.type.toUpperCase() }}
                </span>
              </template>
            </Column>
            <Column field="amount" header="Jumlah">
              <template #body="slotProps">
                Rp {{ slotProps.data.amount.toLocaleString() }}
              </template>
            </Column>
          </DataTable>
          <div class="mt-4 text-lg font-semibold">
            <p>
              Total Income:
              <span class="text-green-600">Rp {{ totalIncome.toLocaleString() }}</span>
            </p>
            <p>
              Total Expense:
              <span class="text-red-600">Rp {{ totalExpense.toLocaleString() }}</span>
            </p>
          </div>
        </div>

        <Message v-else severity="warn">There were no transactions in this date range.</Message>
      </div>
    </div>
  </div>
</template>

<style scoped>
.p-datatable {
  margin-top: 10px;
}
</style>
