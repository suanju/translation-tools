<template>
    <div class="sticky top-0 z-40 lg:mx-auto lg:max-w-9xl lg:px-8 bg-slate-50 dark:bg-dark-bg-primary">
        <div
            class="flex h-20 items-center gap-x-4  border-gray-200 px-4 shadow-sm sm:gap-x-6 sm:px-6 lg:px-0 lg:shadow-none bg-slate-50 dark:bg-dark-bg-primary">
            <button type="button" class="-m-2.5 p-2.5 text-gray-700 lg:hidden" @click="globalInfo.sidebarShow = true">
                <span class="sr-only">Open sidebar</span>
                <Bars3Icon class="h-6 w-6" aria-hidden="true" />
            </button>
            <!-- 分隔符 -->
            <div class="h-6 w-px bg-gray-200 lg:hidden" aria-hidden="true" />
            <div class="flex flex-1 gap-x-4 self-stretch lg:gap-x-6">
                <div class="relative flex flex-1 items-center">
                    <span class="font-bold text-3xl tracking-widest  text-indigo-600">开始翻译</span>
                </div>
                <div class="flex items-center gap-x-4 lg:gap-x-6">
                    <button type="button" class="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500">
                        <span class="sr-only">View notifications</span>
                        <BellIcon class="h-6 w-6" aria-hidden="true" />
                    </button>
                    <div>
                        <span @click="globalInfo.theme = globalInfo.theme == 'dark' ? 'light' : 'dark'"
                            class="inline-flex items-center rounded-md bg-gray-50 dark:bg-dark-bg-primary  px-2 py-1 text-xs font-medium text-gray-600 ring-1 ring-inset ring-gray-500/10 dark:ring-dark-ring-primary">
                            <DarkIcon v-if="globalInfo.theme === 'dark'" class="w-5 h-5" :fontControlled="false" filled />
                            <LightIcon v-else class="w-5 h-5" :fontControlled="false" filled />
                        </span>
                    </div>
                    <!-- 分隔符 -->
                    <div class="hidden lg:block lg:h-6 lg:w-px lg:bg-gray-200 dark:bg-dark-ring-primary"
                        aria-hidden="true" />
                        <button type="button" class="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500">
                        <span class="sr-only">设置</span>
                        <Cog6ToothIcon class="h-6 w-6" aria-hidden="true" @click="globalInfo.settingShow = true" />
                    </button>
                    <Setting></Setting>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import {
    Bars3Icon,
    BellIcon,
    Cog6ToothIcon
} from "@heroicons/vue/24/outline";

import { useGlobalStore } from "~/store/global"
import LightIcon from "~/assets/icons/light.svg"
import DarkIcon from "~/assets/icons/dark.svg"
import Setting from "./setting.vue"

const globalInfo = useGlobalStore()


onNuxtReady(async () => {
    watch(() => globalInfo.theme, async () => {
        await globalInfo.updateTheme()
    }, { immediate: true })
})
</script>