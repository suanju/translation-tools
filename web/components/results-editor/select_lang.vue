<template>
    <TransitionRoot as="template" :show="translationStore.resultsSelectLangDialogShow">
        <Teleport to="body">
            <Dialog as="div" class="relative z-50" @close="translationStore.resultsSelectLangDialogShow = false">
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
                                class="relative transform overflow-hidden rounded-lg bg-white  dark:bg-dark-bg-darkgray px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-md sm:p-6">
                                <!-- 顶部选择 -->
                                <div class="flex items-center">
                                    <span class="flex-1 text-indigo-600 font-semibold">添加输出语言</span>
                                    <XMarkIcon class="w-4 h-4 text-xs"
                                        @click="translationStore.resultsSelectLangDialogShow = false" />
                                </div>
                                <!-- 选择语言 -->
                                <div class="mt-2">
                                    <input v-model="findText"
                                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 dark:text-slate-200 shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-dark-border-primary placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 dark:bg-dark-bg-darkgray" />
                                </div>
                                <!-- 选择区域 -->
                                <div class="mt-4 border rounded-md border-slate-300 dark:border-dark-border-primary">
                                    <div class="h-80  overflow-y-auto custom-scrollbar scrollbar-thumb-black ">
                                        <legend class="sr-only">Notifications</legend>
                                        <div class="space-y-5 py-4 px-4  hover:bg-gray-100 dark:hover:bg-dark-bg-grey"
                                            v-for="item in findList" :key="item.code">
                                            <div class="relative flex items-start">
                                                <div class="flex h-6 items-center">
                                                    <input aria-describedby="comments-description" name="comments"
                                                        v-model="item.check" type="checkbox"
                                                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600" />
                                                </div>
                                                <div class="ml-3 text-sm leading-6">
                                                    <label for="comments"
                                                        class="font-medium text-gray-900 dark:text-slate-200">{{ item.lang
                                                        }}</label>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                                <!-- 已选择 -->
                                <div class="mt-4 flex items-center justify-center">
                                    <span class="text-xs font-extralight flex-1 dark:text-slate-200">选中 : </span>
                                    <span>
                                        <button type="button" @click="translationStore.uncheckAll"
                                            class="rounded bg-indigo-50 dark:bg-slate-800 px-2 py-1 text-xs font-semibold text-indigo-600 shadow-sm hover:bg-indigo-100 dark:hover:bg-dark-bg-grey">
                                            清除所有
                                        </button></span>
                                </div>

                                <div class="mt-2 flex flex-wrap">
                                    <button type="button" v-for="item in translationStore.selectedList" :key="item.code"
                                        @click="translationStore.uncheckById(item.code)"
                                        class="flex justify-center  rounded bg-inherit px-2 py-1 text-xs font-semibold text-slate-500 shadow-sm hover:bg-indigo-100  dark:hover:bg-dark-bg-grey  mr-2 mt-1 w-14">
                                        <span class="flex-1 dark:text-slate-200">{{ item.code }}</span>
                                        <XMarkIcon class="ml-3 w-3 h-4 text-xs" />
                                    </button>
                                </div>

                                <div class="w-full mt-2">
                                    <button v-wave type="button" @click="translationStore.resultsSelectLangDialogShow = false"
                                        class="w-full justify-center inline-flex items-center gap-x-1.5 rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                                        <CheckCircleIcon class="-ml-0.5 h-5 w-5" aria-hidden="true" />
                                        确认语言
                                    </button>
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
import { Dialog, DialogPanel, TransitionChild, TransitionRoot } from '@headlessui/vue'
import { XMarkIcon, CheckCircleIcon } from '@heroicons/vue/24/outline'
import { useTranslationStore } from "~/store/translation"

const translationStore = useTranslationStore()
const findText = ref("")

const findList = computed(() => {
    return translationStore.originalTranslationLang.filter((item) => {
        return item.lang.includes(findText.value)
    })
})

watch(() => translationStore.originalTranslationLang, (_, nlv) => {
    translationStore.selectedList = nlv.filter((item) => {
        return item.check
    })
    return translationStore.selectedList
}, {
    deep: true
})


</script>

<style lang="scss" scoped>

::-webkit-scrollbar {
    width: 4px;
    border-radius: 8px;
}

::-webkit-scrollbar-thumb {
    background-color: #757575;
    border-radius: 2rem;
}

</style>