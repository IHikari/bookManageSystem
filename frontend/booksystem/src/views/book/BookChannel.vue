<script setup>
import { ref } from 'vue'
import {
  bookGetChannelsService,
  bookDeleteChannelService
} from '../../api/book.js'
import { Delete, Edit } from '@element-plus/icons-vue'
import ChannelEdit from './components/ChannelEdit.vue'
const channelList = ref([])
const loading = ref(false)
const dialog = ref()
const getChannelList = async () => {
  loading.value = true
  const res = await bookGetChannelsService()
  channelList.value = res.data.data
  loading.value = false
}
// 请求分类数据
getChannelList()

const onDeleteChannel = async (row) => {
  await ElMessageBox.confirm('你确定要删除该分类吗', '温馨提示', {
    type: 'warning',
    confirmButtonText: '确认',
    cancelButtonText: '取消'
  })
  await bookDeleteChannelService(row.ID)
  getChannelList()
  ElMessage({
    message: '删除成功',
    type: 'success'
  })
}
const onEditChannel = (row) => {
  dialog.value.open(row)
}
const onAddChannel = () => {
  dialog.value.open({})
}
const onSuccess = () => {
  getChannelList()
}
</script>
<template>
  <page-container title="图书分类">
    <!--具名插槽配置按钮-->
    <template #extra>
      <el-button type="primary" @click="onAddChannel">添加分类</el-button>
    </template>
    <el-table v-loading="loading" :data="channelList" style="width: 100%">
      <el-table-column type="index" label="序号" width="100"></el-table-column>
      <el-table-column prop="ID" label="分类编号"></el-table-column>
      <el-table-column prop="name" label="分类名称"></el-table-column>
      <el-table-column prop="alias" label="分类别名"></el-table-column>
      <el-table-column label="操作" width="100">
        <!--row就是channelList的每一项，$index是下标-->
        <template #default="{ row, $index }">
          <!--:icon 图标  circle 圆形  plain 镂空-->
          <el-button
            :icon="Edit"
            circle
            type="primary"
            plain
            @click="onEditChannel(row)"
          ></el-button>
          <el-button
            :icon="Delete"
            circle
            type="danger"
            plain
            @click="onDeleteChannel(row, $index)"
          ></el-button>
        </template>
      </el-table-column>
      <!--没有数据显示-->
      <template #empty>
        <el-empty description="没有数据"></el-empty>
      </template>
    </el-table>
    <channel-edit ref="dialog" @success="onSuccess"></channel-edit>
  </page-container>
</template>

<style lang="scss" scoped></style>
