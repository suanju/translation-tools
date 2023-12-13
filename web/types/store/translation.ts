export type ResultsSelectLangDialog = boolean
export interface OriginalLangListItem {
    id: number,
    lang: string,
    code: string,
    original: boolean,
    results: boolean,
    check?: boolean,
}

export interface PlatformItem {
    name: string
    code: string
    default: boolean
}

export type PlatformList = PlatformItem[]

export type OriginalLangList = OriginalLangListItem[]

export type SelectedList = OriginalLangListItem[]

export interface TranslationLangItem {
    lang: string
    code: string
}
export type translationLangList = TranslationLangItem[]
