<template>
  <div class="app-container schedule" style="">
    <CalendarApp />
  </div>
</template>

<script>
import CalendarApp from '@/components/Calendar/CalendarApp.vue'

export default {
  components: {
    CalendarApp
  },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
    }
  }
}
</script>

<style lang="scss" scoped>
.schedule {
  display: flex;
  min-height: calc(100vh - 50px);
}
</style>