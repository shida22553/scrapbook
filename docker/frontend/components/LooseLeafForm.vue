<template>
  <v-card-text>
    <v-form v-show="isEditMode">
      <v-textarea
        label="content"
        v-model="looseLeaf.Content"
        ref="contentTextArea"
        @blur="submit"
      ></v-textarea>
    </v-form>
    <v-progress-circular
      indeterminate
      v-show="isWaitingResponse" />
    <div v-show="!isEditMode && !isWaitingResponse">
      <div @click="clickContent">
        {{ looseLeaf.Content || 'Add new loose leaf' }}
      </div>
      <v-chip-group
        column
      >
      </v-chip-group>
    </div>
  </v-card-text>
</template>

<script>
export default {
  props: {
    initialLooseLeaf: Object,
    isWaitingResponse: Boolean,
    isEditMode: Boolean
  },
  data () {
    return {
      looseLeaf: {
        ID: this.initialLooseLeaf.ID,
        Content: this.initialLooseLeaf.Content
      }
    }
  },
  methods: {
    submit () {
      this.$emit('setEditMode', false)
      if (this.initialLooseLeaf.Content === this.looseLeaf.Content) { return }
      this.$emit('submit', this.looseLeaf)
    },
    clickContent () {
      this.$emit('setEditMode', true)
      this.$nextTick(() => this.$refs.contentTextArea.focus())
    }
  }
}
</script>
