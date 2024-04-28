<script setup>
import { ref } from 'vue'
import { userEditService, userAddService } from '../../../api/user.js'

const dialogVisible = ref(false)
const formRef = ref()
const formModel = ref({
  ID: 0,
  name: '',
  password: '',
  permission: '',
  CreatedAt: ''
})
const rules = {
  ID: [
    { required: true, message: '请输入学工号', trigger: 'blur' },
    { pattern: /^[0-9]*$/, message: '学工号必须是数字', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' },
    {
      pattern: /^([\u4e00-\u9fa5]{2,5}|[a-zA-Z.\s]{2,20})$/,
      message: '姓名必须是(2-5)位中文汉字或(2-20)位英文字母',
      trigger: 'blur'
    }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      pattern: /^\S{6,15}$/,
      message: '密码长度必须是6-15位的非空字符串',
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
    await userEditService(formModel.value)
    ElMessage({
      message: '编辑成功',
      type: 'success'
    })
  } else {
    formModel.value.ID = parseInt(formModel.value.ID)
    if (formModel.value.permission === ''){
      formModel.value.permission = 'false'
    }
    await userAddService(formModel.value)
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
  console.log(formModel.value)
}

// 向外暴露方法
defineExpose({
  open
})

</script>

<template>
  <el-dialog
    v-model="dialogVisible"
    :title="formModel.CreatedAt ? '编辑用户' : '添加用户'"
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
      <el-form-item label="学工号" prop="ID">
        <el-input
          v-model="formModel.ID"
          :disabled="!!formModel.CreatedAt"
          placeholder="请输入学工号"
        ></el-input>
      </el-form-item>
      <el-form-item label="姓名" prop="name">
        <el-input
          v-model="formModel.name"
          :disabled="!!formModel.CreatedAt"
          placeholder="请输入姓名"
        ></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="formModel.password" placeholder="请输入密码" type="password">
        </el-input>
      </el-form-item>
      <el-form-item label="权限" prop="permission">
        <el-radio-group v-model="formModel.permission">
          <el-radio :label="true">管理员</el-radio>
          <el-radio :label="false" :disabled="formModel.permission">普通用户</el-radio>
        </el-radio-group>
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
