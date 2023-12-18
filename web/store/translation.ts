import type { ResultsSelectLangDialog, PlatformItem, PlatformList, OriginalLangListItem, OriginalLangList, SelectedList, translationLangList, TranslationLangItem } from '~/types/store/translation';
import { defineStore } from "pinia"
import { httpGetLangList, httpGetPlatformList } from '~/apis/translation';


export const useTranslationStore = defineStore("translation", () => {
    //选择语言弹出层
    let resultsSelectLangDialogShow = ref<ResultsSelectLangDialog>(false)
    //可选接口列表
    const platformList = ref<PlatformList>([])
    //当前选择的接口
    const selectPlatform = ref<PlatformItem>()
    //可选语言数组
    const originalLangList = ref<OriginalLangList>([])
    //可作为原始语言
    const originalTranslationLang = ref<translationLangList>([])
    //选中的原始语言
    const selectedOriginalLang = ref<TranslationLangItem>()
    //选中结果语言(直接在 originalLangList 过滤得到)
    const selectedList = ref<SelectedList>([])
    //接口信息
    const APPID = ref("")
    const KEY = ref("")

    //初始化请求
    const initTranslation = async () => {
        if (!platformList.value.length) {
            await getPlatformList()
            await getLangList()
        }
    }
    const getPlatformList = async () => {
        platformList.value = (await httpGetPlatformList()).data.platform_list
        platformList.value.filter((item) => {
            item.default && (selectPlatform.value = item)
        })
    }

    const getLangList = async () => {
        const data = await httpGetLangList({
            platform: selectPlatform.value.code
        })
        originalLangList.value = []
        data.data.lang_list.map((item) => {
            item.original && (
                originalLangList.value.push(<OriginalLangListItem>{
                    id: item.id,
                    lang: item.lang,
                    code: item.code,
                    original: item.original,
                    results: item.results,
                    check: false
                })
            )
        })
        //默认自动选择
        selectedOriginalLang.value = originalLangList.value.filter((item) =>{
            return item.lang === "自动选择 (AUTO)"
        })[0]
        filterLang()
    }

    const filterLang = () => {
        originalTranslationLang.value = []
        originalLangList.value.forEach((item) => {
            console.log(item.results,item.lang)
            item.results && (originalTranslationLang.value.push({
                lang: item.lang,
                code: item.code
            }));
        })
        selectedList.value = []
    }


    const uncheckById = (code: string) => {
        originalLangList.value.map((item, index) => {
            if (item.code == code) originalLangList.value[index].check = false
        })
    }

    const uncheckAll = () => {
        originalLangList.value.map((_, index) => {
            originalLangList[index].check = false
        })
    }

    return {
        resultsSelectLangDialogShow,
        platformList,
        selectPlatform,
        originalLangList,
        originalTranslationLang,
        selectedOriginalLang,
        selectedList,
        APPID,
        KEY,
        initTranslation,
        getPlatformList,
        getLangList,
        uncheckById,
        uncheckAll
    }

}, {
    persist: true,
})