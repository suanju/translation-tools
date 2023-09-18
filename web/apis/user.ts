import { HttpRegisterReq, HttpRegisterResq, HttpLoginReq, HttpLoginResq } from '~/types/apis/user';


export const httpLogin = (params: HttpLoginReq) => {
    return useHttp().post<HttpLoginResq>('/user/login', params);
}

export const httpRegiste = (params: HttpRegisterReq) => {
    return useHttp().post<HttpRegisterResq>('/user/register', params);
}