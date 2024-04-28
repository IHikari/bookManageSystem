<script setup>
import { useUserStore } from '../../stores'
import { ref, computed} from 'vue'
import { userEditInfoService } from '../../api/user.js'
const {
  user: {ID, name, nickname, email,permission },
  getUser
} = useUserStore()

const userInfo = ref({ ID, name, nickname, email,permission })
const formRef = ref()
const rules = {
  nickname: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' },
    {
      pattern: /^\S{2,10}$/,
      message: '昵称必须是2-10位的非空字符串',
      trigger: 'blur'
    }
  ],
  email: [
    { required: true, message: '请输入用户邮箱', trigger: 'blur' },
    { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }
  ]
}
const submitForm =async () =>{
  // 等待校验结果
  await formRef.value.validate()
  // 提交
  await userEditInfoService(userInfo.value)
  // 通知user模块,进行数据更新
  getUser()
  ElMessage({
    message: '编辑成功',
    type: 'success'
  })
}
// 用户权限
const userPermissionText = computed(() => {
  return userInfo.value.permission ? '管理员' : '普通用户'
})
</script>

<template>
  <page-container title="基本资料">
    <el-row>
      <el-col :span="12">
        <el-form
          :model="userInfo"
          :rules="rules"
          ref="formRef"
          label-width="100px"
          size="large"
        >
          <el-form-item label="学工号">
            <el-input v-model="userInfo.ID" disabled></el-input>
          </el-form-item>
          <el-form-item label="登录名称">
            <el-input v-model="userInfo.name" disabled></el-input>
          </el-form-item>
          <el-form-item label="用户权限">
            <el-input v-model="userPermissionText" disabled></el-input>
          </el-form-item>
          <el-form-item label="用户昵称" prop="nickname">
            <el-input v-model="userInfo.nickname"></el-input>
          </el-form-item>
          <el-form-item label="用户邮箱" prop="email">
            <el-input v-model="userInfo.email"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm">提交修改</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </page-container>
</template>