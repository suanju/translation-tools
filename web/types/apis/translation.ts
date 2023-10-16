export interface HttpGetLangListReq { }

export interface HttpGetLangListItem {
    id: number
    lang: string
    code: string
    original: boolean
    results: boolean
    check?: boolean
}
export interface HttpGetLangListResq {
    lang_list: HttpGetLangListItem[]
}