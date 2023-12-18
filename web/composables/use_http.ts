import axios from 'axios';
import type { AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig } from 'axios';
import { useToast, POSITION } from "vue-toastification";
import { useTranslationStore } from '~/store/translation';

interface Result {
    code: number;
    message: string
}

// 请求响应参数，包含data
interface ResultData<T = any> extends Result {
    data?: T;
}


enum RequestEnums {
    TIMEOUT = 60000,//置超时时间
    SUCCESS = 200, // 请求成功
    OPERATIONFAIL = 500, // 操作失败
    NOTLOGIN = 303, // 操作失败
    FAIL = 999, // 请求失败
}


class RequestHttp {
    service: AxiosInstance;
    public constructor() {
        //获取公共配置
        const translationStore = useTranslationStore()
        const runtimeConfig = useRuntimeConfig();
        const config = <AxiosRequestConfig>{
            baseURL: runtimeConfig.public.baseApi,
            timeout: RequestEnums.TIMEOUT as number,
        }
        this.service = axios.create(config);

        /**
         * 请求拦截器
         */
        this.service.interceptors.request.use(
            (config: InternalAxiosRequestConfig) => {
                translationStore.selectPlatform?.code && (config.headers['platform_code'] = translationStore.selectPlatform.code)
                translationStore.APPID && (config.headers['app_id'] = translationStore.APPID)
                translationStore.KEY && (config.headers['app_key'] = translationStore.KEY)
                return {
                    ...config,
                }
            },
            (error: AxiosError) => {
                // 请求报错
                Promise.reject(error)
            }
        )

        /**
         * 响应拦截器
         */
        this.service.interceptors.response.use(
            (response: AxiosResponse) => {
                const { data } = response;
                if (data.code == RequestEnums.OPERATIONFAIL) {
                    return Promise.reject(data);
                }
                // 全局错误信息拦截（防止下载文件得时候返回数据流，没有code，直接报错）
                if (data.code && data.code !== RequestEnums.SUCCESS) {
                    console.error(data)
                    return Promise.reject(data)
                }
                return data;
            },
            (error: AxiosError) => {
                console.log(error)
                if (error.code == "ERR_NETWORK") {
                    return Promise.reject(error)
                }
                const { response } = error;
                console.log(response)
                if (response) {
                    this.handleCode(response.status)
                }
                if (!window.navigator.onLine) {
                    useToast().error("网络异常", {
                        position: POSITION.BOTTOM_RIGHT
                    })
                }
            }
        )
    }

    handleCode(code: number): void {
        switch (code) {
            default:
                //错误响应
                break;
        }
    }

    // 常用方法封装
    get<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.get(url, { params });
    }
    post<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.post(url, params);
    }
    put<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.put(url, params);
    }
    delete<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.delete(url, { params });
    }
}
export const useHttp = () => new RequestHttp(); 