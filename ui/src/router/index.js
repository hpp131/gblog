import LoginView from '../views/LoginView.vue'
import BackendLayout from '../views/backend/BackendLayout.vue'
import FrontendLayout from '../views/frontend/FrontendLayout.vue'

import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/backend',
      name: 'backend',
      component: BackendLayout
    },
    {
      path: '/frontend',
      name: 'frontend',
      component: FrontendLayout
    }
  ]
})

export default router
