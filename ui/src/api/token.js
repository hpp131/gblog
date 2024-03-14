import client from './client'

export const LOGIN = (data) => client({
    url: '/gblogs/api/v1/tokens/',
    method: 'post',
    data: data
})
