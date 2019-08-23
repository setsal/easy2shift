<template>
  <div class="app-container">
      <aside>
      開放決定要排哪些月份的班<br />
      注意: 要先開放才能讓網管們排班!
    </aside>
    <el-form ref="form" :model="form" label-width="150px">
      <el-form-item label="選擇日期">
      <el-date-picker
        v-model="month"
        type="month"
        value-format="yyyy-MM"
        :picker-options="pickerOptions"
        placeholder="選擇月份">
      </el-date-picker>
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
  name: 'Manage',
  computed: {
    ...mapGetters([
      'name',
      'roles'
    ])
  },
  data() {
    return {
      form:{
      },
      month: '',
      loading: false,
      pickerOptions: {
        disabledDate: (time) => {
          return this.dealDisabledDate(time)
        }
      }
    }
  },
  methods: {
      onSubmit() {  
        // submit request
        this.loading = true
        this.$store.dispatch('admin/month', this.month).then(() => {
          this.$message( this.month + ' 已經開放排班~')
          this.loading = false
        }).catch(() => {
          this.$message( 'somethin error : (')
        })           
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
        return time.getTime() < startDate 
      }      
    }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
