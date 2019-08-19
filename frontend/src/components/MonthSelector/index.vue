<template>
  <div class="monthselector">
    <div class="selector">
      <button @click="subMonth" class="changeMonthButton">◀</button>
      <div class="title">{{ month.format('YYYY-MM')}}</div>
      <button @click="addMonth" class="changeMonthButton">▶</button>
      </div>
  </div>
</template>

<script>

export default {
  name: 'MonthSelector',
  data() {
    return {
      
    }
  },
  props: { 
    month: { type: Object },
    onChange: { type: Function }
  },
  methods: {
    addMonth() {
      const newMonth = this.month.clone().add(1,'M')
      this.onChange(newMonth)
    },
    subMonth() {
      const newMonth = this.month.clone().subtract(1,'M')
      this.onChange(newMonth)
    }
  }
}
</script>

<style scoped>
.changeMonthButton {
  border-radius: 50%;
  border: 0;
  background-color: rgb(34,34,34);
  box-shadow: 2px 2px 4px 1px rgba(0,0,0,0.1);
  color: white;
  font-weight: 900;
  width: 30px;
  height: 30px;
  cursor: pointer;  
}
.title {
  font-size: 20px;
}
</style>

<style  lang="scss">
.monthselector {
  height: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
  .selector {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 200px;
    height: 100%;
  }
}
</style>
