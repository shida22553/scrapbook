<template>
  <v-app dark>
    <v-main>
      <v-container fluid>
        <div v-show="currentUser != null">
          <v-row class="ma-2">
            <v-btn class="ma-2" to="/" icon nuxt>
              <v-icon>
                mdi-arrow-left
              </v-icon>
            </v-btn>
            <v-spacer></v-spacer>
            <span class="ma-2" v-show="$route.path !== '/user'">
              <v-btn to="/user" icon nuxt>
                <v-icon>
                  mdi-account
                </v-icon>
              </v-btn>
            </span>
          </v-row>
        </div>
        <nuxt :nuxt-child-key="$route.fullPath" />
      </v-container>
    </v-main>
  </v-app>
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
