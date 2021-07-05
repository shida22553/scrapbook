<template>
  <v-app dark>
    <v-main>
      <v-container fluid>
        <v-row class="ma-2" v-show="$route.path !== '/user'">
          <v-spacer></v-spacer>
          <v-btn to="/user" icon nuxt>
            <v-icon>
              mdi-account
            </v-icon>
          </v-btn>
        </v-row>
        <v-row class="ma-2" v-show="$route.path === '/user'">
          <v-btn to="/" icon nuxt>
            <v-icon>
              mdi-arrow-left
            </v-icon>
          </v-btn>
        </v-row>
        <nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
</script>
