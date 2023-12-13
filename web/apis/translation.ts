import type { HttpGetLangListReq, HttpGetLangListResp , HttpGetPlatformListResp,HttpGetTranslationJsonReq,HttpGetTranslationJsonResp } from '~/types/apis/translation';



export const httpGetPlatformList = () => {
    return useHttp().get<HttpGetPlatformListResp>('/lang/get_platform_list');
}
export const httpGetLangList = (params: HttpGetLangListReq) => {
    return useHttp().get<HttpGetLangListResp>('/lang/get_lang_list', params);
}

export const httpGetTranslationJson = (params: HttpGetTranslationJsonReq) => {
    return useHttp().post<HttpGetTranslationJsonResp>('/translation/translation_json', params);
}
