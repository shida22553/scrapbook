<template>
  <v-card-text>
    <v-form v-show="isEditMode">
      <v-textarea
        label="note"
        v-model="cutting.Note"
        ref="noteTextArea"
        @blur="submit"
      ></v-textarea>
    </v-form>
    <v-progress-circular
      indeterminate
      v-show="isWaitingResponse" />
    <div v-show="!isEditMode && !isWaitingResponse">
      <div @click="clickNote">
        {{ cutting.Note || 'Stick new cutting' }}
      </div>
      <v-chip-group
        column
      >
        <v-chip
          small
          v-for="tag in tags"
          :key="tag.ID"
        >
          {{ tag.Name }}
        </v-chip>
      </v-chip-group>
    </div>
  </v-card-text>
</template>

<script>
export default {
  props: {
    initialCutting: Object,
    initialTags: Array,
    isWaitingResponse: Boolean,
    isEditMode: Boolean
  },
  data () {
    return {
      cutting: {
        ID: this.initialCutting.ID,
        Note: this.initialCutting.Note
      },
      tags: this.initialTags
    }
  },
  methods: {
    submit () {
      this.$emit('setEditMode', false)
      if (this.initialCutting.Note === this.cutting.Note) { return }
      this.$emit('submit', this.cutting)
    },
    clickNote () {
      this.$emit('setEditMode', true)
      this.$nextTick(() => this.$refs.noteTextArea.focus())
    }
  }
}
</script>
