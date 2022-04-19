


<template>

  <v-container>


    

    <v-layout wrap>
      <v-flex xs12 mr-1 ml-1>
        <v-card>
          <v-card-title primary-title>
            <div>
              <h2 class="headline mb-0">{{this.Billsplit.Name}}  -  {{this.Expense.Name}}</h2>
            </div>
          <v-card-subtitle>
              {{this.Expense.CreatedAt}}
          </v-card-subtitle>
          </v-card-title >

          <v-card-actions>
            <v-btn text color="green" @click="back">back</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>


    <v-layout row wrap>
      <v-flex xs12>
        <div class="text-xs-center">
          <v-card>

            <div>
                              <v-subheader>Payer</v-subheader>
                              {{Expense.PayerName}}
            </div>
            <div>
                              <v-subheader>Amount</v-subheader>
                              {{parseFloat(Expense.Amount).toFixed( 2 )}}$
            </div>
                
             <v-list >
                <v-subheader>Participants</v-subheader>
                <v-list-item
                    v-for="(item, index) in this.Expense.Participants"
                    :key="index"
                >
                <v-list-item-content>
                    <v-list-item-title> {{item}} </v-list-item-title>
                </v-list-item-content>
                </v-list-item>
            </v-list>

          </v-card>
      </div>
      </v-flex>
    </v-layout>

  </v-container>

</template>

<script>
import axios from 'axios'

export default {
  props: ['id','expense_id'],
  data () {
    return {
      valid: true,
      Expense: '',
      Billsplit : ''
    }
  },


  mounted () {

axios.interceptors.request.use(request => {
  console.log('Starting Request', request)
  return request
})

    axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
    axios
      .get(process.env.VUE_APP_BACK_ADDR+'/expense/'+this.expense_id)
      .then(response => {
        this.Expense = response.data
        console.log(this.Expense)
      })
      .catch(error => {
        console.log(error)
      })  

    axios
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id)
      .then(response => {
        this.Billsplit = response.data
        console.log(this.Billsplit)
      })
      .catch(error => {
        console.log(error)
      })

  },
  methods: {
    back () {
      this.$router.push('/billsplit/'+this.id)
    },
  }
}
</script>
<style>
</style>
