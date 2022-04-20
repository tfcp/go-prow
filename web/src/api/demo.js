import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/demo/user-list',
    method: 'get',
    params,
    baseURL: "http://127.0.0.1:8008/"
  })
}

export function getDetail(id) {
  return request({
    url: '/demo/user-detail?id='+id,
    method: 'get',
    baseURL: "http://127.0.0.1:8008/"
  })
}

export function Delete(id) {
  return request({
    url: '/demo/user-delete?id='+id,
    method: 'post',
    baseURL: "http://127.0.0.1:8008/"
  })
}

export function enable(id) {
  let params = {
    id: id,
    status: 1
  }
  return request({
    url: '/demo/user-change',
    method: 'post',
    params,
    baseURL: "http://127.0.0.1:8008/"
  })
}

export function disable(id) {
  let params = {
    id: id,
    status: 2
  }
  return request({
    url: '/demo/user-change',
    method: 'post',
    params,
    baseURL: "http://127.0.0.1:8008/"
  })
}

export function save(params) {
  return request({
    url: '/demo/user-save',
    method: 'post',
    params,
    baseURL: "http://127.0.0.1:8008/"
  })
}
