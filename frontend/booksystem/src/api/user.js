import request from '../utils/request.js'

// 注册接口
export const userRegisterService = ({ ID, name, password, repassword }) => {
  return request.post('/api/reg', { ID, name, password, repassword })
}

// 登录接口(为以上写法的简写)
export const userLoginService = ({ ID, password }) =>
  request.post('/api/login', {
    ID,
    password
  })

// 获取用户的基本信息
export const userGetInfoService = () => request.get('/api/my/userinfo')

// 修改用户基本信息
export const userEditInfoService = ({ ID, nickname, email }) =>
  request.put('/api/my/userinfo', { ID, nickname, email })

// 修改用户头像
export const userEditAvatarService = (data) =>
  request.patch(
    '/api/my/avatar', data)

//更新用户密码
export const userUpdatePassService = ({ old_pwd, new_pwd, re_pwd }) =>
  request.patch('/api/my/updatepwd', { old_pwd, new_pwd, re_pwd })


//获取用户列表
export const userGetListService = (params) =>
  request.get('/api/user/list', {
    params
  })

// 删除用户列表
export const userDeleteService = (id) =>
  request.delete('/api/user/info', {
    params: { id }
  })

// 管理员添加用户
export const userAddService = (data) =>
  request.post('/api/user/add', data)

// 管理员编辑用户
export const userEditService = (data) =>
  request.put('/api/user/info', data)

// 获取我的图书列表
export const userGetBookListService = () =>
  request.get('/api/my/booklist' )

export const userBookGetDetailService = (id) =>
  request.get('/api/my/book', {
    params: { id }
  })