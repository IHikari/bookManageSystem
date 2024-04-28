<script setup>
// 在vue3 CompositionAPI中
// 1.获取路由对象router useRouter
// const router = useRouter()
// 2.获取路由参数 route useRoute
// const route = useRoute()
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
</script>

<template>
  <div>
    <!--App.vue只需要留一个路由出口router-view即可-->
    <el-config-provider :locale="zhCn">
      <router-view></router-view>
    </el-config-provider>
  </div>
</template>

<style scoped></style>
