export default defineNuxtConfig({
  devtools: { enabled: true },
<<<<<<< HEAD
  typescript: {
    strict: true
  },
  css: ['~/assets/css/main.scss'],
=======
  css: ['~/assets/css/main.css'],
  typescript: {
    strict: true
  },
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
  modules: [
    '@pinia/nuxt',
    '@pinia-plugin-persistedstate/nuxt',
    '@element-plus/nuxt',
    '@nuxtjs/tailwindcss',
<<<<<<< HEAD
    "@tailwindcss/forms",
    'nuxt-svgo'
=======
    "@tailwindcss/forms"
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
  ],
  runtimeConfig: {
    public: {
      baseApi: process.env.NUXT_PUBLIC_API_BASE
    }
  },
  postcss: {
    plugins: {
      autoprefixer: {},
    },
  },
  piniaPersistedstate: {
    cookieOptions: {
      sameSite: 'strict',
    },
    storage: 'localStorage'
<<<<<<< HEAD
  },
  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: '@use "@/assets/css/variables.scss" as *;'
        }
      }
    }
=======
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
  }
})
