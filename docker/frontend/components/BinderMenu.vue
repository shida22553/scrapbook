<template>
  <div class="text-center mb-8">
    <BinderMenuItem
      :name="'All'"
      :color="'grey darken-4'"
      :binderId="''"
    />
    <BinderMenuItem
      v-for="(binder, index) in binders"
      :key="binder.ID"
      :name="binder.Name"
      :color="getBinderColor(index)"
      :binderId="binder.ID.toString()"
    />
  </div>
</template>

<script>
export default {
  components: {
  },
  props: {

  },
  data () {
    return {
      binders: [],
      newBinder: {
        ID: null,
        Name: ''
      },
      page: 1,
      pageSize: 10,
      loadButtonVisible: true,
      isNewMode: false,
      isWaitingResponse: false
    }
  },
  async mounted () {
    const self = this
    const token = await self.$fire.auth.currentUser?.getIdToken(true)
    await self.$axios
      .$get('/binders', {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      .then(function (response) {
        console.log(response)
        self.binders.push(...response)
      })
      .catch(function (error) {
        console.log(error)
      })
  },
  methods: {
    getBinderColor (index) {
      const colors = ['red', 'pink', 'purple', 'deep-purple', 'indigo', 'blue', 'light-blue', 'cyan', 'teal', 'green', 'light-green', 'lime', 'yellow', 'amber', 'orange', 'deep-orange']
      const colorIndex = Math.floor((index / this.binders.length) * colors.length)
      return colors[colorIndex] + ' darken-4'
    }
  }
}
</script>
