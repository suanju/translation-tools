export interface HttpLoginReq {
    email: string
    password: string
}

export interface HttpLoginResq {
    id: number
    email: string
    token: string
}

export interface HttpRegisterReq {
    email: string
    password: string
    password_check: string
}

export interface HttpRegisterResq {
    id: number
    email: string
    token: string
}