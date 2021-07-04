<template>
  <v-card class="mx-auto my-4" max-width="374">
    <LooseLeafForm
    :initialLooseLeaf="looseLeaf"
    :initialTags="tags"
    :isWaitingResponse="isWaitingResponse"
    :isEditMode="isEditMode"
    @submit="updateLooseLeaf"
    @setEditMode="setEditMode"/>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn
        x-small
        @click="deleteLooseLeaf(looseLeaf.ID)"
      >
        Remove
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: {
    looseLeaf: Object,
    tags: Array
  },
  data () {
    return {
      isEditMode: false,
      isWaitingResponse: false
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
          console.log(response)
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
          console.log(response)
          self.$emit('removeLooseLeaf', id)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
