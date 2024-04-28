<script setup>
import { Delete, Edit } from '@element-plus/icons-vue'
import { userGetListService, userDeleteService } from '../../api/user.js'
import { ref } from 'vue'
import UserEdit from './components/UserEdit.vue'
const userList = ref([]) //文章列表
const total = ref(0) //总条数
const loading = ref(false) //loading状态
// 定义请求参数对象
const params = ref({
  pagenum: 1, // 当前页
  pagesize: 5, // 当前生效的每页条数
  input: ''
})
const dialog = ref()
// 基于params参数获取文章列表
const getUserList = async () => {
  loading.value = true

  const res = await userGetListService(params.value)
  userList.value = res.data.data
  total.value = res.data.total
  loading.value = false
}
getUserList()

// 处理分页逻辑
const onSizeChange = (size) => {
  // 只要是每页条数变化，原本正在访问的当前页意义不大了,数据大概率已经不在原来那一页了
  // 重新从第一页渲染即可
  params.value.pagenum = 1
  params.value.pagesize = size
  // 基于最新的当前页和每页条数，渲染数据
  getUserList()
}

const onCurrentChange = (page) => {
  // 更新当前页
  params.value.pagenum = page
  // 基于最新的当前页和每页条数，渲染数据
  getUserList()
}

// 搜索逻辑  => 按最新的条件,重新检索，从第一页开始展示
const onSearch = () => {
  params.value.pagenum = 1 //重置页面
  getUserList()
}

// 重置逻辑 => 将筛选条件清空,重新检索，从第一页开始展示
const onReset = () => {
  params.value.pagenum = 1 //重置页面
  params.value.input = ''
  getUserList()
}


// 添加逻辑
const onAddUser = () => {
  dialog.value.open({})
}
// 编辑逻辑
const onEditUser = (row) => {
  row.password = ''
  dialog.value.open(row)
}

// 删除逻辑
const onDeleteUser = async (row) => {
  await userDeleteService(row.ID)
  ElMessage({
    message: '删除成功',
    type: 'success'
  })
  getUserList()
}

// 添加或者编辑成功的回调
const onSuccess = (type) => {
  if (type === 'add') {
    // 如果是添加,最好渲染最后一页
    const lastPage = Math.ceil((total.value + 1) / params.value.pagesize)
    // 更新成最大页码数, 在渲染
    params.value.pagenum = lastPage
    getUserList()
  }
  // 如果是编辑,直接渲染当前页
  getUserList()
}
</script>
<template>
  <page-container title="用户列表">
    <!--具名插槽配置按钮-->
    <template #extra>
      <el-button type="primary" @click="onAddUser">新增用户</el-button>
    </template>
    <!--表单区域-->
    <el-form inline>
      <el-form-item>
        <el-input v-model="params.input" placeholder="请输入学工号或姓名"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="onSearch" type="primary">搜索</el-button>
        <el-button @click="onReset">重置</el-button>
      </el-form-item>
    </el-form>
    <!--表格区域-->
    <el-table :data="userList" v-loading="loading">
      <el-table-column label="学工号" prop="ID"></el-table-column>
      <el-table-column label="姓名" prop="name"></el-table-column>
      <el-table-column label="昵称" prop="nickname"></el-table-column>
      <el-table-column label="邮箱" prop="email"></el-table-column>
      <el-table-column label="权限" prop="permission">
        <template v-slot="scope">
          <span v-if="scope.row.permission">管理员</span>
          <span v-else>普通用户</span>
        </template>
      </el-table-column>
      <!--利用作用域插槽row 可以获取当前行的数据  => v-for 遍历的 item-->
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button
            :icon="Edit"
            circle
            plain
            type="primary"
            @click="onEditUser(row)"
          ></el-button>
          <el-button
            :icon="Delete"
            circle
            plain
            type="danger"
            @click="onDeleteUser(row)"
          ></el-button>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty description="没有数据"></el-empty>
      </template>
    </el-table>


    <!--分页区域-->
    <el-pagination
      v-model:current-page="params.pagenum"
      v-model:page-size="params.pagesize"
      :page-sizes="[3, 5, 7, 10]"
      :background="true"
      layout="jumper, total, sizes, prev, pager, next"
      :total="total"
      @size-change="onSizeChange"
      @current-change="onCurrentChange"
      style="margin-top: 20px; justify-content: flex-end"
    />
    <user-edit ref="dialog" @success="onSuccess"></user-edit>
  </page-container>
</template>

<style lang="scss" scoped></style>
