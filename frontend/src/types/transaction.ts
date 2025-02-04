export interface Transaction {
  amount: string
  description: string
  type: 'expense' | 'income'
  userId: string
  categoryName: string
}
