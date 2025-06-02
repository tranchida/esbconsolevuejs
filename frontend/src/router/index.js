import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/in/dashboard'
  },
  {
    path: '/:env/dashboard',
    name: 'Dashboard',
    component: () => import('../views/Dashboard.vue')
  },
  {
    path: '/:env/jms',
    name: 'JMS',
    component: () => import('../views/JMS.vue')
  },
  {
    path: '/:env/audit',
    name: 'Audit',
    component: () => import('../views/Audit.vue')
  },
  {
    path: '/:env/admin',
    name: 'Admin',
    component: () => import('../views/Admin.vue')
  },
  
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 