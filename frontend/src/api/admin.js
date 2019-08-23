import request from '@/utils/request'

export function setMonth(data) {
  return request({
    url: '/admin/month',
    method: 'post',
    data
  })
}
