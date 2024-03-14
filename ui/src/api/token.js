import client from './client'

export const LOGIN = (data) => client({
    url: '/gblogs/api/v1',
    method: 'post',
    data: data
})
