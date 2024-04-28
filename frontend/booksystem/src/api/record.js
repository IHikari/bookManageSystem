import request from '../utils/request.js'


export const recordGetService = () => request.get('/api/record/list')