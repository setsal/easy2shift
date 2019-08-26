<template>
  <div class="ds-month">
    <div class="ds-week-header">
      <div class="ds-week-header-day"
            v-for="(weekday) in weekdays"
            :key="weekday"
            >
            {{ weekday }}
      </div>
    </div>
    <DayRow />
  </div>
</template>

<script>
import DayRow from '@/components/Calendar/DayRow.vue'

export default {
  name: "WeekViews",
  components: {
    DayRow
  },
  props: {
    weekdays: {
      type: Array,
      default() {
        return ["Sun","Mon","Tue","Wed","Thu","Fri","Sat"];
      }
    }
  },
  methods: {}
};
</script>

<style scoped lang="scss">
.ds-month {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: white;

  .ds-week-header {
    display: flex;
    user-select: none;
  }

  .ds-week-header-day {
    flex: 1 0 20px;
    border-right: #e0e0e0 1px solid;
    color: #757575;
    padding: 4px;
    padding-bottom: 0px;

    &.ds-week-header-today {
      color: #4285f4;
      font-weight: 500;
    }
  }
}
</style>
