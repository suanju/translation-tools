export const useUserStore = defineStore("user", () => {
    let userInfoData = reactive<userInfo>({
        id: 0,
        username: "",
        photo: "",
        token: "",
        created_at: "",
        liveData: {
            address: "",
            key: ""
        },
        //消息相关
        unreadNotice: 0
    })

    const setUserInfo = (info: userInfoRes) => {
        userInfoData.username = info.username
        userInfoData.id = info.id
        userInfoData.photo = info.photo
        userInfoData.token = info.token
        userInfoData.created_at = info.created_at
    }

    const setUnreadNotice = (num: number) => {
        userInfoData.unreadNotice = num
    }

    const loginOut = () => {
        userInfoData.username = ""
        userInfoData.id = 0
        userInfoData.photo = ""
        userInfoData.token = ""
        userInfoData.created_at = ""
        userInfoData.unreadNotice = 0
        //清空消息
        let chat = useChatListStore()
        chat.chatListData = []
        chat.tid = 0
        chat.tUsername = ""
    }
    return {
        userInfoData,
        setUserInfo,
        setUnreadNotice,
        loginOut
    }

}, {
    persist: true,
})