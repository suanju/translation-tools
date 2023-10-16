<template>
  <div class="editor-box rounded">
    <!-- 工具栏 -->
    <div class="h-16 flex items-center border rounded-rl-2xl px-6 bg-white">
      <div class="flex items-center flex-auto">
        <div class="mt-2 flex flex-wrap w-full" ref="containerRef">
          <button type="button" v-for="item in editorInfo.selectedList" :key="item.code"
            @click="editorInfo.uncheckById(item.id)"
            class="flex justify-center  rounded bg-inherit px-2 py-1 text-xs font-semibold text-slate-500 shadow-sm hover:bg-indigo-100 mr-2 mt-1 w-14">
            <span class="flex-1">{{ item.code }}</span>
            <XMarkIcon class="ml-3 w-3 h-4 text-xs" />
          </button>
        </div>
      </div>
      <button type="button" @click="editorInfo.resultsSelectLangDialogShow = true"
        class=" w-32 rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
        选择语言
      </button>
    </div>
    <!-- 选择语言弹出层 -->
    <SelectLang></SelectLang>
    <!-- 代码编辑器 -->
    <div ref="codeEditorRef" class="editor-box"></div>
  </div>
</template>

<script lang="ts" setup>
import * as monaco from "monaco-editor";
import 'monaco-editor/esm/vs/basic-languages/javascript/javascript.contribution';
import SelectLang from "./select_lang.vue";
import { XMarkIcon } from '@heroicons/vue/24/outline'
import { useEditorStore } from "~/store/editor"
import { useGlobalStore } from "~/store/global";


const codeEditorRef = ref(null);
const containerRef = ref(null)
const editorInfo = useEditorStore()
const globalInfo = useGlobalStore()

watch(() => editorInfo.selectedList, (_, nlv) => {
  console.log(containerRef.value.offsetWidth)
}, {
  deep: true
})


onMounted(() => {
  //加载主题
  const theme = globalInfo.theme == "light" ? "vs" : "brilliance-black"
  codeEditorRef.value.style.height = "calc(100vh - 64px - 5rem - 4.2rem)";
  monaco.editor.create(codeEditorRef.value as HTMLElement, {
    theme, // 主题
    value: "", // 默认显示的值
    language: "json",
    folding: true, // 是否折叠
    foldingHighlight: true, // 折叠等高线
    foldingStrategy: "indentation", // 折叠方式  auto | indentation
    showFoldingControls: "always", // 是否一直显示折叠 always | mouseover
    disableLayerHinting: true, // 等宽优化
    emptySelectionClipboard: false, // 空选择剪切板
    selectionClipboard: false, // 选择剪切板
    automaticLayout: true, // 自动布局
    codeLens: false, // 代码镜头
    scrollBeyondLastLine: false, // 滚动完最后一行后再滚动一屏幕
    colorDecorators: true, // 颜色装饰器
    accessibilitySupport: "off", // 辅助功能支持  "auto" | "off" | "on"
    lineNumbers: "on", // 行号 取值： "on" | "off" | "relative" | "interval" | function
    lineNumbersMinChars: 5, // 行号最小字符   number
    readOnly: false, //是否只读  取值 true | false
    cursorSmoothCaretAnimation: "on", // 是否启用光标平滑插入动画  当你在快速输入文字的时候 光标是直接平滑的移动还是直接"闪现"到当前文字所处位置
  });


});


</script>

<style></style>
