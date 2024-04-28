<script setup>
import { Delete, Edit } from '@element-plus/icons-vue'
import ChannelSelect from './components/ChannelSelect.vue'
import { bookDeleteService, bookGetListService } from '../../api/book.js'
import { ref } from 'vue'
import BookEdit from './components/BookEdit.vue'
const bookList = ref([]) //文章列表
const total = ref(0) //总条数
const loading = ref(false) //loading状态
// 定义请求参数对象
const params = ref({
  pagenum: 1, // 当前页
  pagesize: 5, // 当前生效的每页条数
  cate_id: '',
  input: ''
})

// 基于params参数获取文章列表
const getBookList = async () => {
  loading.value = true

  const res = await bookGetListService(params.value)
  bookList.value = res.data.data
  total.value = res.data.total
  loading.value = false
}
getBookList()

// 处理分页逻辑
const onSizeChange = (size) => {
  // 只要是每页条数变化，原本正在访问的当前页意义不大了,数据大概率已经不在原来那一页了
  // 重新从第一页渲染即可
  params.value.pagenum = 1
  params.value.pagesize = size
  // 基于最新的当前页和每页条数，渲染数据
  getBookList()
}

const onCurrentChange = (page) => {
  // 更新当前页
  params.value.pagenum = page
  // 基于最新的当前页和每页条数，渲染数据
  getBookList()
}

// 搜索逻辑  => 按最新的条件,重新检索，从第一页开始展示
const onSearch = () => {
  params.value.pagenum = 1 //重置页面
  getBookList()
}

// 重置逻辑 => 将筛选条件清空,重新检索，从第一页开始展示
const onReset = () => {
  params.value.pagenum = 1 //重置页面
  params.value.cate_id = ''
  params.value.input = ''
  getBookList()
}

const bookEditRef = ref()
// 添加逻辑
const onAddBook = () => {
  bookEditRef.value.open({})
}
// 编辑逻辑
const onEditBook = (row) => {
  bookEditRef.value.open(row)
}

// 删除逻辑
const onDeleteBook = (row) => {
  bookDeleteService(row.ID)
  ElMessage({
    message: '删除成功',
    type: 'success'
  })
  getBookList()
}

// 添加或者编辑成功的回调
const onSuccess = (type) => {
  if (type === 'add') {
    // 如果是添加,最好渲染最后一页
    const lastPage = Math.ceil((total.value + 1) / params.value.pagesize)
    // 更新成最大页码数, 在渲染
    params.value.pagenum = lastPage
    getBookList()
  }
  // 如果是编辑,直接渲染当前页
  getBookList()
}
</script>
<template>
  <page-container title="图书管理">
    <!--具名插槽配置按钮-->
    <template #extra>
      <el-button type="primary" @click="onAddBook">添加图书</el-button>
    </template>
    <!--表单区域-->
    <el-form inline>
      <el-form-item label="图书分类:">
        <!--Vue2 => v-model= :value 和 @input的简写-->
        <!--Vue3 => v-model= :modelValue 和 @update:modelValue的简写-->

        <!--Vue3 => v-model:cid   = :cid 和 @update:cid的简写-->
        <channel-select v-model="params.cate_id" width="200px"></channel-select>
      </el-form-item>
      <el-form-item>
        <el-input v-model="params.input" placeholder="请输入书名"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="onSearch" type="primary">搜索</el-button>
        <el-button @click="onReset">重置</el-button>
      </el-form-item>
    </el-form>
    <!--表格区域-->
    <el-table :data="bookList" v-loading="loading">
      <el-table-column label="图书编号" prop="ID"></el-table-column>
      <el-table-column label="名称" prop="bookname">
        <template #default="{ row }">
          <el-link type="primary" :underline="false">{{
            row.bookname
          }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="图片" prop="image">
        <template #default="{ row }">
          <el-image :src="row.image" style="width: 100px; height: 100px;" fit="cover"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="作者" prop="author"></el-table-column>
      <el-table-column label="类别" prop="cate_name"></el-table-column>
      <el-table-column label="出版社" prop="publisher"></el-table-column>
      <el-table-column label="数量" prop="number"></el-table-column>
      <!--利用作用域插槽row 可以获取当前行的数据  => v-for 遍历的 item-->
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button
            :icon="Edit"
            circle
            plain
            type="primary"
            @click="onEditBook(row)"
          ></el-button>
          <el-button
            :icon="Delete"
            circle
            plain
            type="danger"
            @click="onDeleteBook(row)"
          ></el-button>
        </template>
      </el-table-column>
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

    <!--添加编辑的抽屉-->
    <book-edit ref="bookEditRef" @success="onSuccess"></book-edit>
  </page-container>
</template>

<style lang="scss" scoped></style>
