<template>
  <div>
    <v-card class="mx-auto my-4" max-width="374">
      <BinderForm :initialBinder="newBinder" :isWaitingResponse="isWaitingResponse" :isEditMode="isNewMode" @submit="createBinder" @setEditMode="setNewMode"/>
    </v-card>
    <Binder
    v-for="(binder, index) in binders"
    :key="binder.ID"
    :binder="binder"
    :color="getBinderColor(index)"
    @replaceBinder="replaceBinder"
    @removeBinder="removeBinder" />
    <infinite-loading @infinite="infiniteHandler"></infinite-loading>
    <div class="d-flex justify-center">
      <v-btn v-show="loadButtonVisible" @click="loadBinders">
        load
      </v-btn>
    </div>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'

export default {
  components: {
    InfiniteLoading
  },
  props: {

  },
  data () {
    return {
      binders: [],
      newBinder: {
        ID: null,
        Name: ''
      },
      page: 1,
      pageSize: 30,
      loadButtonVisible: true,
      isNewMode: false,
      isWaitingResponse: false
    }
  },
  methods: {
    async createBinder (binder) {
      const self = this
      self.isNewMode = false
      if (binder.Name === '') { return }
      self.isWaitingResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$post('/binders', {
          name: binder.Name
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          binder.Name = ''
          self.binders.unshift(response)
          self.isWaitingResponse = false
        })
        .catch(function (error) {
          console.log(error)
          self.isWaitingReponse = false
        })
    },
    async infiniteHandler ($state) {
      const self = this
      let newBinders = await self.getBinders()
      if (newBinders) {
        if (newBinders.length > 0) {
          newBinders = self.removeOverlappedBinders(newBinders)
          self.page += 1
          self.binders.push(...newBinders)
          $state.loaded()
        } else {
          self.loadButtonVisible = false
          $state.complete()
        }
      }
    },
    async getBinders () {
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      let newBinders = null
      await self.$axios
        .$get('/binders', {
          headers: {
            Authorization: `Bearer ${token}`
          },
          params: {
            page: self.page,
            pageSize: self.pageSize
          }
        })
        .then(function (response) {
          newBinders = response
        })
        .catch(function (error) {
          console.log(error)
        })
      return newBinders
    },
    async loadBinders () {
      const self = this
      let newBinders = await self.getBinders()
      if (newBinders) {
        if (newBinders.length > 0) {
          newBinders = self.removeOverlappedBinders(newBinders)
          self.page += 1
          self.binders.push(...newBinders)
        } else {
          self.loadButtonVisible = false
        }
      }
    },
    removeOverlappedBinders (newBinders) {
      const self = this
      const binderIds = self.binders.map(binder => binder.ID)
      return newBinders.filter((newBinder) => {
        return !binderIds.includes(newBinder.ID)
      })
    },
    setNewMode (newMode) {
      this.isNewMode = newMode
    },
    replaceBinder (response) {
      const binder = this.binders.find(binder => binder.ID === response.ID)
      binder.Name = response.Name
    },
    removeBinder (id) {
      const index = this.binders.findIndex(binder => binder.ID === id)
      this.binders.splice(index, 1)
      this.page = Math.ceil(this.binders.length / this.pageSize)
    },
    getBinderColor (index) {
      const colors = ['red', 'pink', 'purple', 'deep-purple', 'indigo', 'blue', 'light-blue', 'cyan', 'teal', 'green', 'light-green', 'lime', 'yellow', 'amber', 'orange', 'deep-orange']
      const colorIndex = index % this.binders.length
      return colors[colorIndex] + ' darken-4'
    }
  }
}
</script>
