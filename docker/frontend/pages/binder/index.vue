<template>
  <div>
    <div v-if="currentUser != null">
      <BinderList />
    </div>
  </div>
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
