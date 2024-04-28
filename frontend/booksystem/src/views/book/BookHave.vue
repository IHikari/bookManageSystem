<script setup>
import { userGetBookListService } from '../../api/user.js'
import { ref } from 'vue'
import BookLend from './components/BookLend.vue'

const bookList = ref([]) //文章列表
const total = ref(0) //总条数
const loading = ref(false) //loading状态

const getUserBookList = async () => {
  loading.value = true
  const res = await userGetBookListService()
  bookList.value = res.data.data
  total.value = res.data.total
  loading.value = false
}
getUserBookList()

const bookLendRef = ref()


const onLendBook = (row) => {
  console.log(row)
  bookLendRef.value.open(row)
}
const onSuccess = () => {
  getUserBookList()
}
</script>
<template>
  <page-container title="我的图书">
    <!--表格区域-->
    <el-table :data="bookList" v-loading="loading">
      <el-table-column label="图书编号" prop="ID"></el-table-column>
      <el-table-column label="图书名称" prop="bookname"></el-table-column>
      <el-table-column label="图书图片" prop="image">
        <template #default="{ row }">
          <el-image :src="row.image" style="width: 100px; height: 100px;" fit="cover"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="作者" prop="author"></el-table-column>
      <el-table-column label="类别" prop="cate_name"></el-table-column>
      <el-table-column label="出版社" prop="publisher"></el-table-column>
      <el-table-column label="持有数量" prop="number"></el-table-column>
      <!--利用作用域插槽row 可以获取当前行的数据  => v-for 遍历的 item-->
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button
            plain
            type="success"
            @click="onLendBook(row)"
          >还书</el-button>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty description="没有数据"></el-empty>
      </template>
    </el-table>
    <!--添加编辑的抽屉-->
    <book-lend ref="bookLendRef" @success="onSuccess"></book-lend>
  </page-container>
</template>

<style lang="scss" scoped></style>
