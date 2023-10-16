<template>
  <ClientOnly >
    <div class="main border rounded-2xl">
      <div class="flex">
        <div class="w-1/2">
          <OriginalEditot v-if="monacoReady"></OriginalEditot>
        </div>
        <div class="w-1/2">
          <ResultslEditot v-if="monacoReady"></ResultslEditot>
        </div>
      </div>
      <div>
        <EditorErrorTips></EditorErrorTips>
      </div>
    </div>
    <!-- 加载中 -->
    <template #fallback>
      <EditorLoading></EditorLoading>
    </template>
  </ClientOnly>
</template>

<script lang="ts" setup>

import OriginalEditot from "~/components/original-editor/editor.vue";
import ResultslEditot from "~/components/results-editor/editor.vue";
import EditorErrorTips from "~/components/editor-error-tips/tips.vue";
import EditorLoading from "~/components/loading/editor_loading.vue";


const monacoReady = ref(false);

onNuxtReady(async () => {
  const monacoWorker = (await import("~/utlis/monaco_worker"))
  const initMonacoWorker = monacoWorker.initMonacoWorker;
  const getYamlITextModel = monacoWorker.getYamlITextModel;
  const monaco = (await import ("monaco-editor"));
  console.log(monaco)
  //加载主题
  const theme = (await import('monaco-themes/themes/Brilliance Black.json')).default
  // @ts-ignore 版本问题无法取出相关类型
  monaco.editor.defineTheme('brilliance-black', theme);
  //开启语法提示worker
  initMonacoWorker();
  getYamlITextModel();
  monacoReady.value = true
})

</script>
