<template>
  <ClientOnly>
    <div class="main border rounded-2xl dark:border-dark-border-primary">
      <div class="flex">
        <!-- 切换按钮 -->
        <div v-ripple @click="submit" class="w-24 h-24 bg-indigo-600 z-50 absolute rounded-full flex items-center justify-center"
          style="left:calc(50% + 5.2rem); top:56%">
          <switchIcon class="w-12 h-12 text-white" :fontControlled="false" filled></switchIcon>
        </div>
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
import switchIcon from "~/assets/icons/switch.svg"
import { useEditorStore, monacoEditor, resultsMonacoEditor } from "~/store/editor";
import { httpGetTranslationJson } from "~/apis/translation";
import type { HttpGetTranslationJsonReq } from "~/types/apis/translation";
import { useTranslationStore } from "~/store/translation";
import { useToast,POSITION } from "vue-toastification";

const monacoReady = ref(false);
const isLoading = ref(true)
const editorStore = useEditorStore()
const translationStore = useTranslationStore()

const submit = async () => {
  resultsMonacoEditor.setValue("")
  if (editorStore.editorError.length) {
    editorStore.editorTipsShow = true
    return false
  }
  monacoEditor.getAction("editor.action.formatDocument").run()
  //进行请求转化  
  //值翻译单语言
  const resp = await httpGetTranslationJson(<HttpGetTranslationJsonReq>{
    isKeyAsTr: true,
    original: translationStore.selectedOriginalLang.code,
    results: 'en',
    json: JSON.parse(monacoEditor.getValue())
  })
  const text = JSON.stringify(resp.data.json)
  resultsMonacoEditor.setValue(text)
  resultsMonacoEditor.getAction("editor.action.formatDocument").run()
  useToast().success("翻译成功",{
    position:POSITION.BOTTOM_RIGHT
  })

}


onNuxtReady(async () => {
  const monacoWorker = (await import("~/utlis/monaco_worker"))
  const initMonacoWorker = monacoWorker.initMonacoWorker;
  const getYamlITextModel = monacoWorker.getYamlITextModel;
  const monaco = (await import("monaco-editor"));
  //加载主题
  // const theme = (await import('monaco-themes/themes/Twilight.json')).default
  const theme = (await import('~/assets/monaco/themes/dark.json')).default
  // @ts-ignore 版本问题无法取出相关类型
  monaco.editor.defineTheme('brilliance-black', theme);
  //开启语法提示worker
  initMonacoWorker();
  getYamlITextModel();
  monacoReady.value = true
})

</script>
