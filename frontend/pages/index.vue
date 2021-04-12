<template>
  <div>
    <h1>Welcome <span v-if="currentUser">{{ currentUser.email }}</span></h1>
    <v-text-field
      v-model='user.email'
      label='E-mail'
    ></v-text-field>
    <v-text-field
      v-model='user.password'
      label='Password'
    ></v-text-field>
    <v-btn class='mr-4' @click='signin'>signin</v-btn>
    <hr>
    <v-text-field
      v-model='newUser.name'
      :counter='10'
      label='Name'
    ></v-text-field>
    <v-text-field
      v-model='newUser.email'
      label='E-mail'
    ></v-text-field>
    <v-text-field
      v-model='newUser.password'
      label='Password'
    ></v-text-field>
    <v-btn class='mr-4' @click='signup'>signup</v-btn>
  </div>
</template>

<script>
export default {
  data () {
    return {
      user: {
        email: '41243251@example.com',
        password: 'password'
      },
      newUser: {
        name: 'name',
        email: `${Math.floor(Math.random() * 100000000)}@example.com`,
        password: 'password'
      }
    }
  },
  mounted () {
  },
  computed: {
    currentUser () {
      return this.$store.state.user
    }
  },
  methods: {
    async signup () {
      const self = this
      await self.$fire.auth.createUserWithEmailAndPassword(self.newUser.email, self.newUser.password)
    },
    async signin () {
      const self = this
      await self.$fire.auth.signInWithEmailAndPassword(self.user.email, self.user.password)
      // const token = await userCredential.user.getIdToken(true)
    }
  }
}
</script>
