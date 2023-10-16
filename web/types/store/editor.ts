import * as monaco from "monaco-editor";

export type EditorErrerList =  monaco.editor.IMarker[]
export type EditorTipsShow = boolean
export type ResultsSelectLangDialog = boolean

export interface OriginalLangListItem {
    id:number,
    lang:string,
    code:string,
    original: boolean,
    results: boolean,
    check?:boolean,
}

export type OriginalLangList = OriginalLangListItem[]

export type SelectedList = OriginalLangListItem[]