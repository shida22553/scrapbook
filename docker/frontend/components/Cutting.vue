<template>
  <v-card class="mx-auto my-4" max-width="374">
    <CuttingForm
    :initialCutting="cutting"
    :initialTags="tags"
    :isWaitingResponse="isWaitingResponse"
    :isEditMode="isEditMode"
    @submit="updateCutting"
    @setEditMode="setEditMode"/>
    <v-card-actions>
      <v-btn
        text
        @click="deleteCutting(cutting.ID)"
      >
        Delete
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: {
    cutting: Object,
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
    async updateCutting (cutting) {
      const self = this
      self.editMode = false
      if (cutting.Note === '') { return }
      this.isWaitingResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$put(`/cuttings/${cutting.ID}`, {
          note: cutting.Note
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          console.log(response)
          self.$emit('replaceCutting', response)
          self.isWaitingResponse = false
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    async deleteCutting (id) {
      if (!confirm('Delete this?')) { return }
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$delete(`/cuttings/${id}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          console.log(response)
          self.$emit('removeCutting', id)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
