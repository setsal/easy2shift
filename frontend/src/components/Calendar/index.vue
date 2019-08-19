<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="150px">
      <el-form-item label="你的 ID">
       {{ name }}
      </el-form-item>
      <el-form-item label="選擇日期">
        <el-date-picker
        type="dates"
        v-model="form.activeDays"
        :picker-options="pickerOptions"
        placeholder="選擇日期">
      </el-date-picker>
      </el-form-item>
      <el-form-item label="有需要備註什麼嗎?">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">送出</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Calendar',
  computed: {
    ...mapGetters([
      'name',
      'roles'
    ])
  },  
  data() {
    return {
      form: {
        name: '',
        activeDays: '',
        desc: ''
      },
      pickerOptions: {
        disabledDate: (time) => {
          return this.dealDisabledDate(time)
        }
      }
    }
  },
  props: { 
    month: Object
  },
  methods: {
    onSubmit() {
      console.log(this.name)
      this.$message( this.name + ' 填好班表囉~')
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })   
    },
    dealDisabledDate(time) {
      let currDate = new Date()
      currDate.setMonth(currDate.getMonth() + 1, 0 )
      let startDate = currDate.getTime()
      // console.log(currDate.getMonth())
      currDate.setMonth(currDate.getMonth() + 2, 0 )
      let endDate = currDate.getTime()
      return time.getTime() < startDate || time.getTime() > endDate || time.getDay() === 6 || time.getDay() === 0
    }
  }
}
</script>

<style scoped>
</style>
