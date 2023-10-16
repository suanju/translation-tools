import { UserInfo } from '~/types/store/users';
import { defineStore } from "pinia"


export const useUserStore = defineStore("user", () => {
    let userInfoData = reactive<UserInfo>({
        id: 0,
<<<<<<< HEAD
=======
        username: "",
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
        email: "",
        token: "",
    })

    const setUserInfo = (info: UserInfo) => {
        userInfoData.id = info.id
<<<<<<< HEAD
=======
        userInfoData.username = info.username
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
        userInfoData.email = info.email
        userInfoData.token = info.token
    }

    const loginOut = () => {
        userInfoData.id = 0
<<<<<<< HEAD
=======
        userInfoData.username = ""
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
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