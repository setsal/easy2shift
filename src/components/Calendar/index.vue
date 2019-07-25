<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="150px">
      <el-form-item label="你的 ID">
       {{ name }}
      </el-form-item>
      <el-form-item label="有需要備註什麼嗎?">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">送出</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
    <div class="block">
    <span class="demonstration">選擇日期</span>
      <el-date-picker
        :inline="true"
        type="dates"
        v-model="value4"
        placeholder="选择一个或多个日期">
      </el-date-picker>
    </div>
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
        name: ''
      },
      value4: ''
    }
  },
  props: { 
    month: Object
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
    }
  }
}
</script>

<style scoped>
</style>
