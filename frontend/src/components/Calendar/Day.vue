<template>
  <div class="ds-day">
    <div>
      <a class="ds-dom"></a>
      <!-- <span class="ds-first-day" v-if="showMonth">
        {{ month }}
      </span>-->
    </div>

    <!-- //event must be here -->
    <!-- <template v-for="(event, i) in visibleEvents"> -->
    <!-- <ds-calendar-event
    ></ds-calendar-event>-->
    <!-- </template> -->
  </div>
</template>

<script>
export default {
  name: "Day",

  props: {},

  computed: {},

  methods: {}
};
</script>

<style scoped lang="scss">
.ds-day {
  flex: 1;
  width: 0;
  border-right: #e0e0e0 1px solid;
  border-bottom: #e0e0e0 1px solid;
  padding: 4px;
  overflow: hidden;
  user-select: none;

  .ds-dom {
    text-decoration: none;
    color: #212121;
    width: 24px;
    height: 24px;
    line-height: 24px;
    text-align: center;
    user-select: none;

    &:hover {
      text-decoration: underline;
    }

    &.ds-today-dom {
      border-radius: 12px;
      background-color: #4285f4;
      color: white;
      display: inline-block;
      position: relative;
      z-index: 1;
    }
  }
}

.ds-out-calendar {
  color: #757575;
}
</style>
