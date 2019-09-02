<template>
<div class="calendar-event">
    <div class="ds-calendar-event">
      <span >
          <strong class="ds-ev-title">admin</strong>
      </span>
    </div>
</div>
</template>

<script>
export default {
  name: "CalendarEvent",

  props: {},

  computed: {},

  methods: {}
};
</script>

<style scoped lang="scss">
.ds-calendar-event {
  margin: 1px;
  color: white;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  padding-left: 0.5em;
  font-size: 12px;
  cursor: pointer;
  position: absolute;
  user-select: none;
  display: block;
  margin: 0;
  border-radius: 2px;
  pointer-events: all;
  
  top: 0px;
  background-color: rgb(33, 150, 243);
  color: rgb(255, 255, 255);
  left: 0px;
  right: 0px;
  text-decoration: inherit;
}

.calendar-event {
  width: 100%;
  height: 18px;
  pointer-events: none;
  position: relative;
  line-height: 1.5;
}
</style>
