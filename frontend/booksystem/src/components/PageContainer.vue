<script setup>
defineProps({
  title:{
    required:true,
    type:String
  }
})
</script>
<template>
  <el-card class="page-container">
    <template #header>
      <div class="header">
        <span>{{ title }}</span>
        <div class="extra">
          <!--考虑复用,具名插槽定义按钮-->
          <slot name="extra"></slot>
        </div>
      </div>
    </template>

    <!--考虑复用,内容不要写死,使用插槽slot-->
    <slot> </slot>
  </el-card>
</template>

<style lang="scss" scoped>
.page-container {
  min-height: 100%;
  box-sizing: border-box;

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}
</style>
<script setup></script>
