<script setup>
import {
  Management,
  Promotion,
  UserFilled,
  User,
  Crop,
  EditPen,
  SwitchButton,
  CaretBottom,
  Notification,
  Setting,
  Edit,
  List,
  Share
} from '@element-plus/icons-vue'
import avatar from '../../assets/default.png'
import { useUserStore } from '../../stores'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

const userStore = useUserStore()
const router = useRouter()
onMounted(() => {
  userStore.getUser()
})

// 下拉列表逻辑
const handleCommand = async (key) => {
  if (key === 'logout') {
    // 退出操作
    await ElMessageBox.confirm('你确定要进行退出吗', '温馨提示', {
      type: 'warning',
      confirmButtonText: '确认',
      cancelButtonText: '取消'
    })
    // 1.清楚本地的数据(token + 用户信息user)
    userStore.removeToken()
    userStore.setUser({})
    // 2.跳转至登录页
    router.push('/login')
  } else {
    // 跳转操作
    router.push(`/user/${key}`)
  }
}
</script>

<template>
  <el-container class="layout-container">
    <el-aside width="200px">
      <!-- 新增标题 -->
      <div>
        <h1 class="sidebar-title">图书管理系统</h1>
      </div>
      <!--
      el-menu 菜单组件
      :default-active="$route.path"       配置默认高亮的菜单项
      router   router选项开启， el-menu-item的index就是点击跳转的路径

      el-menu-item 菜单项
          index="/book/channel" 配置的是访问的跳转路径,配合default-active的值实现高亮
      -->
      <el-menu
        active-text-color="#ffd04b"
        background-color="#232326"
        :default-active="$route.path"
        text-color="#fff"
        router
      >
        <el-menu-item index="/list">
          <el-icon><Share /></el-icon>
          <span>图书列表</span>
        </el-menu-item>
        <el-menu-item index="/book/have">
          <el-icon><List /></el-icon>
          <span>我的图书</span>
        </el-menu-item>
        <el-menu-item index="/record">
          <el-icon><Notification /></el-icon>
          <span>借阅记录</span>
        </el-menu-item>
        <el-sub-menu index="/user">
          <!--多级菜单的标题  具名插槽，title-->
          <template #title>
            <el-icon><UserFilled /></el-icon>
            <span>个人中心</span>
          </template>
          <!--展开的内容     默认插槽-->
          <el-menu-item index="/user/profile">
            <el-icon><User /></el-icon>
            <span>基本资料</span>
          </el-menu-item>
          <el-menu-item index="/user/avatar">
            <el-icon><Crop /></el-icon>
            <span>更换头像</span>
          </el-menu-item>
          <el-menu-item index="/user/password">
            <el-icon><EditPen /></el-icon>
            <span>重置密码</span>
          </el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="/manage" v-if="userStore.user.permission">
          <!--多级菜单的标题  具名插槽，title-->
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>管理</span>
          </template>
          <el-menu-item index="/book/channel">
            <el-icon><Management /></el-icon>
            <span>分类管理</span>
          </el-menu-item>
          <el-menu-item index="/book/manage">
            <el-icon><Promotion /></el-icon>
            <span>图书管理</span>
          </el-menu-item>
          <el-menu-item index="/user/manage">
            <el-icon><Edit /></el-icon>
            <span>用户管理</span>
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header>
        <div>
          欢迎您：<strong>{{
            userStore.user.nickname || userStore.user.name
          }}</strong>
        </div>
        <!--展示给用户看的，默认看到的-->
        <el-dropdown placement="bottom-end" @command="handleCommand">
          <span class="el-dropdown__box">
            <el-avatar :src="userStore.user.user_pic || avatar" />
            <el-icon><CaretBottom /></el-icon>
          </span>

          <!--折叠的下拉部分 command对应路由名字-->
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="profile" :icon="User"
                >基本资料</el-dropdown-item
              >
              <el-dropdown-item command="avatar" :icon="Crop"
                >更换头像</el-dropdown-item
              >
              <el-dropdown-item command="password" :icon="EditPen"
                >重置密码</el-dropdown-item
              >
              <el-dropdown-item command="logout" :icon="SwitchButton"
                >退出登录</el-dropdown-item
              >
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
      <el-footer>图书管理系统</el-footer>
    </el-container>
  </el-container>
</template>

<style lang="scss" scoped>
.layout-container {
  height: 100vh;
  .el-aside {
    background-color: #232323;
    &__logo {
      height: 120px;
    }
    .el-menu {
      border-right: none;
    }
  }
  .el-header {
    background-color: #fff;
    display: flex;
    align-items: center;
    justify-content: space-between;
    .el-dropdown__box {
      display: flex;
      align-items: center;
      .el-icon {
        color: #999;
        margin-left: 10px;
      }

      &:active,
      &:focus {
        outline: none;
      }
    }
  }
  .el-footer {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    color: #666;
  }
}
// 添加对标题的自定义样式
.sidebar-title {
  margin: 20px 10px; /* 添加标题与其他元素之间的垂直间距 */
  padding: 10px; /* 添加内边距，使标题更具吸引力 */
  font-size: 1.5rem; /* 设置标题字体大小 */
  color: #fff; /* 设置标题颜色为白色 */
  background-color: #232323; /* 设置标题背景颜色 */
}
</style>
