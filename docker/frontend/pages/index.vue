<template>
  <v-row>
    <v-col>
      <div v-if="currentUser != null">
        <BinderMenu/>
        <LooseLeafList :initialBinderId="null" />
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
