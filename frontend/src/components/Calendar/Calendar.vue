<template>
  <div class="ds-calendar">
    <div class="ds-month-view">
      <WeekViews :month="month" class="ds-expand" />
    </div>
  </div>
</template>

<script>
import WeekViews from '@/components/Calendar/WeekViews.vue'

export default {
  name: 'Calendar',
  components: {
    WeekViews
  },
  data() {
    return {
      month: ''
    }
  },
  created() {
    this.month = this.$moment().add(1, 'M')
    console.log('Initial Calendar')
  }
}
</script>

<style scoped lang="scss">
.ds-calendar {
  width: 100%;
  height: 100%;
  position: relative;
  padding: 0px !important;
}

.ds-month-view {
  width: 100%;
  height: 100%;

  .ds-today {
    background-color: rgba(0, 0, 0, 0.04);
  }
}
</style>
