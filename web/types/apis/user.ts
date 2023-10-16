<<<<<<< HEAD
export interface HttpLoginReq {
    email: string
    password: string
}

export interface HttpLoginResq {
    id: number
    email: string
    token: string
}

=======
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
export interface HttpRegisterReq {
    email: string
    password: string
    password_check: string
}

export interface HttpRegisterResq {
<<<<<<< HEAD
    id: number
    email: string
    token: string
=======
    email: string
    password: string
    password_check: string
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
}