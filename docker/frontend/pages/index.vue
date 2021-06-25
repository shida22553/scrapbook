<template>
  <div>
    <h1>Welcome <span v-if="currentUser">{{ currentUser.email }}</span></h1>
    <v-btn class="mr-4" @click="signout">
      signout
    </v-btn>
    <div v-if="currentUser != null">
      <v-card class="mx-auto my-4" max-width="374">
        <CuttingForm :initialCutting="newCutting" :isWaitingResponse="isWaitingResponse" :isEditMode="isNewMode" @submit="createCutting" @setEditMode="setNewMode"/>
      </v-card>
      <Cutting
      v-for="cutting in cuttings"
      :key="cutting.ID"
      :cutting="cutting"
      :tags="tags"
      @replaceCutting="replaceCutting"
      @removeCutting="removeCutting" />
      <infinite-loading @infinite="infiniteHandler"></infinite-loading>
      <div class="d-flex justify-center">
        <v-btn v-show="loadButtonVisible" @click="loadCuttings">
          load
        </v-btn>
      </div>
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
      newCutting: {
        ID: null,
        Note: ''
      },
      page: 1,
      pageSize: 2,
      loadButtonVisible: true,
      isNewMode: false,
      isWaitingResponse: false,
      tags: [
        {
          ID: 1,
          Name: 'aaa'
        }, {
          ID: 2,
          Name: 'bbb'
        }
      ]
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
    async createCutting (cutting) {
      const self = this
      self.isNewMode = false
      if (cutting.Note === '') { return }
      self.isWaitingResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$post('/cuttings', {
          note: cutting.Note
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          cutting.Note = ''
          self.cuttings.unshift(response)
          self.isWaitingResponse = false
        })
        .catch(function (error) {
          console.log(error)
          self.isWaitingReponse = false
        })
    },
    async infiniteHandler ($state) {
      const self = this
      let newCuttings = await self.getCuttings()
      if (newCuttings) {
        if (newCuttings.length > 0) {
          newCuttings = self.removeOverlappedCuttings(newCuttings)
          self.page += 1
          self.cuttings.push(...newCuttings)
          $state.loaded()
        } else {
          self.loadButtonVisible = false
          $state.complete()
        }
      }
    },
    async getCuttings () {
      console.log(this.page)
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      let newCuttings = null
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
          newCuttings = response
        })
        .catch(function (error) {
          console.log(error)
        })
      return newCuttings
    },
    async loadCuttings () {
      const self = this
      let newCuttings = await self.getCuttings()
      if (newCuttings) {
        if (newCuttings.length > 0) {
          newCuttings = self.removeOverlappedCuttings(newCuttings)
          self.page += 1
          self.cuttings.push(...newCuttings)
        } else {
          self.loadButtonVisible = false
        }
      }
    },
    removeOverlappedCuttings (newCuttings) {
      const self = this
      const cuttingIds = self.cuttings.map(cutting => cutting.ID)
      return newCuttings.filter((newCutting) => {
        return !cuttingIds.includes(newCutting.ID)
      })
    },
    setNewMode (newMode) {
      this.isNewMode = newMode
    },
    replaceCutting (response) {
      const cutting = this.cuttings.find(cutting => cutting.ID === response.ID)
      cutting.Note = response.Note
    },
    removeCutting (id) {
      const index = this.cuttings.findIndex(cutting => cutting.ID === id)
      this.cuttings.splice(index, 1)
      this.page = Math.ceil(this.cuttings.length / this.pageSize)
    }
  }
}
</script>
