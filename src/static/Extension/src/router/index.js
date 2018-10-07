import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Home from '@/components/Home'
import Introduction from '@/components/Introduction'
import Order from '@/components/Order'
import OrderNew from '@/components/OrderNew'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: 'Login'
    },
    {
      path: '/Login',
      name:'Login',
      component: Login
    },
    {
      path: '/Home',
      // name: 'Home',
      component: Home,
      children: [
        // {
        //   path: '/',
        //   component: Introduction
        // },
        {
          path: '/Self',
          name: 'Self',
          component: Introduction
        },
        {
          path: '/Order',
          name: 'Order',
          component: Order
        },
        {
          path: '/OrderNew',
          name: 'OrderNew',
          component: OrderNew
        },
        {
          path: '/OrderManage',
          name: 'OrderManage',
          component: Order
        },
        {
          path: '/Summary',
          name: 'Summary',
          component: Introduction
        },
        {
          path: '/Introduction',
          name: 'Introduction',
          component: Introduction
        }
      ]
    }
  ]
})
