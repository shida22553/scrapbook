<template>
  <div>
    <h1>Welcome <span v-if="currentUser">{{ currentUser.email }}</span></h1>
    <v-text-field
      v-model="user.email"
      label="E-mail"
    />
    <v-text-field
      v-model="user.password"
      label="Password"
    />
    <v-btn class="mr-4" @click="signin">
      signin
    </v-btn>
    <nuxt-link to="/signup">Sign up</nuxt-link>
  </div>
</template>

<script>
export default {
  data () {
    return {
      user: {
        email: '18051976@example.com',
        password: 'password'
      }
    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user
    }
  },
  methods: {
    async signin () {
      const self = this
      await self.$fire.auth.signInWithEmailAndPassword(self.user.email, self.user.password)
      location.href = '/'
    }
  }
}
</script>
