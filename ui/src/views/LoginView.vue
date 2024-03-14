<script setup>
import { LOGIN } from '@/api/token';
import { state } from '@/stores/app';
import { ref } from 'vue';



const form = ref({
  username: '',
  password: '',
  isRead: true,
});

const handleSubmit = (data) => {
  // 此处更换为api调用
  console.log(data);
  LOGIN(form.value).then(function(response){
    // 关于localstorage
    state.value.token = response
  })
};

const formRules = {
  username: {
    required: true,
    message:'Lack of username'
  },
  password:[
    {
      required: true,
      message:'Lack of password'
    },
    {
      minLength:6,
      message:'Please input great than 6 characters'
    }
  ]
}

</script>

<template>
  <div class="content">
    <a-form :model="form" :rules="formRules" class="login-form"  @submit="handleSubmit">
      <div class="title">GBLOG SYSTEM</div>
      <a-form-item hide-label field="username"  label="">
        <a-input v-model="form.username" placeholder="please enter your username..." >
          <!-- prefix是arco design提供的具名插槽,该插槽并没有在我们本地代码中定义 -->
          <template #prefix>
            <icon-user />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item hide-label field="password" label="">
        <a-input v-model="form.password" placeholder="please enter your password..." >
          <template #prefix>
            <icon-lock />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item hide-label field="isRead">
        <a-checkbox style="display: flex; justify-content: center" v-model="form.isRead">记住我</a-checkbox>
      </a-form-item>
      <a-form-item hide-label>
        <a-button type="primary" html-type="submit" style="width: 100%;">Submit</a-button>
      </a-form-item>
    </a-form>
  </div>

</template>

<style lang="css" scoped>
.title{
  margin-bottom: 20px;
  display: flex;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
}

.content{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
}

.login-form{
  display: flex;
  justify-content: center;
  width: 500px;
  height: 100%;
  margin-top: 100px;
}

</style>
