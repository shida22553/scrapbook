<template>
  <v-card class="mx-auto my-4" max-width="374">
    <div v-show="!isEditMode">
      <v-card-text>
        <div @click="isEditMode = true">
          {{ cutting.Note }}
        </div>
      </v-card-text>

      <v-divider class="mx-4"></v-divider>
      <v-card-actions>
        <v-btn
          text
          @click="$emit('deleteCutting', cutting.ID)"
        >
          Delete
        </v-btn>
      </v-card-actions>
    </div>
    <div v-show="isEditMode">
      <v-card-text>
        <CuttingForm :initialCutting="cutting" @submit="updateCutting" @setEditMode="setEditMode"/>
      </v-card-text>
    </div>
  </v-card>
</template>

<script>
export default {
  props: {
    cutting: Object
  },
  data () {
    return {
      isEditMode: false
    }
  },
  methods: {
    setEditMode (editMode) {
      this.isEditMode = editMode
    },
    updateCutting (cutting) {
      this.$emit('updateCutting', cutting)
    }
  }
}
</script>
