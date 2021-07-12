<template>
  <v-chip
    class="ma-2"
    label
    :color="color"
    :to="`loose_leaf?binderId=${binderId}`"
    exact
    nuxt
  >
    <v-icon left>
      mdi-book
    </v-icon>
    {{ name }}
  </v-chip>
</template>

<script>
export default {
  props: {
    color: String,
    name: String,
    binderId: String
  },
  data () {
    return {
    }
  },
  methods: {
  }
}
</script>
