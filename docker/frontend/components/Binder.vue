<template>
  <v-card class="mx-auto my-4" max-width="374" :class="color">
    <BinderForm
    :initialBinder="binder"
    :isWaitingResponse="isWaitingResponse"
    :isEditMode="isEditMode"
    @submit="updateBinder"
    @setEditMode="setEditMode"/>
  </v-card>
</template>

<script>
export default {
  props: {
    binder: Object,
    color: String
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
    async updateBinder (binder) {
      const self = this
      self.editMode = false
      if (binder.name === '') { return }
      this.isWaitingResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$put(`/binders/${binder.ID}`, {
          name: binder.Name
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          // console.log(response)
          self.$emit('replaceBinder', response)
          self.isWaitingResponse = false
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
