<template>
  <div class="main border rounded-2xl">
    <ClientOnly fallback-tag="span" fallback="Loading comments...">
      <div class="flex">
        <div class="flex-1"><OriginalEditot></OriginalEditot></div>
        <div class="flex-auto"><ResultslEditot></ResultslEditot></div>
      </div>
      <div>
      <EditorErrorTips></EditorErrorTips>
      </div>
    </ClientOnly>
  </div>
</template>

<script lang="ts" setup>
import OriginalEditot from "~/components/original-editor/editor.vue";
import ResultslEditot from "~/components/results-editor/editor.vue";
import EditorErrorTips from "~/components/editor-error-tips/tips.vue"
</script>
