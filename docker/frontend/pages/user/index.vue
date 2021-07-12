<template>
  <div>
    <v-card
      class="mx-auto my-4"
      max-width="400"
    >
      <v-card-text>
        <v-form>
          <v-progress-circular
          indeterminate
          v-show="isWaitingResponse" />
          <v-text-field
            v-show="!isWaitingResponse"
            v-model="name"
            label="Name"
            required
            @blur="updateName"
          ></v-text-field>
        </v-form>
      </v-card-text>
    </v-card>
    <v-card
      class="mx-auto my-4"
      max-width="400"
    >
      <v-list flat>
        <v-list-item-group>
          <v-list-item nuxt to="/binder">
            <v-list-item-content>
              <v-list-item-title>Manage binders</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-divider></v-divider>

          <v-list-item @click="signout">
            <v-list-item-content>
              <v-list-item-title>Sign out</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card>
  </div>
</template>

<script>
export default {
  components: {
  },
  data () {
    return {
      id: null,
      name: '',
      isWaitingResponse: false
    }
  },
  async mounted () {
    const self = this
    const token = await self.$fire.auth.currentUser?.getIdToken(true)
    await self.$axios
      .$get('/user', {
        headers: {
          Authorization: `Bearer ${token}`
        },
        data: {}
      })
      .then(function (response) {
        self.name = response.Name
        self.id = response.ID
        // console.log(response)
      })
      .catch(function (error) {
        console.log(error)
      })
  },
  methods: {
    async updateName () {
      const self = this
      self.isWaitingResponse = true
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      // console.log(self.name)
      await self.$axios
        .$put(`/users/${self.id}`, {
          name: self.name
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .then(function (response) {
          self.name = response.Name
          self.isWaitingResponse = false
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    async signout () {
      const self = this
      await self.$fire.auth.signOut()
    }
  }
}
</script>
