<template>
  <div>
    <v-textarea
      label="body"
      v-model="cuttingNote"
    ></v-textarea>
    <v-btn class="mr-4" @click="updateCutting">
      update
    </v-btn>
    </div>
</template>

<script>
export default {
  data () {
    return {
      cuttingNote: ''
    }
  },
  async mounted () {
    const self = this
    const token = await self.$fire.auth.currentUser?.getIdToken(true)
    if (token != null) {
      await self.$axios
        .$get(`/cuttings/${self.$route.params.id}`, {
          headers: {
            Authorization: `Bearer ${token}`
          },
          data: {}
        })
        .then(function (response) {
          self.cuttingNote = response.Note
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  methods: {
    async updateCutting () {
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$put(`/cuttings/${self.$route.params.id}`, {
          note: self.cuttingNote
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          console.log(response)
          self.$router.push('/')
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
