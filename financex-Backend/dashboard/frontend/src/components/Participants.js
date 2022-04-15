


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
              <h2  class="headline mb-0">Manage participants</h2>
            </div>
              </v-card-subtitle>
          <v-card-actions>
            <v-btn text color="green" @click="back">back</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>

  <v-form v-on:submit.prevent="addNewParticipant">

    <v-text-field
      v-model="newPartName"
      id="new-part-name"
      label="Add a Participant"
      :counter="50"
      :rules="nameRules"
    ></v-text-field>
    <v-btn
    @click="addNewParticipant()"
    >Add</v-btn>
  </v-form>
  
             <v-list >
                <v-subheader>Participants</v-subheader>
                <v-list-item
                    v-for="(item, index) in Participants"
                    :key="index"
                >
                <v-list-item-content>
                    <v-list-item-title> {{item.Name}} </v-list-item-title>
                </v-list-item-content>
                </v-list-item>
                <v-list-item
                    v-for="(item, index) in newParticipants"
                    :key="index"
                >
                <v-list-item-content>
                    <v-list-item-title> {{item.Name}} </v-list-item-title>
                </v-list-item-content>
                <v-btn v-on:click="newParticipants.splice(index, 1)">Remove</v-btn>
                </v-list-item>
            </v-list>


                    <v-btn
          :disabled="!valid"
          color="success"
          class="mr-4"
          @click="newBillsplit"
        >
          Validate
        </v-btn>



  </v-container>



</template>

<script>
import axios from 'axios'

export default {
 props: ['id'],
  data () {
    return {
      valid: true,
      newPartName: '',
      newParticipants: [],
      Participants: [],


      nameRules: [
        v => (v && v.length <= 50) || 'Name must be less than 50 characters',
      ],

      BillSplit: '',
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
      this.$router.push('/billsplit/'+this.id)
    },




    addNewParticipant: function () {
      this.newParticipants.push({
        Name: this.newPartName
      })
      console.log(this.newParticipants)
          var names = this.newParticipants.map(function(item) {
      return item['Name'];
    });
        console.log(names)
      this.newPartName = ''
    },

    newBillsplit () {
        console.log(this.newParticipants)

    var names = this.newParticipants.map(function(item) {
      return item['Name'];
    });
        console.log(names)

      axios.post(process.env.VUE_APP_BACK_ADDR+'/billsplit/'+this.id+'/participants/new', names, 
        {
        headers: {
          'Content-Type': 'application/json;charset=utf-8', 
          'Access-Control-Allow-Origin': '*'
        }
       });
    },
  }
}
</script>
<style>
</style>
