<template>
  <div class="ds-week">
    <template v-for="(day, i) in days">
    <Day :key="i" :day="day"/>
  </template>
  </div>
</template>
<script>
import Day from '@/components/Calendar/Day.vue'

export default {
  name: "DayRow",
  components: {
    Day
  },
  props: {
    days: {
      required: true,
      type: Array
    }
  },

  methods: {}
};
</script>

<style scoped lang="scss">
.ds-week {
  display: flex;
  flex: 1;
}
</style>
