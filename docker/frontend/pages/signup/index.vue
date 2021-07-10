<template>
  <div>
    <v-text-field
      v-model="newUser.name"
      :counter="10"
      label="Name"
    />
    <v-text-field
      v-model="newUser.email"
      label="E-mail"
    />
    <v-text-field
      v-model="newUser.password"
      label="Password"
    />
    <v-btn class="mr-4" @click="signup">
      signup
    </v-btn>
  </div>
</template>

<script>
export default {
  data () {
    return {
      newUser: {
        name: 'name',
        email: `${Math.floor(Math.random() * 100000000)}@example.com`,
        password: 'password'
      }
    }
  },
  methods: {
    async signup () {
      const self = this
      const userCredential = await self.$fire.auth.createUserWithEmailAndPassword(self.newUser.email, self.newUser.password)
      const token = await userCredential.user.getIdToken(true)
      console.log(token)
      await self.$axios
        .$post('/users', {
          name: self.newUser.name
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
          self.$fire.auth.signOut()
        })
    }
  }
}
</script>
