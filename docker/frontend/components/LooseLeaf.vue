<template>
  <v-card class="mx-auto my-4" max-width="374" :class="getBinderColor(binders.findIndex(binder => binder.ID == looseLeaf.BinderID), 'darken-4')">
    <LooseLeafForm
    :initialLooseLeaf="looseLeaf"
    :isWaitingResponse="isWaitingResponse"
    :isEditMode="isEditMode"
    @submit="updateLooseLeaf"
    @setEditMode="setEditMode"/>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn
        x-small
        @click="isBindersVisible = !isBindersVisible"
      >
        Bind
      </v-btn>
      <!-- <v-btn
        x-small
        @click="deleteLooseLeaf(looseLeaf.ID)"
      >
        Remove
      </v-btn> -->
    </v-card-actions>
    <v-card-actions v-show="isBindersVisible">
      <v-spacer></v-spacer>
      <v-progress-circular
      indeterminate
      v-show="isWaitingPutBinderResponse" />
      <v-chip-group column v-show="!isWaitingPutBinderResponse">
        <v-chip
        v-for="(binder, index) in binders"
        :key="binder.id"
        :color="getBinderColor(index, 'darken-3')"
        :outlined="binder.ID != looseLeaf.BinderID"
        @click="bind(looseLeaf, binder.ID)"
        x-small>
          {{ binder.Name }}
        </v-chip>
      </v-chip-group>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: {
    looseLeaf: Object,
    binders: Array
  },
  data () {
    return {
      isEditMode: false,
      isWaitingResponse: false,
      isBindersVisible: false,
      isWaitingPutBinderResponse: false
    }
  },
  methods: {
    setEditMode (editMode) {
      this.isEditMode = editMode
    },
    async updateLooseLeaf (looseLeaf) {
      const self = this
      self.editMode = false
      if (looseLeaf.Content === '') { return }
      this.isWaitingResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$put(`/loose_leafs/${looseLeaf.ID}`, {
          content: looseLeaf.Content
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          // console.log(response)
          self.$emit('replaceLooseLeaf', response)
          self.isWaitingResponse = false
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    async deleteLooseLeaf (id) {
      if (!confirm('Delete this?')) { return }
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$delete(`/loose_leafs/${id}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          // console.log(response)
          self.$emit('removeLooseLeaf', id)
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    getBinderColor (index, option) {
      const colors = ['red', 'pink', 'purple', 'deep-purple', 'indigo', 'blue', 'light-blue', 'cyan', 'teal', 'green', 'light-green', 'lime', 'yellow', 'amber', 'orange', 'deep-orange']
      const colorIndex = Math.floor((index / this.binders.length) * colors.length)
      return `${colors[colorIndex]} ${option}`
    },
    async bind (looseLeaf, binderId) {
      const self = this
      self.isWaitingPutBinderResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$put(`/loose_leafs/${looseLeaf.ID}/binder`, {
          BinderId: binderId
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          self.$emit('replaceLooseLeaf', response)
          self.isWaitingPutBinderResponse = false
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
