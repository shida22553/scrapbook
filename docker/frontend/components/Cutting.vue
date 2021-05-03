<template>
  <v-card class="mx-auto my-4" max-width="374">

    <v-card-text>
      <div>
        {{ note }}
      </div>
    </v-card-text>

    <v-divider class="mx-4"></v-divider>
    <v-card-actions>
      <nuxt-link class="text-decoration-none" :to="`/cuttings/${id}`">
        <v-btn
          text
        >
          Edit
        </v-btn>
      </nuxt-link>
      <v-btn
        text
        @click="deleteCutting(id)"
      >
        Delete
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: {
    id: {
      type: Number,
      required: true
    },
    note: {
      type: String,
      required: true
    }
  },
  methods: {
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
          self.$emit('delete', self.id)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
