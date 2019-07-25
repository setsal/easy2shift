<template>
  <div class="app-container">
    <MonthSelector :month="month" :onChange="onChangeMonth" />
    <Calendar :month="month" />
  </div>
</template>

<script>
import MonthSelector from '@/components/MonthSelector'
import Calendar from '@/components/Calendar'
import moment from 'moment'

export default {
  components: {
    MonthSelector,
    Calendar
  },
  data() {
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      },
      month: moment().add(1, 'M')
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    },
    onChangeMonth(value) {
      this.month = value
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

