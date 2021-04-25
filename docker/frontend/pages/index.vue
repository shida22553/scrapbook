<template>
  <div>
    <h1>Welcome <span v-if="currentUser">{{ currentUser.email }}</span></h1>
    <v-btn class="mr-4" @click="signout">
      signout
    </v-btn>
    <div v-if="currentUser != null">
      <v-textarea
        label="body"
        v-model="cuttingNote"
      ></v-textarea>
      <v-btn class="mr-4" @click="stick">
        stick
      </v-btn>
      <p v-for="cutting in cuttings" v-bind:key="cutting.id">
        {{ cutting.note }}
      </p>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      cuttings: [],
      cuttingNote: ''
    }
  },
  async mounted () {
    const self = this
    const token = await self.$fire.auth.currentUser?.getIdToken(true)
    if (token != null) {
      await self.$axios
        .$get('/cuttings', {
          headers: {
            Authorization: `Bearer ${token}`
          },
          data: {}
        })
        .then(function (response) {
          self.cuttings = response.data
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user
    }
  },
  methods: {
    async signout () {
      const self = this
      await self.$fire.auth.signOut()
      location.reload()
    },
    async stick () {
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$post('/cuttings', {
          note: self.cuttingNote
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          console.log(response)
          self.cuttingNote = ''
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
