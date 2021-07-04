<template>
  <v-card-text>
    <v-form v-show="isEditMode">
      <v-text-field
        label="Name"
        v-model="binder.Name"
        ref="nameTextField"
        @blur="submit"
      ></v-text-field>
    </v-form>
    <v-progress-circular
      indeterminate
      v-show="isWaitingResponse" />
    <div v-show="!isEditMode && !isWaitingResponse">
      <div @click="clickName">
        {{ binder.Name || 'Add new binder' }}
      </div>
    </div>
  </v-card-text>
</template>

<script>
export default {
  props: {
    initialBinder: Object,
    isWaitingResponse: Boolean,
    isEditMode: Boolean
  },
  data () {
    return {
      binder: {
        ID: this.initialBinder.ID,
        Name: this.initialBinder.Name
      }
    }
  },
  methods: {
    submit () {
      this.$emit('setEditMode', false)
      if (this.initialBinder.Name === this.binder.Name) { return }
      this.$emit('submit', this.binder)
    },
    clickName () {
      this.$emit('setEditMode', true)
      this.$nextTick(() => this.$refs.nameTextField.focus())
    }
  }
}
</script>
