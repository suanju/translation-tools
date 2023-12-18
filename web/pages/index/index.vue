<template>
  <ClientOnly>
    <div class="main border rounded-2xl dark:border-dark-border-primary">
      <div class="flex">
        <!-- 切换按钮 -->
        <div v-wave @click="submit"
          class="w-28 h-28 bg-indigo-600 z-50 absolute rounded-full flex items-center justify-center"
          style="left:calc(50% + 5.2rem); top:56%;">
          <switchIcon class="w-16 h-16 text-white" :fontControlled="false" filled></switchIcon>
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
import { useToast, POSITION } from "vue-toastification";

const monacoReady = ref(false);
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
  if (!translationStore.selectedList.length) {
    useToast().error("未添加翻译语言", {
      position: POSITION.BOTTOM_RIGHT
    })
    return false;
  }
  if (translationStore.selectedList.length == 1) {
    editorStore.isLoadingResult = true
    //翻译单语言
    try {
      const resp = await httpGetTranslationJson(<HttpGetTranslationJsonReq>{
        isKeyAsTr: true,
        original: translationStore.selectedOriginalLang.code,
        results: translationStore.selectedList[0].code,
        json: monacoEditor.getValue()
      })
      const text = resp.data.json
      resultsMonacoEditor.updateOptions({ readOnly: false });
      resultsMonacoEditor.setValue(text)
      await resultsMonacoEditor.getAction("editor.action.formatDocument").run()
      resultsMonacoEditor.updateOptions({ readOnly: false });
      useToast().success("翻译成功", {
        position: POSITION.BOTTOM_RIGHT
      })
    } catch (err) {
      editorStore.isLoadingResult = false
      useToast().error(err.msg ? err.msg : "发生异常", {
        position: POSITION.BOTTOM_RIGHT
      })
    }
  } else {
    //翻译多语言
    useToast().warning("暂不支持翻译多种类", {
        position: POSITION.BOTTOM_RIGHT
    })
    return false
  }

  editorStore.isLoadingResult = false
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
