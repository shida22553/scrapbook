<template>
  <v-form>
    <v-textarea
      label="note"
      v-model="cutting.Note"
      @blur="submit"
    ></v-textarea>
  </v-form>
</template>

<script>
export default {
  props: {
    initialCutting: Object
  },
  data () {
    return {
      cutting: {
        ID: this.initialCutting.ID,
        Note: this.initialCutting.Note
      }
    }
  },
  methods: {
    submit () {
      this.$emit('setEditMode', false)
      if (this.initialCutting.Note === this.cutting.Note) { return }
      this.$emit('submit', this.cutting)
    }
  }
}
</script>
