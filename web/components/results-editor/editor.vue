<template>
  <div class="editor-box rounded">
    <!-- 工具栏 -->
    <div
      class="h-16 flex items-center border-b border-l rounded-tr-2xl px-6 bg-white dark:bg-dark-bg-primary dark:border-dark-border-primary ">
      <div class="flex items-center flex-auto">
        <div class="mt-2 flex flex-wrap w-full" ref="containerRef">
          <button type="button" v-for="item in translationStore.selectedList" :key="item.code"
            @click="translationStore.uncheckById(item.code)"
            class="flex justify-center  rounded bg-inherit px-2 py-1 text-xs font-semibold text-slate-500 shadow-sm hover:bg-indigo-100 mr-2 mt-1 w-14">
            <span class="flex-1">{{ item.code }}</span>
            <XMarkIcon class="ml-3 w-3 h-4 text-xs" />
          </button>
        </div>
      </div>
      <button v-wave type="button" @click="selectLang"
        class=" w-32 rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
        选择语言
      </button>
    </div>
    <!-- 选择语言弹出层 -->
    <SelectLang></SelectLang>
    <!-- 加载中 -->
    <div ref="loadingRef" v-show="editorStore.isLoadingResult" class="bg-white h-full relative">
      <span>
        <loading v-model:active="editorStore.isLoadingResult" :can-cancel="false" :z-index="30" :is-full-page="false">
          <template v-slot>
            <div class="flex flex-col items-center">
              <div>
                <LoadingIcon class="h-16 w-16 text-gray-400  animate-spin" aria-hidden="true" :fontControlled="false" />
              </div> 
              <p class="mt-4">翻译中时间可能较长请耐心等待  ...</p>
            </div>
          </template>
        </loading>
      </span>
    </div>
    <!-- 代码编辑器 -->
    <div v-show="!editorStore.isLoadingResult" ref="codeEditorRef" class="editor-box"></div>
  </div>
</template>

<script lang="ts" setup>
import * as monaco from "monaco-editor";
import 'monaco-editor/esm/vs/basic-languages/javascript/javascript.contribution';
import Loading from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/css/index.css';
import SelectLang from "./select_lang.vue";
import { XMarkIcon } from '@heroicons/vue/24/outline'
import { useGlobalStore } from "~/store/global";
import { useTranslationStore } from "~/store/translation";
import { setResultsMonacoEditor, useEditorStore } from "~/store/editor"
import { useToast, POSITION } from "vue-toastification";
import LoadingIcon from "~/assets/icons/loading.svg"


const codeEditorRef = ref(null);
const loadingRef = ref(null);
const containerRef = ref(null)
const translationStore = useTranslationStore()
const editorStore = useEditorStore()
const globalStore = useGlobalStore()

const selectLang = () => {
  if (editorStore.isLoadingResult) {
    useToast().warning("正在翻译中请稍后", {
      position: POSITION.BOTTOM_RIGHT
    })
    return false
  }
  translationStore.resultsSelectLangDialogShow = true
}

onMounted(() => {
  //加载主题
  const theme = globalStore.theme == "light" ? "vs" : "brilliance-black"
  codeEditorRef.value.style.height = "calc(100vh - 64px - 5rem - 4.2rem)";
  loadingRef.value.style.height = "calc(100vh - 64px - 5rem - 4.2rem)";
  let resultsExample = monaco.editor.create(codeEditorRef.value as HTMLElement, {
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
    readOnly: true, //是否只读  取值 true | false
    cursorSmoothCaretAnimation: "on", // 是否启用光标平滑插入动画  当你在快速输入文字的时候 光标是直接平滑的移动还是直接"闪现"到当前文字所处位置

  });
  setResultsMonacoEditor(resultsExample)

});


</script>

<style lang="scss" scoped>
:deep(.monaco-editor) {
  border-radius: 1rem;
}

:deep(.overflow-guard) {
  border-radius: 1rem;
}
</style>
