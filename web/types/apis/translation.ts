
export interface HttpGetPlatformListItem {
    name: string
    code: string
    default: boolean
}

export interface HttpGetPlatformListResp {
    platform_list: HttpGetPlatformListItem[]
}

export interface HttpGetLangListReq {
    platform: string
}

export interface HttpGetLangListItem {
    id: number
    lang: string
    code: string
    original: boolean
    results: boolean
    check?: boolean
}

export interface HttpGetLangListResp {
    lang_list: HttpGetLangListItem[]
}

export interface HttpGetTranslationJsonReq {
    isKeyAsTr: boolean
    original: string
    results: string
    json: Record<string, string>
}

export interface HttpGetTranslationJsonResp {
    json: Record<string, string>
}