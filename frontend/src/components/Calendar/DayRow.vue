<template>
  <div class="ds-week">
    <!-- <template v-for="(day, i) in days"> -->

    <Day />

    <!-- </template> -->
  </div>
</template>
<script>
import Day from '@/components/Calendar/Day.vue'

export default {
  name: "DayRow",
  components: {
    Day
  },
  props: {},

  methods: {}
};
</script>

<style scoped lang="scss">
.ds-week {
  display: flex;
  flex: 1;
}
</style>
