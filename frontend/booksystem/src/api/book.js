// 获取图书分类
import request from '../utils/request.js'

// 分类:获取图书分类
export const bookGetChannelsService = () => request.get('/api/cate/list')

// 分类:添加图书分类
export const bookAddChannelService = (data) =>
  request.post('/api/cate/add', data)

// 分类:编辑图书分类
export const bookEditChannelService = (data) =>
  request.put('/api/cate/info', data)

// 分类:删除图书分类
export const bookDeleteChannelService = (id) =>
  request.delete('/api/cate/del', {
    params: { id }
  })

// 图书: 获取图书列表
export const bookGetListService = (params) =>
  request.get('/api/letter/list', {
    params
  })

// 图书: 添加图书
// 注意：data需要是一个formData格式的对象
export const bookAddService = (data) => request.post('/api/letter/add', data)

export const bookGetDetailService = (id) =>
  request.get('/api/letter/info', {
    params: { id }
  })

// 图书: 编辑图书
export const bookEditService = (data) => request.put('/api/letter/info', data)
export const bookDeleteService = (id) =>
  request.delete('/api/letter/info', {
    params: { id }
  })


// 借书
export const bookBorrowService =(data) => request.post('/api/record/borrow', data)
export const bookLendService =(data) => request.post('/api/record/lend', data)

