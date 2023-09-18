export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  typescript: {
    strict: true
  },
  modules: [
    '@pinia/nuxt',
    '@pinia-plugin-persistedstate/nuxt',
    '@element-plus/nuxt',
    '@nuxtjs/tailwindcss',
    "@tailwindcss/forms"
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
  }
})
