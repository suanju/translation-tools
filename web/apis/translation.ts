import type { HttpGetLangListReq, HttpGetLangListResq } from '~/types/apis/translation';


export const httpGetLangList = (params: HttpGetLangListReq) => {
    return useHttp().get<HttpGetLangListResq>('/translation/get_lang_list', params);
}
