<script setup>
import { userRegisterService, userLoginService } from '../../api/user.js'
import { User, Lock, Postcard } from '@element-plus/icons-vue'
import { ref, watch } from 'vue'
import { useUserStore } from '../../stores/index.js'
import { useRouter } from 'vue-router'

const isRegister = ref(false)
const form = ref()
// 整个的用于提交的form数据对象
const formModel = ref({
  ID: '',
  name: '',
  password: '',
  repassword: ''
})
// 整个表单的校验规则
// 1.非空校验 required: true 非空校验  message:消息提示   trigger:触发校验时机  blur:失焦  change:改变
// 2.长度校验 min:**, max: **
// 3.正则校验 pattern: 正则校验
// 4.自定义校验  =>  自己写逻辑
// validator: (rule, value, callback)
// (1)rule 当前校验规则相关的信息
// (2)value 所校验的表单元素目前的表单值
// (3)callback 无论成功失败，都需要 callback 回调
//     -callback()校验成功
//     -callback(new Error(错误信息)) 校验失败
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
      message: '密码必须是6-15位的非空字符',
      trigger: 'blur'
    }
  ],
  repassword: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      pattern: /^\S{6,15}$/,
      message: '密码必须是6-15位的非空字符',
      trigger: 'blur'
    },
    // 自定义校验
    {
      validator: (rule, value, callback) => {
        //判断 value 和当前form中手机的password是否一致
        if (value !== formModel.value.password) {
          callback(new Error('两次密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}
// 登录逻辑
const userStore = useUserStore()
const router = useRouter()
const login = async () => {
  await form.value.validate()
  formModel.value.ID = parseInt(formModel.value.ID)
  const res = await userLoginService(formModel.value)
  userStore.setToken(res.data.data.token)

  ElMessage({
    message: '登录成功',
    type: 'success'
  })
  router.push('/')
}

//注册逻辑
const register = async () => {
  // 注册成功之前，先进行校验，校验成功 ->请求，验证失败 ->自动提示
  await form.value.validate()
  formModel.value.ID = parseInt(formModel.value.ID)
  await userRegisterService(formModel.value)
  ElMessage({
    message: '注册成功',
    type: 'success'
  })
  isRegister.value = false
}
// 切换的时候，重置表单内容

watch(isRegister, () => {
  formModel.value = {
    ID: '',
    name: '',
    password: '',
    repassword: ''
  }
})
</script>

<template>
  <!--
     1.结构
      el-row表示一行，一行分成24份
      el-col表示列
      (1) :span="12" 代表在一行中，占12份
      (2) :span="6" 代表在一行中，占6份
      (3) :offset="2" 代表在一行中，左侧的margin份数


      el-form 整个表单组件
      el-form-item 表单中的一行
      el-input 表单元素(输入框)
     2.校验
     (1)el-form => :model="ruleForm" 绑定整个form的数据对象  {***, ***, ***}
     (2)el-form => :rules="rules" 绑定整个rules规则对象      {***, ***, ***}
     (3)表单元素 => v-model="rulesForm.***" 给表单元素，绑定form的子属性
     (4)el-form-item => prop配置生效的是哪个校验规则
 -->
  <el-row class="login-page">
    <!--左边一半的背景图-->
    <el-col :span="15" class="bg"></el-col>
    <el-col :span="6" :offset="2" class="form">
      <!--注册相关表单-->
      <el-form
        :model="formModel"
        :rules="rules"
        ref="form"
        size="large"
        autocomplete="off"
        v-if="isRegister"
      >
        <el-form-item>
          <h1>注册</h1>
        </el-form-item>
        <el-form-item prop="ID">
          <el-input
            v-model="formModel.ID"
            :prefix-icon="Postcard"
            placeholder="请输入学工号"
          ></el-input>
        </el-form-item>
        <el-form-item prop="name">
          <el-input
            v-model="formModel.name"
            :prefix-icon="User"
            placeholder="请输入姓名"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formModel.password"
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item prop="repassword">
          <el-input
            v-model="formModel.repassword"
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入再次密码"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button
            @click="register"
            class="button"
            type="primary"
            auto-insert-space
          >
            注册
          </el-button>
        </el-form-item>
        <el-form-item class="flex">
          <el-link type="info" :underline="false" @click="isRegister = false">
            ← 返回
          </el-link>
        </el-form-item>
      </el-form>
      <el-form
        :model="formModel"
        :rules="rules"
        ref="form"
        size="large"
        autocomplete="off"
        v-else
      >
        <el-form-item>
          <h1>登录</h1>
        </el-form-item>
        <el-form-item prop="ID">
          <el-input
            v-model="formModel.ID"
            :prefix-icon="Postcard"
            placeholder="请输入学工号"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formModel.password"
            name="password"
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button
            @click="login"
            class="button"
            type="primary"
            auto-insert-space
            >登录</el-button
          >
        </el-form-item>
        <el-form-item class="flex">
          <el-link type="info" :underline="false" @click="isRegister = true">
            注册 →
          </el-link>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<style lang="scss" scoped>
.login-page {
  height: 100vh;
  background-color: #fff;
  .bg {
    background: url('../../assets/log_bg.jpg') no-repeat center / cover;
    border-radius: 0 20px 20px 0;
  }
  .form {
    display: flex;
    flex-direction: column;
    justify-content: center;
    user-select: none;
    .title {
      margin: 0 auto;
    }
    .button {
      width: 100%;
    }
    .flex {
      width: 100%;
      display: flex;
      justify-content: space-between;
    }
  }
}
</style>
