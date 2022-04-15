<template>

  <v-container v-if="loading">
    <div class="text-xs-center">
        <v-progress-circular
          indeterminate
          :size="150"
          :width="8"
          color="green">
        </v-progress-circular>
      </div>
  </v-container>

  <v-container v-else>
    <v-layout row wrap>
      <v-flex xs12 mr-1 ml-1>


        <v-card>
          <v-card-title primary-title>
            <div>
              <h2 class="headline mb-0">{{this.BillSplit.Name}}</h2>
            </div>
            <v-card-subtitle>
              Created At {{this.BillSplit.CreatedAt}}
              </v-card-subtitle>
          </v-card-title>
          <v-card-actions>
            <v-btn text color="green" @click="back">back</v-btn>
          </v-card-actions>

            <v-card-actions>
            <v-btn @click="newExpense"> new expense</v-btn> <v-btn @click="manageParticipants"> manage participants</v-btn>
          </v-card-actions>

        </v-card>

      </v-flex>
    </v-layout>


    <v-layout row wrap>
      <v-flex xs12>
        <div class="text-xs-center">
          <v-card>
            <v-subheader>Participants</v-subheader>
            <v-row flat>
              <v-col v-for="(item, index) in this.Participants" :key="index">
              <v-card>{{ item.Name }}</v-card>
              </v-col>
            </v-row>
          </v-card>
        </div>
      </v-flex>
    </v-layout>


    <v-layout row wrap>
      <v-flex xs12>
        <div class="text-center">
          <v-card>
             <v-list flat>
                <v-subheader>Expenses</v-subheader>
                <v-list-item
                v-for="item in Expenses"
                :key="item.Uuid"
                >

                <v-list-item-content>
                    <v-list-item-title  v-text="item.Name"></v-list-item-title>
                    <v-list-item-subtitle>   Payed By: <b>{{item.PayerName}}</b>   Amount: <b>{{parseFloat(item.Amount).toFixed( 2 )}}€</b>    Payed at: <b>{{item.CreatedAt}}</b>      </v-list-item-subtitle>
                </v-list-item-content>

                <v-list-item-action>
                  <v-btn @click="getExpense(item.Uuid)" icon>
                    <v-icon color="grey lighten-1">mdi-information</v-icon>
                  </v-btn>
                </v-list-item-action>


                </v-list-item>

                
            </v-list>
          </v-card>
      </div>
      </v-flex>
    </v-layout>

    <v-card-actions>
    <v-btn text color="green" @click="balance">View balance</v-btn>
    </v-card-actions>
  </v-container>
</template>

<script>
import axios from 'axios'
export default {
  props: ['id'],
  data () {
    return {
      BillSplit: '',
      Expenses: [],
      Participants : [],
      Balance: '',
      Debts: [],
      loading: false
    }
  },
  mounted () {
    axios
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id)
      .then(response => {
        this.BillSplit = response.data
      })
      .catch(error => {
        console.log(error)
      })

    axios
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id+'/expenses')
      .then(response => {
        this.Expenses = response.data
      })
      .catch(error => {
        console.log(error)
      })


    axios
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id+'/participants')
      .then(response => {
        this.Participants = response.data
      })
      .catch(error => {
        console.log(error)
      })

  },
  methods: {
  back () {
    this.$router.push('/')
  },
  balance () {
    this.$router.push({
        path: '/billsplit/'+this.id+'/balance',
        })
  },
  newExpense () {
    this.$router.push({
        path: '/billsplit/'+this.id+'/new',
        })
  },
  getExpense(expenseId) {
    this.$router.push({
        path: '/billsplit/'+this.id+'/expense/'+expenseId,
        })
  },
  manageParticipants() {
    this.$router.push({
        path: '/billsplit/'+this.id+'/participants'
        })
  }


  }
}
</script><style>

</style>