<template>
  <v-row>
    <v-col>
      <div v-if="currentUser != null">
        <LooseLeafList />
      </div>
    </v-col>
  </v-row>
</template>

<script>
export default {
  components: {
  },
  data () {
    return {

    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user
    }
  },
  methods: {

  }
}
</script>
