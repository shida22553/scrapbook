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
      <v-btn class="mr-4" @click="create">
        stick
      </v-btn>
      <Cutting v-for="cutting in cuttings" :key="cutting.ID" :note="cutting.Note" :id="cutting.ID"></Cutting>
      <infinite-loading @infinite="infiniteHandler"></infinite-loading>
    </div>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'

export default {
  components: {
    InfiniteLoading
  },
  data () {
    return {
      cuttings: [],
      cuttingNote: '',
      page: 1,
      pageSize: 5
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
    async create () {
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
    },
    async infiniteHandler ($state) {
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$get('/cuttings', {
          headers: {
            Authorization: `Bearer ${token}`
          },
          params: {
            page: self.page,
            pageSize: self.pageSize
          }
        })
        .then(function (response) {
          console.log(response)
          if (response.length) {
            self.page += 1
            self.cuttings.push(...response)
            $state.loaded()
          } else {
            $state.complete()
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
