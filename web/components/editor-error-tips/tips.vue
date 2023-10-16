<template>
    <div class="tips-box flex items-center justify-center h-full absolute top-1/2 right-0 transform translate-y-[-50%]">
        <div :class="{ 'bg-neutral-700': !editorInfo.editorError.length, 'bg-red-600': editorInfo.editorError.length }"
            class="tips flex items-center justify-center w-8 h-8 rounded-r-lg rounded-full ">
            <TipsIcon class="text-white w-6 h-6" :fontControlled="false"
                @click="editorInfo.editorTipsShow = !editorInfo.editorTipsShow" />
        </div>
        <div class="tips-error-box max-w-xl h-full bg-white  shadow-md border-l-2"
            :style="{ 'width': editorInfo.editorTipsShow ? '20em' : '0em' }">
            <div class="flex items-center justify-center">
                <h4 class="text-3xl font-normal p-8 border-b-2 w-4/5 text-center whitespace-pre">语法错误</h4>
            </div>
            <!-- 无错误 -->
            <div class="mt-16 mx-10" v-show="!editorInfo.editorError.length">
                <div class="flex items-center justify-center">
                    <TipsCorrectIcon class="w-10 h-10" :fontControlled="false" filled />
                </div>
                <div class="mt-8 w-60"> 没有检测到任何语法错误。你的文件看起来完全没有问题!</div>
            </div>
            <!-- 错误列表 -->
            <div class="mt-10 mx-10 " v-show="editorInfo.editorError.length">
                <div v-for="(item, index) in editorInfo.editorError" :key="index"
                    class="flex pl-4 justify-center flex-col border-l-4 border-l-red-600 w-60 h-24 min-h-full bg-gray-100 rounded-r-lg mb-8">
                    <div class="text-indigo-400">Line : {{ item.startLineNumber }} Column : {{ item.startColumn }}</div>
                    <div class=""> {{ item.message }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import TipsIcon from "~/assets/icons/tips.svg"
import TipsCorrectIcon from "~/assets/icons/tips-correct.svg"
import { useEditorStore } from "~/store/editor"


const editorInfo = useEditorStore()

watch(() => editorInfo.editorError.length, (val) => {
    editorInfo.editorTipsShow = !!val
})

</script>

<style lang="scss" scoped>
.tips-box {
    z-index: 999;
    transition: all 0.1s ease;

    .tips-error-box {
        overflow: hidden;
        transition: width 0.3s ease;
    }
}</style>