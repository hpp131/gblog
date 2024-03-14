import { Message } from '@arco-design/web-vue';
import axios from 'axios';

// 创建httpInstance实例
const htptInstance = axios.create({
    timeout: 5000,
})

// axios分别提供了request和response的拦截器，这里我们只用到response拦截器
htptInstance.interceptors.response.use(
    // 请求成功时
    (success) => {
        return success.data
    },
    // 请求失败时
    (error) => {
        var msg = error.Message
        if (error.response.data && error.response.data.Message){
            msg = err.response.data.Message
        }
        
        // 使用arco design提供的"全局提示 Message"组件
        Message.error({
            content:msg,
            duration:2000
        })
        return Promise.reject(err)
    },
)

export default htptInstance