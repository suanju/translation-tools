import type { RouterConfig } from '@nuxt/schema'

export default <RouterConfig>{
    routes: (_routes) => [
        {
            name: 'login',
            path: '/login',
            component: () => import('~/pages/login/login.vue')
        },
        {
            name: 'register',
            path: '/register',
            component: () => import('~/pages/login/register.vue')
        }
    ],
}
