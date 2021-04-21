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
    <hr>
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
    <v-btn class="mr-4" @click="signout">
      signout
    </v-btn>
    <div v-if="currentUser != null">
      <v-textarea
        label="body"
        v-model="cuttingNote"
      ></v-textarea>
      <v-btn class="mr-4" @click="stick">
        stick
      </v-btn>
      <p v-for="cutting in cuttings" v-bind:key="cutting.id">
        {{ cutting.note }}
      </p>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      user: {
        email: '18051976@example.com',
        password: 'password'
      },
      newUser: {
        name: 'name',
        email: `${Math.floor(Math.random() * 100000000)}@example.com`,
        password: 'password'
      },
      cuttings: [],
      cuttingNote: ''
    }
  },
  async mounted () {
    const self = this
    const token = await self.$fire.auth.currentUser?.getIdToken(true)
    if (token != null) {
      await self.$axios
        .$get('/cuttings', {
          headers: {
            Authorization: `Bearer ${token}`
          },
          data: {}
        })
        .then(function (response) {
          self.cuttings = response.data
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user
    }
  },
  methods: {
    async signup () {
      const self = this
      const userCredential = await self.$fire.auth.createUserWithEmailAndPassword(self.newUser.email, self.newUser.password)
      const token = await userCredential.user.getIdToken(true)
      await self.$axios
        .$post('/users', {}, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
    },
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
    },
    async signout () {
      const self = this
      await self.$fire.auth.signOut()
      location.reload()
    },
    async stick () {
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      await self.$axios
        .$post('/cuttings', {
          note: self.cuttingNote
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          console.log(response)
          self.cuttingNote = ''
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
