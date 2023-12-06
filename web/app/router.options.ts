import type { RouterConfig } from '@nuxt/schema'

export default <RouterConfig>{
    routes: (_routes) => [
        { 
            name: 'hone',
            path: '',
            component: () => import('~/pages/layout.vue'),
            children:[
                {
                    name:"translation",
                    path :"",
                    component: () => import('~/pages/index/index.vue'),
                },
                {
                    name: 'setting',
                    path: '/setting',
                    component: () => import('~/pages/setting/setting.vue')
                }
            ]
        }
    ],
}
