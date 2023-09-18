<template>
  <div class="editor-box rounded">
    <!-- 工具栏 -->
    <div class="h-16 flex items-center border rounded-rl-2xl px-6 bg-white">
      <div class="flex items-center flex-auto">
      </div>
      <button type="button"
        className=" w-32 rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
        选择语言
      </button>
    </div>
    <!-- 代码编辑器 -->
    <div ref="codeEditorRef" class="editor-box"></div>
  </div>
</template>

<script lang="ts" setup>
import * as monaco from "monaco-editor";
import 'monaco-editor/esm/vs/basic-languages/javascript/javascript.contribution';

const codeEditorRef = ref(null);

let monacoEditor: monaco.editor.IStandaloneCodeEditor



onMounted(() => {
  codeEditorRef.value.style.height = "calc(100vh - 64px - 5rem - 4.2rem)";
  monacoEditor = monaco.editor.create(codeEditorRef.value as HTMLElement, {
    theme: "vs", // 主题
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
