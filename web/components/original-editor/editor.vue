<template>
  <div class="editor-box rounded">
    <!-- 工具栏 -->
    <div class="h-16 flex items-center border rounded-tl-2xl px-6 bg-white">
      <div class="flex items-center flex-auto">
        <Listbox as="div" v-model="langSelected">
          <div class="relative">
            <ListboxButton
              class="relative w-full cursor-default rounded-md bg-white py-1.5 pl-3 pr-10 text-left text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6">
              <span class="block truncate font-extralight">{{ langSelected.name }}</span>
              <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
              </span>
            </ListboxButton>

            <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100"
              leave-to-class="opacity-0">
              <ListboxOptions
                class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                <ListboxOption as="template" v-for="lang in langArr" :key="lang.id" :value="lang"
                  v-slot="{ active, langSelected }">
                  <li :class="[
                    active ? 'bg-indigo-600 text-white' : 'text-gray-900',
                    'relative cursor-default select-none py-2 pl-3 pr-7',
                  ]">
                    <span :class="[
                      langSelected ? 'font-semibold' : 'font-normal',
                      'block truncate',
                    ]" class="font-extralight">{{ lang.name }}</span>
                    <span v-if="langSelected" :class="[
                      active ? 'text-white' : 'text-indigo-600',
                      'absolute inset-y-0 right-0 flex items-center pr-4',
                    ]">
                      <CheckIcon class="h-5 w-5" aria-hidden="true" />
                    </span>
                  </li>
                </ListboxOption>
              </ListboxOptions>
            </transition>
          </div>
        </Listbox>
        <Listbox as="div" class="ml-2" v-model="formatSelected">
          <div class="relative">
            <ListboxButton
              class="relative w-full cursor-default rounded-md bg-white py-1.5 pl-3 pr-10 text-left text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6">
              <span class="block truncate font-extralight">{{
                formatSelected.language
              }}</span>
              <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
              </span>
            </ListboxButton>

            <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100"
              leave-to-class="opacity-0">
              <ListboxOptions
                class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                <ListboxOption as="template" v-for="format in formatArr" :key="format.id" :value="format"
                  v-slot="{ active, formatSelected }">
                  <li :class="[
                    active ? 'bg-indigo-600 text-white' : 'text-gray-900',
                    'relative cursor-default select-none py-2 pl-3 pr-7',
                  ]">
                    <span :class="[
                      formatSelected ? 'font-semibold' : 'font-normal',
                      'block truncate',
                    ]" class="font-extralight">{{ format.language }}</span>
                    <span v-if="formatSelected" :class="[
                      active ? 'text-white' : 'text-indigo-600',
                      'absolute inset-y-0 right-0 flex items-center pr-4',
                    ]">
                      <CheckIcon class="h-5 w-5" aria-hidden="true" />
                    </span>
                  </li>
                </ListboxOption>
              </ListboxOptions>
            </transition>
          </div>
        </Listbox>
      </div>

      <button type="button"
        class=" w-32 rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
        上传文件
      </button>
    </div>
    <!-- 代码编辑器 -->
    <div ref="codeEditorRef" class="editor-box"></div>
  </div>
</template>

<script lang="ts" setup>
import * as monaco from "monaco-editor";
import 'monaco-editor/esm/vs/basic-languages/javascript/javascript.contribution';

import {
  Listbox,
  ListboxButton,
  ListboxOption,
  ListboxOptions,
} from "@headlessui/vue";
import { CheckIcon, ChevronUpDownIcon } from "@heroicons/vue/20/solid";
import { useEditorStore } from "~/store/editor"
import { useGlobalStore } from "~/store/global";

const codeEditorRef = ref(null);
const editorInfo = useEditorStore()

const langArr = [
  { id: 1, name: "中文 (ZH)" },
  { id: 2, name: "英语 (EN)" },
  { id: 3, name: "德语 (DE)" },
];

const formatArr = [
  { id: 1, language: "json" },
  { id: 2, language: "yaml" },
];

const langSelected = ref(langArr[0]);
const formatSelected = ref(formatArr[0]);
const globalInfo = useGlobalStore()

let monacoEditor: monaco.editor.IStandaloneCodeEditor

watch(formatSelected, (val) => {
  monaco.editor.setModelLanguage(monacoEditor.getModel(), val.language);
})


onMounted(() => {
  //加载主题
  const theme = globalInfo.theme == "light" ? "vs" : "brilliance-black"
  codeEditorRef.value.style.height = "calc(100vh - 64px - 5rem - 4.2rem)";
  monacoEditor = monaco.editor.create(codeEditorRef.value as HTMLElement, {
    theme, // 主题
    value: "{{'得到' ：  ，， '$schema':  1} 12}", // 默认显示的值
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

  //记录错误信息
  monaco.editor.onDidChangeMarkers(([uri]) => {
    const markers = monaco.editor.getModelMarkers({ resource: uri })
    editorInfo.setEditorError(markers)
    console.log('markers:', markers.map(
      ({ message, startLineNumber, startColumn, endLineNumber, endColumn }) =>
        `${message} [${startLineNumber}:${startColumn}-${endLineNumber}:${endColumn}]`,
    ))
  })

});


</script>

<style></style>
