export default defineNuxtConfig({
  devtools: { enabled: true },
  typescript: {
    strict: true
  },
  css: ['~/assets/css/main.scss'],
  modules: [
    '@pinia/nuxt',
    '@pinia-plugin-persistedstate/nuxt',
    '@element-plus/nuxt',
    '@nuxtjs/tailwindcss',
    "@tailwindcss/forms",
    'nuxt-svgo'
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
  },
  vite: {
    css: {
        scss: {
          additionalData: '@use "@/assets/css/variables.scss" as *;'
        }
      }
    }
  }
)
