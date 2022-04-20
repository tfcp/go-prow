import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    data,
    baseURL: "http://127.0.0.1:8008/"
  })
}

export function getInfo(token) {
  return request({
    url: '/user/info',
    method: 'get',
    params: { token },
    baseURL: "http://127.0.0.1:8008/"
  })
}
