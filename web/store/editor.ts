import { EditorErrerList,EditorTipsShow } from '~/types/store/editor';
import * as monaco from "monaco-editor";
import { defineStore } from "pinia"


export const useEditorStore = defineStore("editor", () => {
    let editorError= ref<EditorErrerList>([])
    let editorTipsShow = ref<EditorTipsShow>(false)

    const setEditorError = (info:  monaco.editor.IMarker[]) => {
        editorError.value = info
    }
    
    return {
        editorError,
        editorTipsShow,
        setEditorError
    }

}, {
    persist: false,
})