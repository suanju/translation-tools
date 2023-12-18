<template>
    <TransitionRoot as="template" :show="globalStore.settingShow">
        <Teleport to="body">
            <Dialog as="div" class="relative z-50" @close="globalStore.settingShow = false">
                <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100"
                    leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
                    <div
                        class="fixed inset-0 bg-gray-500 dark:bg-black dark:bg-opacity-50  bg-opacity-30 transition-opacity" />
                </TransitionChild>
                <div class="fixed inset-0 z-10 overflow-y-auto">
                    <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                        <TransitionChild as="template" enter="ease-out duration-300"
                            enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                            enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200"
                            leave-from="opacity-100 translate-y-0 sm:scale-100"
                            leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
                            <DialogPanel
                                class="relative transform overflow-hidden rounded-lg bg-white h-[26rem] w-full dark:bg-dark-bg-darkgray text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg ">
                                <!-- 弹窗内容 -->
                                <div class="fixed  w-full sm:p-6">
                                    <div class="flex items-center justify-items-center">
                                        <span class="flex-1 dark:text-slate-200">API</span>
                                        <XMarkIcon @click="globalStore.settingShow = false" class="h-5 w-5 text-gray-400"
                                            aria-hidden="true" />
                                    </div>
                                    <div class="mt-6 h-[20rem] relative">
                                        <div>
                                            <label for="type"
                                                class="block text-sm font-medium leading-6 dark:text-slate-200 text-gray-900">翻译接口</label>
                                            <div class="mt-3">
                                                <Listbox v-model="translationStore.selectPlatform">
                                                    <div class="relative mt-1">
                                                        <ListboxButton
                                                            class="relative w-full dark:bg-dark-bg-darkgray  dark:ring-dark-border-primary dark:text-slate-200 ring-1 cursor-default rounded-lg bg-white py-2 pl-3 pr-10 text-left shadow-sm focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white/75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm">
                                                            <span class="block truncate">{{
                                                                translationStore.selectPlatform.name }}</span>
                                                            <span
                                                                class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                                                                <ChevronUpDownIcon class="h-5 w-5 text-gray-400"
                                                                    aria-hidden="true" />
                                                            </span>
                                                        </ListboxButton>
                                                        <transition leave-active-class="transition duration-100 ease-in"
                                                            leave-from-class="opacity-100" leave-to-class="opacity-0">
                                                            <ListboxOptions
                                                                class="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black/5 focus:outline-none sm:text-sm">
                                                                <ListboxOption v-slot="{ active, selected }"
                                                                    v-for="person in translationStore.platformList"
                                                                    :key="person.name" :value="person" as="template">
                                                                    <li :class="[
                                                                        active ? 'bg-indigo-100 text-indigo-600' : 'text-gray-900',
                                                                        'relative cursor-default select-none py-2 pl-10 pr-4',
                                                                    ]">
                                                                        <span :class="[
                                                                            selected ? 'font-medium' : 'font-normal',
                                                                            'block truncate',
                                                                        ]">{{ person.name }}</span>
                                                                        <span
                                                                            v-if="person.code == translationStore.selectPlatform.code"
                                                                            class="absolute inset-y-0 left-0 flex items-center pl-3 text-indigo-600">
                                                                            <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                                                        </span>
                                                                    </li>
                                                                </ListboxOption>
                                                            </ListboxOptions>
                                                        </transition>
                                                    </div>
                                                </Listbox>
                                            </div>
                                        </div>
                                        <div class="mt-3" v-show="translationStore.selectPlatform.code == 'baidu'">
                                            <label for="appid"
                                                class="block text-sm font-medium leading-6 text-gray-900 dark:text-slate-200">APPID</label>
                                            <div class="mt-2">
                                                <input type="text" v-model="translationStore.APPID" name="appid" id="appid"
                                                    class="block w-full rounded-md border-0 py-1.5  dark:bg-dark-bg-darkgray  dark:ring-dark-border-primary dark:text-slate-200 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    placeholder="请填写您的APPID" />
                                            </div>
                                        </div>
                                        <div class="mt-3">
                                            <label for="key"
                                                class="block text-sm font-medium leading-6 text-gray-900 dark:text-slate-200">KEY</label>
                                            <div class="mt-2">
                                                <input type="text" v-model="translationStore.KEY" name="key" id="key"
                                                    class="block w-full rounded-md border-0 py-1.5 dark:bg-dark-bg-darkgray  dark:ring-dark-border-primary dark:text-slate-200 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    placeholder="请填写您的KEY" />
                                            </div>
                                        </div>
                                        <button @click="save" type="button"
                                            class="rounded-lg w-full  absolute bottom-4 bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">保存设置</button>
                                    </div>
                                </div>
                            </DialogPanel>
                        </TransitionChild>
                    </div>
                </div>
            </Dialog>
        </Teleport>
    </TransitionRoot>
</template>
  
<script lang="ts" setup>
import {
    Listbox,
    ListboxButton,
    ListboxOptions,
    ListboxOption,
} from '@headlessui/vue'
import { CheckIcon, ChevronUpDownIcon, XMarkIcon } from '@heroicons/vue/20/solid'
import { Dialog, DialogPanel, TransitionChild, TransitionRoot } from '@headlessui/vue'
import { useGlobalStore } from '~/store/global'
import { useTranslationStore } from '~/store/translation'
import Swal from 'sweetalert2'

const globalStore = useGlobalStore()
const translationStore = useTranslationStore()

const save = () => {
    globalStore.settingShow = false;
    Swal.fire({
        icon: "success",
        title: "保存成功"
    });
}

watch(() => translationStore.selectPlatform?.code, () => {
    translationStore.getLangList()
})

</script>
