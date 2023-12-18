import type { EditorErrerList, EditorTipsShow, } from '~/types/store/editor';
import * as monaco from "monaco-editor";
import { defineStore } from "pinia"

//编辑器原始Ref
export let monacoEditor =  <monaco.editor.IStandaloneCodeEditor>{} 
export const setMonacoEditor = (example : monaco.editor.IStandaloneCodeEditor) =>{
    monacoEditor = example
}

//结果编辑器原始Ref
export let resultsMonacoEditor =  <monaco.editor.IStandaloneCodeEditor>{} 
export const setResultsMonacoEditor = (example : monaco.editor.IStandaloneCodeEditor) =>{
    resultsMonacoEditor = example
}

export const useEditorStore = defineStore("editor", () => {
    
    
    //编辑器错误列表
    const editorError = ref<EditorErrerList>([])
    //编辑器是否显示错误提示
    const editorTipsShow = ref<EditorTipsShow>(false)
    //是否正在加载结果
    const isLoadingResult = ref(false)

    const setEditorError = (info: monaco.editor.IMarker[]) => {
        editorError.value = info
    }
    return {
        editorError,
        editorTipsShow,
        isLoadingResult,
        setEditorError
    }

}, {
    persist: false,
})