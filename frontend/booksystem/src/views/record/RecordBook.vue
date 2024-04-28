<script setup>
import { ref } from 'vue'
import { recordGetService } from '../../api/record.js'
const recordList = ref([])
const loading = ref(false)
const formatCreatedAt = (row, column) => {
  const timeString = row[column.property]; // 获取时间字符串
  const date = new Date(timeString); // 解析为日期对象

  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`; // 返回格式化后的时间字符串
}
const getRecordList = async () => {
  loading.value = true
  const res = await recordGetService()
  recordList.value = res.data.data
  loading.value = false
}
// 请求分类数据
getRecordList()
</script>
<template>
  <page-container title="借阅记录">
    <el-table v-loading="loading" :data="recordList" style="width: 100%">
      <el-table-column type="index" label="序号" width="100"></el-table-column>
      <el-table-column prop="school_id" label="学工号"></el-table-column>
      <el-table-column prop="name" label="姓名"></el-table-column>
      <el-table-column prop="book_id" label="图书编号"></el-table-column>
      <el-table-column prop="book_name" label="图书书名"></el-table-column>
      <el-table-column prop="created_at" label="时间" :formatter="formatCreatedAt"></el-table-column>
      <el-table-column prop="book_num" label="数量"></el-table-column>
      <el-table-column prop="is_borrow" label="状态">
        <template v-slot="scope">
          <span v-if="scope.row.is_borrow">借书</span>
          <span v-else>还书</span>
        </template>
      </el-table-column>
      <!--没有数据显示-->
      <template #empty>
        <el-empty description="没有数据"></el-empty>
      </template>
    </el-table>
  </page-container>
</template>

<style lang="scss" scoped></style>
