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
    <v-layout wrap>
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
        </v-card>
      </v-flex>
    </v-layout>

    <v-layout row wrap>
      <v-flex xs12>
        <div class="text-xs-center">
          <v-card>
             <v-list >
                <v-subheader>Balance</v-subheader>
                <v-list-item
                v-for="(value, key) in Balance"
                :key="key"
                >
                <v-list-item-content v-if="parseFloat(value) >= 0">
                    <v-list-item-title>  <b>{{key}}</b>   Balance: <b style="color:green">{{parseFloat(value).toFixed( 2 )}}$</b>    </v-list-item-title>
                </v-list-item-content>              
                <v-list-item-content v-else>
                    <v-list-item-title>  <b>{{key}}</b>   Balance:  <b style="color:red">{{parseFloat(value).toFixed( 2 )}}$</b>  </v-list-item-title>
                </v-list-item-content>
                </v-list-item>
            </v-list>
          </v-card>
      </div>
      </v-flex>
    </v-layout>

    <v-layout row wrap>
      <v-flex xs12>
        <div class="text-xs-center">
          <v-card>
             <v-list>
                <v-subheader>Debts</v-subheader>
                <v-list-item
                v-for="(item, index) in Debts"
                :key="index"
                >
                <v-list-item-content>
                    <v-list-item-title>  <b>{{item.Debtor}}</b>   owes <b>{{item.Creditor}}</b>  <b>{{parseFloat(item.Amount).toFixed( 2 )}}$</b>      </v-list-item-title>
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
  props: ['id'],
  data () {
    return {
      Balance: [],
      BillSplit: '',
      Debts: [],
      loading: false
    }
  },
  mounted () {
      axios
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id+'/balance')
      .then(response => {
        this.Balance = response.data
        console.log(this.Balance)
      })
      .catch(error => {
        console.log(error)
      })

      axios
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id)
      .then(response => {
        this.BillSplit = response.data
        console.log(this.BillSplit)
      })
      .catch(error => {
        console.log(error)
      })

      axios
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id+'/debts')
      .then(response => {
        this.Debts = response.data
        console.log(this.Debts)
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
</script><style>

.v-list-item__content{
justify-content: center;
text-align: left;
display: grid;
}
</style>