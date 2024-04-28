<script setup>
import { ref } from 'vue'
import { bookEditChannelService, bookAddChannelService } from '../../../api/book.js'


const dialogVisible = ref(false)
const formRef = ref()
const formModel = ref({
  ID: 0,
  name: '',
  alias: '',
  CreatedAt: ''
})
const rules = {
  ID: [
    {
      required: true,
      message: '请输入分类编号',
      trigger: 'change'
    },
    {
      pattern: /^[0-9]{3,6}$/,
      message: '分类编号必须是3-6位的数字',
      trigger: 'blur'
    }
  ],
  name: [
    {
      required: true,
      message: '请输入分类名称',
      trigger: 'blur'
    },
    {
      pattern: /^\S{1,10}$/,
      message: '分类名必须是1-10位的非空字符',
      trigger: 'blur'
    }
  ],
  alias: [
    {
      required: true,
      message: '请输入分类别名',
      trigger: 'blur'
    },
    {
      pattern: /^[a-zA-Z0-9]{1,15}$/,
      message: '分类名必须是1-15位的字母或数字',
      trigger: 'blur'
    }
  ]
}

const emit = defineEmits(['success'])
// 提交请求
const onSubmit = async () => {
  await formRef.value.validate()
  const isEdit = formModel.value.CreatedAt
  if (isEdit) {
    formModel.value.ID = parseInt(formModel.value.ID)
    await bookEditChannelService(formModel.value)
    ElMessage({
      message: '编辑成功',
      type: 'success'
    })
  } else {
    formModel.value.ID = parseInt(formModel.value.ID)
    await bookAddChannelService(formModel.value)
    ElMessage({
      message: '添加成功',
      type: 'success'
    })
  }
  dialogVisible.value = false
  emit('success')
}

// 组件对外暴露一个方法 open, 基于open传来的参数,区分添加还是编辑
// open({}) => 表单无需渲染,说明是添加
// open({name,ID,...})  => 表单需要渲染,说明是编辑
// open调用后，可以打开弹窗

const open = (row) => {
  dialogVisible.value = true
  formModel.value = { ...row } //添加 -> 重置了表单内容   编辑 -> 回写了表单数据
}

// 向外暴露方法
defineExpose({
  open
})
</script>

<template>
  <el-dialog
    v-model="dialogVisible"
    :title="formModel.CreatedAt ? '编辑分类' : '添加分类'"
    width="30%"
    :before-close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="formModel"
      :rules="rules"
      label-width="100px"
      style="padding-right: 30px"
    >
      <el-form-item label="分类编号" prop="ID">
        <el-input v-model="formModel.ID" placeholder="请输入分类编号" :disabled="formModel.CreatedAt">
        </el-input>
      </el-form-item>
      <el-form-item label="分类名称" prop="name">
        <el-input v-model="formModel.name" placeholder="请输入分类名称">
        </el-input>
      </el-form-item>
      <el-form-item label="分类别名" prop="alias">
        <el-input v-model="formModel.alias" placeholder="请输入分类别名">
        </el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="onSubmit"> 确认 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped></style>
