<script setup>
import { ref } from 'vue'
import ChannelSelect from '../../../views/book/components/ChannelSelect.vue'
import { Plus } from '@element-plus/icons-vue'
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css'
import { bookAddService, bookGetDetailService, bookEditService } from '../../../api/book.js'

// 抽屉显示隐藏
const visibleDrawer = ref(false)

// 默认数据
const defaultForm = {
  bookname: '', // 书名
  author: '', // 作者
  cate_id: '', // 类别id
  publisher: '', // 出版社
  image: '', // 封面图  file对象
  number: '', // 数量
  remarks: '' // 书本描述
}

// 准备数据
const formModel = ref({ ...defaultForm })

// 图片上传相关逻辑
const imageUrl = ref('')
const onSelectFile = (uploadFile) => {
  imageUrl.value = URL.createObjectURL(uploadFile.raw) //预览图片

  // 立即将图片对象，存入 formModel.value.image将来用于提交
  formModel.value.image = uploadFile.raw
}
// 添加
const emit = defineEmits(['success'])
const onAdd = async () => {
  // 注意：当前接口,需要的是formData对象
  // 将普通对象转换成  formData对象
  const fd = new FormData()
  for (let key in formModel.value) {
    fd.append(key, formModel.value[key])
  }

  // 发请求
  if (formModel.value.ID) {
    // 编辑操作
    await bookEditService(fd)
    ElMessage({
      message: '修改成功',
      type: 'success'
    })
    visibleDrawer.value = false
    // 通知到父组件,添加成功了
    emit('success', 'edit')
  } else {
    // 添加操作
    await bookAddService(fd)
    ElMessage({
      message: '添加成功',
      type: 'success'
    })
    visibleDrawer.value = false
    // 通知到父组件,添加成功了
    emit('success', 'add')
  }
}
const onCancel = () => {
  visibleDrawer.value = false
}
// 表单校验规则
const rules = {
  bookname: [
    {
      required: true,
      message: '请输入图书名称',
      trigger: 'blur'
    }
  ],
  cate_id: [
    {
      required: true,
      message: '请输入图书名称',
      trigger: 'blur'
    }
  ],
  author: [
    {
      required: true,
      message: '请输入作者名字',
      trigger: 'blur'
    },
    {
      pattern: /^[\u4E00-\u9FA5A-Za-z\s]+(·[\u4E00-\u9FA5A-Za-z]+)*$/,
      message: '请输入正确格式的人名',
      trigger: 'blur'
    }
  ],
  publisher: [
    {
      required: true,
      message: '请输入出版社名称',
      trigger: 'blur'
    },
    {
      pattern: /^[\u4E00-\u9FA5A-Za-z\s]+(·[\u4E00-\u9FA5A-Za-z]+)*$/,
      message: '请输入正确格式的出版商名称',
      trigger: 'blur'
    }
  ],
  number: [
    {
      required: true,
      message: '请输入图书数量',
      trigger: 'blur'
    },
    {
      pattern: /^[1-9]\d*|0$/,
      message: '图书数量只能为数字',
      trigger: 'blur'
    }
  ],
  image: [
    {
      required: true,
      message: '请上传图书图片',
      trigger: 'blur'
    }
  ]
}

// 组件对外暴露一个方法 open, 基于open传来的参数,区分添加还是编辑
// open({}) => 表单无需渲染,说明是添加
// open({ID,...,...})  => 表单需要渲染,说明是编辑
// open调用后，可以打开抽屉

const editorRef = ref()

const open = async (row) => {
  visibleDrawer.value = true //显示抽屉
  if (row.ID) {
    // 需要基于row.ID发送请求，获取编辑对应的详情数据，进行回显
    const res = await bookGetDetailService(row.ID)
    formModel.value = res.data.data

    // 图片需要单独处理回显
    imageUrl.value = formModel.value.image
    editorRef.value.setText('')
  } else {
    formModel.value = { ...defaultForm } // 基于默认的数据,重置form数据
    // 这里重置了表单数据，但是图片上传img地址,富文本编辑器  => 需要手动重置
    imageUrl.value = ''
    editorRef.value.setText('')
  }
}

// 暴露open方法
defineExpose({
  open
})
</script>

<template>
  <el-drawer
    v-model="visibleDrawer"
    :title="formModel.ID ? '编辑图书' : '添加图书'"
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
        <el-input
          v-model="formModel.bookname"
          placeholder="请输入名称"
        ></el-input>
      </el-form-item>
      <el-form-item label="作者" prop="author">
        <el-input
          v-model="formModel.author"
          placeholder="请输入作者"
        ></el-input>
      </el-form-item>
      <el-form-item label="图书分类" prop="cate_id">
        <channel-select
          v-model="formModel.cate_id"
          width="100%"
        ></channel-select>
      </el-form-item>
      <el-form-item label="出版社" prop="publisher">
        <el-input
          v-model="formModel.publisher"
          placeholder="请输入出版社名称"
        ></el-input>
      </el-form-item>
      <el-form-item label="数量" prop="number">
        <el-input
          v-model="formModel.number"
          placeholder="请输入图书数量"
        ></el-input>
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
        >
          <img v-if="imageUrl" :src="imageUrl" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="图书描述" prop="remarks">
        <div class="editor">
          <quill-editor
            ref="editorRef"
            v-model:content="formModel.remarks"
            content-type="text"
            theme="snow"
          >
          </quill-editor>
        </div>
      </el-form-item>
      <el-form-item>
        <el-button @click="onAdd()" type="primary">提交</el-button>
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
