import { HttpRegisterReq, HttpRegisterResq } from '~/types/apis/user';

export const httpRegiste = (params: HttpRegisterReq) => {
    return useHttp().post<HttpRegisterResq>('/user/login', params);
}