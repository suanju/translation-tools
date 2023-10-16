import type { EditorErrerList, EditorTipsShow, ResultsSelectLangDialog, OriginalLangListItem, OriginalLangList, SelectedList } from '~/types/store/editor';
import * as monaco from "monaco-editor";
import { defineStore } from "pinia"
import { httpGetLangList } from '~/apis/translation';


export const useEditorStore = defineStore("editor", () => {
    //编辑器错误列表
    let editorError = ref<EditorErrerList>([])
    //编辑器是否显示错误提示
    let editorTipsShow = ref<EditorTipsShow>(false)
    //响应结果语言查询
    let resultsSelectLangDialogShow = ref<ResultsSelectLangDialog>(false)
    //可选语言数组
    const originalLangList = ref<OriginalLangList>([])
    //选中语言
    const selectedList = ref<SelectedList>([])

    const setEditorError = (info: monaco.editor.IMarker[]) => {
        editorError.value = info
    }
    const getLangList = async () => {
        const data = await httpGetLangList({})
        data.data.lang_list.map((item) => {
            originalLangList.value.push(<OriginalLangListItem>{
                id: item.id,
                lang: item.lang,
                code: item.code,
                original: item.original,
                results: item.results,
                check: false

            })
        })
    }
    const uncheckById = (id: number) => {
        console.log(originalLangList.value)
        originalLangList.value.map((item, index) => {
            if (item.id == id) originalLangList.value[index].check = false
        })
    }

    const uncheckAll = () => {
        originalLangList.value.map((_, index) => {
            originalLangList[index].check = false
        })
    }

    return {
        editorError,
        editorTipsShow,
        resultsSelectLangDialogShow,
        originalLangList,
        selectedList,
        getLangList,
        setEditorError,
        uncheckById,
        uncheckAll
    }

}, {
    persist: false,
})