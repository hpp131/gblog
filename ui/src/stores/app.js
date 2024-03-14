// 需要使用到vueuse组件库中的"useStorage"功能
// 官网地址https://vueuse.org/

import { useStorage } from "@vueuse/core";

export const state = useStorage(
    'gblog-store',
    {token:null},
    localStorage,
    // 读取localstorage 里面的gblog-store的值作为参数 一起合并
    {mergeDefaults:true}
)