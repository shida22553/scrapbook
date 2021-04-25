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
      const userCredential = await self.$fire.auth.signInWithEmailAndPassword(self.user.email, self.user.password)
      const token = await userCredential.user.getIdToken(true)
      await self.$axios
        .$get('/users', {
          headers: {
            Authorization: `Bearer ${token}`
          },
          data: {}
        })
        .then(function (response) {
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
