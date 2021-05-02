<template>
  <v-card class="mx-auto my-4" max-width="374">

    <v-card-text>
      <div>
        {{ note }}
      </div>
    </v-card-text>

    <v-divider class="mx-4"></v-divider>
    <v-card-actions>
      <nuxt-link :to="`/cuttings/${id}`">
        <v-btn
          text
        >
          Edit
        </v-btn>
      </nuxt-link>
      <v-btn
        text
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
  }
}
</script>
