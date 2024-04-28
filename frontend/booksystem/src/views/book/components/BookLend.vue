<script setup>
import { ref } from 'vue'
import ChannelSelect from '../../../views/book/components/ChannelSelect.vue'
import { Plus } from '@element-plus/icons-vue'
import { bookLendService } from '../../../api/book.js'
import { userBookGetDetailService } from '../../../api/user.js'
import { useUserStore } from '../../../stores/index.js'

// 抽屉显示隐藏
const visibleDrawer = ref(false)
const userStore = useUserStore()
// 默认数据
const defaultForm = {
  ID: '',
  bookname: '', // 书名
  author: '', // 作者
  cate_id: '', // 类别id
  publisher: '', // 出版社
  image: '', // 封面图  file对象
  number: '', // 数量
  book_num: '' //借阅数量
}
// 准备数据
const formModel = ref({ ...defaultForm })
const params = ref({
  book_id: '',
  school_id: '',
  book_num: ''
})
// 图片上传相关逻辑
const imageUrl = ref('')

const checkBookNum = (rule, value, cb) => {
  if (value > formModel.value.number) {
    cb(new Error('归还数量已超过持有上限'))
  } else {
    cb()
  }
}
const rules = {
  book_num: [
    {
      required: true,
      message: '请输入归还数量',
      trigger: 'blur'
    },
    { pattern: /^\d+$/, message: '只能输入数字', trigger: 'blur' },
    { validator: checkBookNum, trigger: 'blur' }
  ]
}
const onSelectFile = (uploadFile) => {
  imageUrl.value = URL.createObjectURL(uploadFile.raw) //预览图片

  // 立即将图片对象，存入 formModel.value.image将来用于提交
  formModel.value.image = uploadFile.raw
}
// 添加
const emit = defineEmits(['success'])
const onLend = async () => {
  // 注意：当前接口,需要的是formData对象
  // 将普通对象转换成  formData对象
  params.value.book_id = formModel.value.ID
  params.value.book_num = formModel.value.book_num
  params.value.school_id = userStore.user.ID
  const fd = new FormData()
  for (let key in params.value) {
    fd.append(key, params.value[key])
  }
  // 发请求
  // 编辑操作
  await bookLendService(fd)
  ElMessage({
    message: '还书成功',
    type: 'success'
  })
  visibleDrawer.value = false
  // 通知到父组件,借书成功了
  emit('success')
}

const onCancel = () => {
  visibleDrawer.value = false
}
// 表单校验规则
// 组件对外暴露一个方法 open, 基于open传来的参数,区分添加还是编辑
// open({}) => 表单无需渲染,说明是添加
// open({ID,...,...})  => 表单需要渲染,说明是编辑
// open调用后，可以打开抽屉

const open = async (row) => {
  visibleDrawer.value = true //显示抽屉
  // 需要基于row.ID发送请求，获取编辑对应的详情数据，进行回显
  const res = await userBookGetDetailService(row.ID)
  formModel.value = res.data.data

  // 图片需要单独处理回显
  imageUrl.value = formModel.value.image
}

// 暴露open方法
defineExpose({
  open
})
</script>

<template>
  <el-drawer
    v-model="visibleDrawer"
    title="归还图书"
    direction="rtl"
    size="50%"
  >
    <!-- 新增图书表单 -->
    <el-form
      :model="formModel"
      :rules="rules"
      ref="formRef"
      label-width="100px"
    >
      <el-form-item label="图书名称" prop="bookname">
        <el-input v-model="formModel.bookname" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="作者" prop="author">
        <el-input v-model="formModel.author" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="图书分类" prop="cate_id">
        <channel-select
          v-model="formModel.cate_id"
          width="100%"
          :disabled="true"
        ></channel-select>
      </el-form-item>
      <el-form-item label="出版社" prop="publisher">
        <el-input v-model="formModel.publisher" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="数量" prop="number">
        <el-input v-model="formModel.number" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="图书封面" prop="image">
        <!--此处需要关闭element-plus的自动上传,不需要配置action等参数
            只需要做前端的本地预览图片即可，无需在提交前上次图标
            语法：URL.createObjectURL(...)创建本地预览的地址，来预览-->
        <el-upload
          class="avatar-uploader"
          :show-file-list="false"
          :auto-upload="false"
          :on-change="onSelectFile"
          :disabled="true"
        >
          <img v-if="imageUrl" :src="imageUrl" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="归还数量" prop="book_num">
        <el-input
          v-model="formModel.book_num"
          placeholder="请输入归还数量"
          type="number"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="onLend()" type="primary">提交</el-button>
        <el-button @click="onCancel()" type="info">取消</el-button>
      </el-form-item>
    </el-form>
  </el-drawer>
</template>

<style lang="scss" scoped>
.avatar-uploader {
  :deep {
    .avatar {
      width: 178px;
      height: 178px;
      display: block;
    }
    .el-upload {
      border: 1px dashed var(--el-border-color);
      border-radius: 6px;
      cursor: pointer;
      position: relative;
      overflow: hidden;
      transition: var(--el-transition-duration-fast);
    }
    .el-upload:hover {
      border-color: var(--el-color-primary);
    }
    .el-icon.avatar-uploader-icon {
      font-size: 28px;
      color: #8c939d;
      width: 178px;
      height: 178px;
      text-align: center;
    }
  }
}
.editor {
  width: 100%;
  :deep(.ql-editor) {
    min-height: 110px;
  }
}
</style>
