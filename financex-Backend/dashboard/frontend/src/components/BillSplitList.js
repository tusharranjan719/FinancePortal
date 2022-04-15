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
  </v-container><v-container v-else grid-list-xl>

        <v-card>
          <v-card-title primary-title>
            <div>
              <h2>Welcome to your bill splits</h2>
            </div>
          </v-card-title>
          <v-card-actions>
            <v-btn @click="newBillSplit"> new Billsplit</v-btn>
          </v-card-actions>
      </v-card>

    <v-layout wrap>



      <v-flex xs4
        v-for="(item, index) in wholeResponse"
        :key="index"
        mb-2>
        <v-card>
          <v-card-title primary-title>
            <div>
              <h2>{{item.Name}}</h2>
            </div>
          </v-card-title><v-card-actions class="justify-center">
            <v-btn text
              color="green"
              @click="singleBillSplit(item.Uuid)"
              >View</v-btn>
          </v-card-actions></v-card>
      </v-flex>
  </v-layout>
  </v-container>
</template>

<script>
import axios from 'axios'
export default {
    data () {
    return {
      wholeResponse: [],
      loading: false
    }
  },
  mounted () {
  axios
    .get(process.env.VUE_APP_BACK_ADDR+'/')
    .then(response => {
      this.wholeResponse = response.data
      this.loading = false
    })
    .catch(error => {
      console.log(error)
    })
  },
    methods: {
    singleBillSplit (id) {
      console.log(id)
      this.$router.push('/billsplit/' + id)
    },
    newBillSplit () {
      this.$router.push('/new')
    }
  }
}
</script>

