<template>
  <div class="ds-expand ds-calendar-app">
    <el-container>
      <el-header style="border-bottom: #e0e0e0 1px solid;">
        <el-row type="flex" class="ds-app-calendar-toolbar">
          <el-col :span="4">Calendar</el-col>
          <el-col :span="4">
            <el-button type="info" @click="handleClick">today</el-button>
          </el-col>
          <el-col :span="2">
            <el-button class="ds-skinny-button" icon="el-icon-arrow-left" circle></el-button>
          </el-col>
          <el-col :span="2">
            <el-button class="ds-skinny-button" icon="el-icon-arrow-right" circle></el-button>
          </el-col>
          <el-col :span="6" :offset="6">
            <el-select v-model="value" placeholder="Month">
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              ></el-option>
            </el-select>
          </el-col>
        </el-row>
      </el-header>
      <el-main class="ds-calendar-container">
        <Calendar />
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Calendar from '@/components/Calendar/Calendar.vue'

export default {
  name: 'CalendarApp',
  components: {
    Calendar
  },
  props: {
    events: {
      type: Array
    },
    readOnly: {
      type: Boolean,
      default: false
    }
  },
  created() {
    console.log('The time is ' + this.$moment().format('HH:mm'))
  },
  data() {
    return {
      activeIndex: '1',
      activeIndex2: '1',
      options: [
        {
          value: 'option1',
          label: 'Month'
        },
        {
          value: 'option2',
          label: 'Schedule'
        }
      ],
      value: ''
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log('123')
    },
    handleClick() {
      console.log('123')
    }
  }
}
</script>

<style lang="scss">
.ds-app-calendar-toolbar {
  .v-toolbar__content {
    border-bottom: 1px solid rgb(224, 224, 224);
  }
}

.ds-skinny-button {
  margin-left: 2px !important;
  margin-right: 2px !important;
}

.ds-expand {
  width: 100%;
  height: 100%;
}

.ds-calendar-container {
  padding: 0px !important;
  position: relative;
}

.v-btn--floating.ds-add-event-today {
  .v-icon {
    width: 24px;
    height: 24px;
  }
}
</style>
