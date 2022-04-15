


<template>

  <v-container>

    <v-card>
          <v-card-title primary-title>
            <div>
              <h2>Create a new Bill Split</h2>
            </div>
          </v-card-title>
          <v-card-actions>
            <v-btn text color="green" @click="back">back</v-btn>
          </v-card-actions>

      </v-card>

  <v-form v-on:submit.prevent="addNewParticipant">

    <v-text-field
        v-model="name"
        :counter="50"
        :rules="nameRules"
        label="Bill split name"
        required
    ></v-text-field>


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
                <v-btn v-on:click="Participants.splice(index, 1)">Remove</v-btn>
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
  data () {
    return {
      valid: true,
      payer : '',
      amount: '',
      name : '',
      newPartName: '',
      Participants: [],


      nameRules: [
        v => (v && v.length <= 50) || 'Name must be less than 50 characters',
      ],

      BillSplit: '',
    }
  },


  methods: {
    back () {
      this.$router.push('/')
    },


    addNewParticipant: function () {
      this.Participants.push({
        Name: this.newPartName
      })
      console.log(this.Participants)
          var names = this.Participants.map(function(item) {
      return item['Name'];
    });
        console.log(names)
      this.newPartName = ''
    },

    newBillsplit () {
        console.log(this.Participants)

    var names = this.Participants.map(function(item) {
      return item['Name'];
    });
        console.log(names)

      const data = {
        Name: this.name,
        Participants: names
      };
      console.log(data);

      axios.post(process.env.VUE_APP_BACK_ADDR+'/billsplit/new', data, 
        {
        headers: {
          'Content-Type': 'application/json;charset=utf-8', 
          'Access-Control-Allow-Origin': '*'
        }
       });
      this.$router.push('/');
    },
  }
}
</script>
<style>
</style>
