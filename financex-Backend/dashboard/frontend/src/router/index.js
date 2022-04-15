import Vue from 'vue'
import VueRouter from 'vue-router'
import BillSplitList from '@/components/BillSplitList'
import BillSplit from '@/components/BillSplit'
import Balance from '@/components/Balance'
import NewExpense from '@/components/NewExpense'
import Expense from '@/components/Expense'
import NewBillsplit from '@/components/NewBillsplit'
import Participants from '@/components/Participants'

Vue.use(VueRouter)

export default new VueRouter({
    mode: 'history',
    routes: [
     {
        path: '/',
        name: 'BillSplitList',
        component: BillSplitList,
        meta: {
          reload: true,
        },
      },
      {
        path: '/billSplit/:id',
        name: 'BillSplit',
        props: true,
        component: BillSplit,
        meta: {
          reload: true,
        },
      },
      {
        path: '/billsplit/:id/balance',
        name: 'Balance',
        props: true,
        component: Balance,
      },
      {
        path: '/billsplit/:id/new',
        name: 'NewExpense',
        props: true,
        component: NewExpense
      },
      {
        path: '/billsplit/:id/expense/:expense_id',
        name: 'Expense',
        props: true,
        component: Expense
      },
      {
        path: '/new',
        name: 'NewBillsplit',
        component: NewBillsplit
      },
      {
        path: '/billsplit/:id/participants',
        name: 'Paticipants',
        component: Participants,
        props: true,
      }

    ]
  })