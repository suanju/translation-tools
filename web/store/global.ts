import { defineStore } from "pinia"
import type { SidebarShow, Theme } from '~/types/store/global'

export const useGlobalStore = defineStore("global", () => {
    let sidebarShow = ref<SidebarShow>(false)
    let theme = ref<Theme>("light")
    let settingShow = ref(false)

    const updateTheme = async () => {
        const monaco = (await import("monaco-editor"));
        console.log("准备设置主题")
        if (theme.value === 'dark') {
            monaco.editor.setTheme("brilliance-black");
            document.querySelector("html").classList.add("dark")
        } else if (theme.value === 'light') {
            monaco.editor.setTheme("vs");
            document.querySelector("html").classList.remove("dark")
        }
    }

    return {
        sidebarShow,
        theme,
        settingShow,
        updateTheme
    }
}, {
    persist: {
        paths: ['sidebarShow','theme']
    },
})