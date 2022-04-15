


<template>

  <v-container>

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
                      <v-card-subtitle>
                          <div>
              <h2  class="headline mb-0">Add new expense</h2>
            </div>
              </v-card-subtitle>
          <v-card-actions>
            <v-btn text color="green" @click="back">back</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>

    <v-form
        ref="form"
        v-model="valid"
      >
        <v-text-field
          v-model="name"
          :counter="50"
          :rules="nameRules"
          label="Expense name"
          required
        ></v-text-field>
  
        <v-text-field
          v-model="amount"
          :rules="amountRules"
          label="Amount"
          required
        ></v-text-field>
  
        <v-select
          v-model="payer"
          item-text="Name"
          :items="this.Participants"
          :rules="[v => v.length>0 || 'Item is required']"
          label="Payer"
        ></v-select>
  
      <v-list flat>
                  <v-list-item v-for="(item) in Participants" :key="item.Name">
                    <v-list-item-action>
                      <v-checkbox
                       v-model="selectedParticipants" 
                       multiple 
                       :value="item.Name"
                       required />
                    </v-list-item-action>
                    <v-list-item-content>
                      <v-list-item-title>{{ item.Name }}</v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
      </v-list>

      <p>Selected participants:</p>
      <v-row flat>
              <v-col v-for="(item, index) in selectedParticipants" :key="index">
                  <v-card>{{ item }}</v-card>
              </v-col>
      </v-row>


        <v-btn
          :disabled="!valid"
          color="success"
          class="mr-4"
          @click="validate"
        >
          Validate
        </v-btn>
    </v-form>


    <v-layout row wrap>
        
    </v-layout>

  </v-container>

</template>

<script>
import axios from 'axios'

export default {
  props: ['id'],
  data () {
    return {
      valid: true,
      payer : '',
      amount: '',
      name : '',
      amountRules: [
        v => !isNaN(v)  || 'Amount is required'
      ], 

      nameRules: [
        v => !!v || 'Name is required',
        v => (v && v.length <= 50) || 'Name must be less than 50 characters',
      ],

      BillSplit: '',
      Participants: [],
      selectedParticipants: []
    }
  },


  mounted () {

axios.interceptors.request.use(request => {
  console.log('Starting Request', request)
  return request
})

    axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
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
      .get(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id+'/participants')
      .then(response => {
        this.Participants = response.data
        console.log(this.Participants)
      })
      .catch(error => {
        console.log(error)
      })


  },
  methods: {
    back () {
      this.$router.push('/billsplit/'+this.id)
    },
    validate () {

      const data = {
        expense: this.name,
        amount: Number(this.amount), 
        payer: this.payer,
        participants: this.selectedParticipants
      };
      console.log(data);

      axios.post(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id+'/expenses/new', data, 
        {
        headers: {
          'Content-Type': 'application/json;charset=utf-8', 
          'Access-Control-Allow-Origin': '*'
        }
       });

      this.$router.push('/billsplit/'+this.id);
    },
  }
}
</script>
<style>
</style>
