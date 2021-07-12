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
    <infinite-loading @infinite="infiniteHandler"></infinite-loading>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'
export default {
  components: {
    InfiniteLoading
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
      pageSize: 30,
      loadButtonVisible: true,
      isNewMode: false,
      isWaitingResponse: false
    }
  },
  mounted () {
  },
  methods: {
    async infiniteHandler ($state) {
      const self = this
      let newBinders = await self.getBinders()
      if (newBinders) {
        if (newBinders.length > 0) {
          newBinders = self.removeOverlappedBinders(newBinders)
          self.page += 1
          self.binders.push(...newBinders)
          $state.loaded()
        } else {
          self.loadButtonVisible = false
          $state.complete()
        }
      }
    },
    async getBinders () {
      const self = this
      const token = await self.$fire.auth.currentUser?.getIdToken(true)
      let newBinders = null
      await self.$axios
        .$get('/binders', {
          headers: {
            Authorization: `Bearer ${token}`
          },
          params: {
            page: self.page,
            pageSize: self.pageSize
          }
        })
        .then(function (response) {
          newBinders = response
        })
        .catch(function (error) {
          console.log(error)
        })
      return newBinders
    },
    removeOverlappedBinders (newBinders) {
      const self = this
      const binderIds = self.binders.map(binder => binder.ID)
      return newBinders.filter((newBinder) => {
        return !binderIds.includes(newBinder.ID)
      })
    },
    getBinderColor (index) {
      const colors = ['red', 'pink', 'purple', 'deep-purple', 'indigo', 'blue', 'light-blue', 'cyan', 'teal', 'green', 'light-green', 'lime', 'yellow', 'amber', 'orange', 'deep-orange']
      const colorIndex = index % this.binders.length
      return colors[colorIndex] + ' darken-4'
    }
  }
}
</script>
