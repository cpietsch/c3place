import Vue from 'vue'
import VueRouter from 'vue-router'
import Draw from '../components/Draw.vue'
import Kiosk from '../components/Kiosk.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'draw',
    component: Draw
  },
  {
    path: '/kiosk',
    name: 'kiosk',
    component: Kiosk
  }
  // {
  //   path: '/about',
  //   name: 'about',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = new VueRouter({
  routes
})

export default router
