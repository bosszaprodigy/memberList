import { createRouter, createWebHistory } from 'vue-router'
// import { useAuthStore } from '../stores/auth'

const routes = [
  {
    path: '/',
    redirect: '/member'
  },
  {
    path: '/member',
    name: 'Member',
    component: () => import('../views/MemberView.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
