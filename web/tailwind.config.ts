/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
    "./app.vue",
  ],
  theme: {
    extend: {
      colors: {
        "dark-bg-primary": '#02050E', 
        "dark-bg-grey": '#3D4047', 
        "dark-bg-darkgray": '#10131B',
        "dark-bg-grey-violet": '#040B21',
        "dark-bg-mauve-violet": '#080D1D',  
        "dark-ring-primary": '#4D4F57', 
        "dark-border-primary": '#4D4F57', 
        "dark-text-primary": '#4D4F57', 
      },
    },
  },
  darkMode: "class",
  plugins: [
    require('@tailwindcss/forms')
  ],
}