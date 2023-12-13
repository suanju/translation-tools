<template>
  <NuxtPage />
</template>

<script setup>
import { useTranslationStore } from "~/store/translation"

const translationStore = useTranslationStore()
onMounted(async () => {
  await translationStore.initTranslation()
})
</script>