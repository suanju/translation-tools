<template>
  <div>
    <TransitionRoot as="template" :show="globalInfo.sidebarShow">
      <Dialog as="div" class="relative z-50 lg:hidden" @close="globalInfo.sidebarShow = false">
        <TransitionChild as="template" enter="transition-opacity ease-linear duration-300" enter-from="opacity-0"
          enter-to="opacity-100" leave="transition-opacity ease-linear duration-300" leave-from="opacity-100"
          leave-to="opacity-0">
          <div class="fixed inset-0 bg-gray-900/80" />
        </TransitionChild>

        <div class="fixed inset-0 flex">
          <TransitionChild as="template" enter="transition ease-in-out duration-300 transform"
            enter-from="-translate-x-full" enter-to="translate-x-0" leave="transition ease-in-out duration-300 transform"
            leave-from="translate-x-0" leave-to="-translate-x-full">
            <DialogPanel class="relative mr-16 flex w-full max-w-xs flex-1">
              <TransitionChild as="template" enter="ease-in-out duration-300" enter-from="opacity-0"
                enter-to="opacity-100" leave="ease-in-out duration-300" leave-from="opacity-100" leave-to="opacity-0">
                <div class="absolute left-full top-0 flex w-16 justify-center pt-5">
                  <button type="button" class="-m-2.5 p-2.5" @click="globalInfo.sidebarShow = false">
                    <span class="sr-only">Close sidebar</span>
                    <XMarkIcon class="h-6 w-6 text-white" aria-hidden="true" />
                  </button>
                </div>
              </TransitionChild>
              <div class="flex grow flex-col gap-y-5 overflow-y-auto bg-white dark:bg-dark-bg-primary px-6 pb-4">
                <div class="flex h-16 shrink-0 items-center">
                  <img class="h-8 w-auto" :src="LogoIcon" alt="logo" />
                  <span class="ml-2  text-lg font-medium font-serif dark:text-white"> Mfn Transition</span>
                </div>
                <nav class="flex flex-1 flex-col">
                  <ul role="list" class="flex flex-1 flex-col gap-y-7">
                    <li>
                      <ul role="list" class="-mx-2 space-y-1">
                        <li v-for="item in navigation" :key="item.name">
                          <span  :class="[
                            item.current
                              ? 'bg-gray-50 dark:bg-dark-bg-grey-violet text-indigo-600'
                              : 'text-gray-700 dark:text-slate-200 hover:text-indigo-600 hover:bg-gray-50 dark:hover:bg-dark-bg-mauve-violet',
                            'group flex gap-x-3 rounded-md p-2 text-base leading-6 font-semibold',
                          ]">
                            <component :is="item.icon" :class="[
                              item.current
                                ? 'text-indigo-600'
                                : 'text-gray-400  group-hover:text-indigo-600',
                              'h-6 w-6 shrink-0',
                            ]" aria-hidden="true" />
                            {{ item.name }}
                          </span>
                        </li>
                      </ul>
                    </li>
                    <li>
                      <div class="text-xs font-semibold leading-6 text-gray-400">
                        Your teams
                      </div>
                    </li>
                  </ul>
                </nav>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </Dialog>
    </TransitionRoot>

    <!-- 侧边栏-->
    <div class="hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
      <!-- 侧边栏组件 -->
      <div
        class="flex grow flex-col gap-y-5 overflow-y-auto border-r border-gray-200 dark:border-dark-border-primary bg-white dark:bg-dark-bg-primary px-6 pb-4">
        <div class="flex h-16 shrink-0 items-center">
          <img class="h-8 w-auto" :src="LogoIcon" alt="logo" />
          <span class="ml-2  text-lg font-medium font-serif dark:text-white"> Mfn Transition</span>
        </div>
        <nav class="flex flex-1 flex-col">
          <ul role="list" class="flex flex-1 flex-col gap-y-7">
            <li>
              <ul role="list" class="-mx-2 space-y-1">
                <li v-for="item in navigation" :key="item.name" @click="jumpMenu(item)">
                  <span :class="[
                    item.current
                      ? 'bg-gray-50 dark:bg-dark-bg-grey-violet text-indigo-600'
                      : 'text-gray-700 dark:text-slate-200 hover:text-indigo-600 hover:bg-gray-50 dark:hover:bg-dark-bg-mauve-violet',
                    'group flex gap-x-3 rounded-md p-3 text-base leading-6 font-semibold h-16 items-center',
                  ]">
                    <component :is="item.icon" :class="[
                      item.current
                        ? 'text-indigo-600'
                        : 'text-gray-400 group-hover:text-indigo-600',
                      'h-6 w-6 shrink-0',
                    ]" aria-hidden="true" />
                    {{ item.name }}
                  </span>
                </li>
              </ul>
            </li>
          </ul>
        </nav>
      </div>
    </div>

    <div class="lg:pl-72">
      <!-- 顶部导航 -->
      <HeadNavigation></HeadNavigation>
      <main class="main pt-8 bg-slate-50 dark:bg-dark-bg-primary">
        <div class="mx-auto max-w-9xl px-4 sm:px-6 lg:px-8">
          <!-- 中心内容 -->
          <NuxtPage />
        </div>
      </main>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  Dialog,
  DialogPanel,
  TransitionChild,
  TransitionRoot,
} from "@headlessui/vue";
import {
  CalendarIcon,
  Cog6ToothIcon,
  FolderIcon,
  HomeIcon,
  UsersIcon,
  XMarkIcon,
} from "@heroicons/vue/24/outline";
import LogoIcon from "~/assets/logo/logo.png"
import { useGlobalStore } from "~/store/global"
import HeadNavigation from "~/components/layout-item/head_navigation.vue"

const router = useRouter();

const navigation = [
  { name: "开始翻译", link: "translation", icon: HomeIcon, current: true },
  { name: "历史记录", link: "#", icon: UsersIcon, current: false },
  { name: "立即订阅", link: "#", icon: FolderIcon, current: false },
  { name: "接口设置", link: "setting", icon: Cog6ToothIcon, current: false },
  { name: "帮助中心", link: "#", icon: CalendarIcon, current: false },
];

const globalInfo = useGlobalStore()

const jumpMenu = (item) => {
  console.log(item)
  router.push({ name: item.link });
}

</script>

<style lang="scss" scoped>
.main {
  min-height: calc(100vh - 5rem);
}
</style>
