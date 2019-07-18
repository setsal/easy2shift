<template>
  <div>
    hello
  </div>
</template>

<script>
export default {
  name: 'MonthSelector',
  props: {
    isActive: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    toggleClick() {
      this.$emit('toggleClick')
    }
  }
}
</script>

<style scoped>

</style>
