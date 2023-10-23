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
                            <LightIcon v-else class="w-5 h-5" :fontControlled="false"
                                filled />
                        </span>
                    </div>
                    <!-- 分隔符 -->
                    <div class="hidden lg:block lg:h-6 lg:w-px lg:bg-gray-200 dark:bg-dark-ring-primary"
                        aria-hidden="true" />
                    <!-- 用户下拉 -->
                    <Menu as="div" class="relative">
                        <MenuButton class="-m-1.5 flex items-center p-1.5">
                            <span class="sr-only">Open user menu</span>
                            <span class="hidden lg:flex lg:items-center">
                                <span class="ml-4 text-sm font-semibold leading-6 text-gray-900" aria-hidden="true">
                                    <UserIcon  v-show="globalInfo.theme === 'light'" class="w-8 h-8" :fontControlled="false" filled />
                                    <UserDarkIcon v-show="globalInfo.theme === 'dark'" class="w-8 h-8" :fontControlled="false" filled />
                                </span>
                            </span>
                        </MenuButton>
                        <transition enter-active-class="transition ease-out duration-100"
                            enter-from-class="transform opacity-0 scale-95" enter-to-class="transform opacity-100 scale-100"
                            leave-active-class="transition ease-in duration-75"
                            leave-from-class="transform opacity-100 scale-100"
                            leave-to-class="transform opacity-0 scale-95">
                            <MenuItems
                                class="absolute right-0 z-10 mt-2.5 w-48 origin-top-right rounded-md bg-white  dark:bg-dark-bg-moderate-grey py-2 shadow-lg ring-1 ring-gray-900/5 focus:outline-none">
                                <!-- 用户下拉选择 -->
                                <UserMenuInfo />
                            </MenuItems>
                        </transition>
                    </Menu>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import {
    Menu,
    MenuButton,
    MenuItems,
} from "@headlessui/vue";
import {
    Bars3Icon,
    BellIcon,
} from "@heroicons/vue/24/outline";

import UserIcon from "~/assets/icons/user.svg"
import UserDarkIcon from "~/assets/icons/user-dark.svg"
import UserMenuInfo from "./user_menu_info.vue"
import { useGlobalStore } from "~/store/global"
import LightIcon from "~/assets/icons/light.svg"
import DarkIcon from "~/assets/icons/dark.svg"

const globalInfo = useGlobalStore()


onNuxtReady(async () => {
    watch(() => globalInfo.theme, async () => {
        await globalInfo.updateTheme()
    }, { immediate: true })
})
</script>