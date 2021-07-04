<template>
  <div>
    <v-card class="mx-auto my-4" max-width="374">
      <LooseLeafForm :initialLooseLeaf="newLooseLeaf" :isWaitingResponse="isWaitingResponse" :isEditMode="isNewMode" @submit="createLooseLeaf" @setEditMode="setNewMode"/>
    </v-card>
    <LooseLeaf
    v-for="looseLeaf in looseLeafs"
    :key="looseLeaf.ID"
    :looseLeaf="looseLeaf"
    :tags="tags"
    @replaceLooseLeaf="replaceLooseLeaf"
    @removeLooseLeaf="removeLooseLeaf" />
    <infinite-loading @infinite="infiniteHandler"></infinite-loading>
    <div class="d-flex justify-center">
      <v-btn v-show="loadButtonVisible" @click="loadLooseLeafs">
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
      looseLeafs: [],
      newLooseLeaf: {
        ID: null,
        Content: ''
      },
      page: 1,
      pageSize: 10,
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
  methods: {
    async createLooseLeaf (looseLeaf) {
      const self = this
      self.isNewMode = false
      if (looseLeaf.Content === '') { return }
      self.isWaitingResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$post('/loose_leafs', {
          content: looseLeaf.Content
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          looseLeaf.Content = ''
          self.looseLeafs.unshift(response)
          self.isWaitingResponse = false
        })
        .catch(function (error) {
          console.log(error)
          self.isWaitingReponse = false
        })
    },
    async infiniteHandler ($state) {
      const self = this
      let newLooseLeafs = await self.getLooseLeafs()
      if (newLooseLeafs) {
        if (newLooseLeafs.length > 0) {
          newLooseLeafs = self.removeOverlappedLooseLeafs(newLooseLeafs)
          self.page += 1
          self.looseLeafs.push(...newLooseLeafs)
          $state.loaded()
        } else {
          self.loadButtonVisible = false
          $state.complete()
        }
      }
    },
    async getLooseLeafs () {
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      let newLooseLeafs = null
      await self.$axios
        .$get('/loose_leafs', {
          headers: {
            Authorization: `Bearer ${token}`
          },
          params: {
            page: self.page,
            pageSize: self.pageSize
          }
        })
        .then(function (response) {
          newLooseLeafs = response
        })
        .catch(function (error) {
          console.log(error)
        })
      return newLooseLeafs
    },
    async loadLooseLeafs () {
      const self = this
      let newLooseLeafs = await self.getLooseLeafs()
      if (newLooseLeafs) {
        if (newLooseLeafs.length > 0) {
          newLooseLeafs = self.removeOverlappedLooseLeafs(newLooseLeafs)
          self.page += 1
          self.looseLeafs.push(...newLooseLeafs)
        } else {
          self.loadButtonVisible = false
        }
      }
    },
    removeOverlappedLooseLeafs (newLooseLeafs) {
      const self = this
      const looseLeafIds = self.looseLeafs.map(looseLeaf => looseLeaf.ID)
      return newLooseLeafs.filter((newLooseLeaf) => {
        return !looseLeafIds.includes(newLooseLeaf.ID)
      })
    },
    setNewMode (newMode) {
      this.isNewMode = newMode
    },
    replaceLooseLeaf (response) {
      const looseLeaf = this.looseLeafs.find(looseLeaf => looseLeaf.ID === response.ID)
      looseLeaf.Content = response.Content
    },
    removeLooseLeaf (id) {
      const index = this.looseLeafs.findIndex(looseLeaf => looseLeaf.ID === id)
      this.looseLeafs.splice(index, 1)
      this.page = Math.ceil(this.looseLeafs.length / this.pageSize)
    }
  }
}
</script>
