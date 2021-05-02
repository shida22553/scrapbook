<template>
  <v-card class="mx-auto my-4" max-width="374">

    <v-card-text>
      <div>
        {{ note }}
      </div>
    </v-card-text>

    <v-divider class="mx-4"></v-divider>
  </v-card>
</template>

<script>
export default {
  props: {
    note: {
      type: String,
      required: true
    }
  }
}
</script>
