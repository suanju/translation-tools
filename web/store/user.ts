import { UserInfo } from '~/types/store/users';
import { defineStore } from "pinia"


export const useUserStore = defineStore("user", () => {
    let userInfoData = reactive<UserInfo>({
        id: 0,
        username: "",
        email: "",
        token: "",
    })

    const setUserInfo = (info: UserInfo) => {
        userInfoData.id = info.id
        userInfoData.username = info.username
        userInfoData.email = info.email
        userInfoData.token = info.token
    }

    const loginOut = () => {
        userInfoData.id = 0
        userInfoData.username = ""
        userInfoData.email = ""
        userInfoData.token = ""
    }
    return {
        userInfoData,
        setUserInfo,
        loginOut
    }

}, {
    persist: true,
})