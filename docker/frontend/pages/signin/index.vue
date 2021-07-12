<template>
  <div>
    <h1>Welcome <span v-if="currentUser">{{ currentUser.email }}</span></h1>
    <v-text-field
      v-model="user.email"
      label="E-mail"
      :rules="[required]"
    />
    <v-text-field
      v-model="user.password"
      label="Password"
      :rules="[required]"
    />
    <v-btn class="mr-4" @click="signin" :disabled="disabled">
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
        email: '4124325221@example.com',
        password: 'password'
      },
      required: value => !!value || 'required'
    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user
    },
    disabled () {
      return !this.user.email || !this.user.password
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
