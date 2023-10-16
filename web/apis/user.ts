<<<<<<< HEAD
import { HttpRegisterReq, HttpRegisterResq, HttpLoginReq, HttpLoginResq } from '~/types/apis/user';


export const httpLogin = (params: HttpLoginReq) => {
    return useHttp().post<HttpLoginResq>('/user/login', params);
}

export const httpRegiste = (params: HttpRegisterReq) => {
    return useHttp().post<HttpRegisterResq>('/user/register', params);
=======
import { HttpRegisterReq, HttpRegisterResq } from '~/types/apis/user';

export const httpRegiste = (params: HttpRegisterReq) => {
    return useHttp().post<HttpRegisterResq>('/user/login', params);
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
}