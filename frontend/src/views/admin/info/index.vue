<template>
  <div class="dashboard-container">
    <div class="dashboard-text">Admin page </div>
    <div class="dashboard-text">roles need admin to visit. <br />your roles: <span v-for="role in roles" :key="role">{{ role }}</span></div>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Info',
  computed: {
    ...mapGetters([
      'name',
      'roles'
    ])
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
