<template>
    <div class="flex min-h-full flex-1 flex-col justify-center py-12 sm:px-6 lg:px-8">
        <div class="sm:mx-auto sm:w-full sm:max-w-md">
            <img class="mx-auto h-10 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
                alt="Your Company" />
            <h2 class="mt-6 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                注册您的账户
            </h2>
        </div>
        <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-[480px]">
            <div class="bg-white px-6 py-12 shadow sm:rounded-lg sm:px-12">
<<<<<<< HEAD
                <Form class="space-y-6" @submit="submit" :validation-schema="schema" v-slot="{ errors }">
=======
                <Form class="space-y-6" @submit="submit" :validation-schema="schema" v-slot="{ meta, errors }">
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
                    <div>
                        <label for="email" name="email"
                            class="block text-sm font-medium leading-6 text-gray-900">Email</label>
                        <div class="relative mt-2 rounded-md shadow-sm">
<<<<<<< HEAD
                            <Field type="email" name="email" id="email" :class="{
                                ...importClass,
                                'placeholder:text-red-300': errors.email,
                                'ring-red-500': errors.email,
                                'text-red-900': errors.email,
                                'focus:ring-red-500': errors.email,
                            }" v-model="form.email" />
=======
                            <Field type="email" name="email" id="email"
                                :class="{ 'placeholder:text-red-300': errors.email, 'ring-red-500': errors.email, 'text-red-900': errors.email, 'focus:ring-red-500': errors.email }"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                v-model="form.email" />
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
                            <div v-show="errors.email"
                                class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
                                <ExclamationCircleIcon class="h-5 w-5 text-red-500" aria-hidden="true" />
                            </div>
<<<<<<< HEAD
=======
                        </div>
                        <ErrorMessage name="email">
                            <p class="mt-2 text-xs text-red-600 font-extralight" id="email-error">{{ errors.email }}</p>
                        </ErrorMessage>
                    </div>
                    <div>
                        <label for="password" name="password"
                            class="block text-sm font-medium leading-6 text-gray-900">密码</label>
                        <div class="relative mt-2 rounded-md shadow-sm">
                            <Field type="password" autocomplete="password" name="password" id="password"
                                :class="{ 'placeholder:text-red-300': errors.password, 'ring-red-500': errors.password, 'text-red-900': errors.password, 'focus:ring-red-500': errors.password }"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                v-model="form.password" />
                            <div v-show="errors.password"
                                class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
                                <ExclamationCircleIcon class="h-5 w-5 text-red-500" aria-hidden="true" />
                            </div>
                        </div>
                        <ErrorMessage name="password">
                            <p class="mt-2 text-xs text-red-600 font-extralight" id="email-error">{{ errors.password }}</p>
                        </ErrorMessage>
                    </div>


                    <div>
                        <label for="password_check" name="password_check"
                            class="block text-sm font-medium leading-6 text-gray-900">确认密码</label>
                        <div class="relative mt-2 rounded-md shadow-sm">
                            <Field type="password" name="password_check" id="password_check" autocomplete="password_check"
                                :class="{ 'placeholder:text-red-300': errors.password_check, 'ring-red-500': errors.password_check, 'text-red-900': errors.password_check, 'focus:ring-red-500': errors.password_check }"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                v-model="form.password_check" />
                            <div v-show="errors.password_check"
                                class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
                                <ExclamationCircleIcon class="h-5 w-5 text-red-500" aria-hidden="true" />
                            </div>
                        </div>
                        <ErrorMessage name="password_check">
                            <p class="mt-2 text-xs text-red-600 font-extralight" id="email-error">{{ errors.password_check
                            }}</p>
                        </ErrorMessage>
                    </div>


                    <div class="flex items-center justify-between">
                        <div class="flex items-center">
                            <input id="remember-me" name="remember-me" type="checkbox"
                                class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600" />
                            <label for="remember-me" class="ml-3 block text-sm leading-6 text-gray-900">记住密码</label>
                        </div>

                        <div class="text-sm leading-6">
                            <a href="#" class="font-semibold text-indigo-600 hover:text-indigo-500">忘记密码</a>
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
                        </div>
                        <ErrorMessage name="email">
                            <p class="mt-2 text-xs text-red-600 font-extralight" id="email-error">
                                {{ errors.email }}
                            </p>
                        </ErrorMessage>
                    </div>
                    <div>
                        <label for="password" name="password"
                            class="block text-sm font-medium leading-6 text-gray-900">密码</label>
                        <div class="relative mt-2 rounded-md shadow-sm">
                            <Field type="password" autocomplete="password" name="password" id="password" :class="{
                                ...importClass,
                                'placeholder:text-red-300': errors.password,
                                'ring-red-500': errors.password,
                                'text-red-900': errors.password,
                                'focus:ring-red-500': errors.password,
                            }" v-model="form.password" />
                            <div v-show="errors.password"
                                class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
                                <ExclamationCircleIcon class="h-5 w-5 text-red-500" aria-hidden="true" />
                            </div>
                        </div>
                        <ErrorMessage name="password">
                            <p class="mt-2 text-xs text-red-600 font-extralight" id="email-error">
                                {{ errors.password }}
                            </p>
                        </ErrorMessage>
                    </div>
<<<<<<< HEAD
                    <div>
                        <label for="password_check" name="password_check"
                            class="block text-sm font-medium leading-6 text-gray-900">确认密码</label>
                        <div class="relative mt-2 rounded-md shadow-sm">
                            <Field type="password" name="password_check" id="password_check" autocomplete="password_check"
                                :class="{
                                    ...importClass,
                                    'placeholder:text-red-300': errors.password_check,
                                    'ring-red-500': errors.password_check,
                                    'text-red-900': errors.password_check,
                                    'focus:ring-red-500': errors.password_check,
                                }" v-model="form.password_check" />
                            <div v-show="errors.password_check"
                                class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
                                <ExclamationCircleIcon class="h-5 w-5 text-red-500" aria-hidden="true" />
                            </div>
                        </div>
                        <ErrorMessage name="password_check">
                            <p class="mt-2 text-xs text-red-600 font-extralight" id="email-error">
                                {{ errors.password_check }}
                            </p>
                        </ErrorMessage>
                    </div>
                    <div>
                        <button id="submit" type="submit" :disabled="isLoading"
                            class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                            <span v-show="isLoading">
                                <LoadingIcon class="icon animate-spin w-6 h-6" :fontControlled="false" />
                            </span>
                            <span v-show="!isLoading">register</span>
                        </button>
                    </div>
                </Form>
=======
                </Form>

>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
                <div>
                    <div class="relative mt-10">
                        <div class="absolute inset-0 flex items-center" aria-hidden="true">
                            <div class="w-full border-t border-gray-200" />
                        </div>
                    </div>
                </div>
            </div>
            <p class="mt-10 text-center text-sm text-gray-500">
                已有账户 ?
                {{ " " }}
                <span @click="router.push({ name: 'login' })"
                    class="font-semibold leading-6 text-indigo-600 hover:text-indigo-500">去登入</span>
            </p>
        </div>
    </div>
</template>
<<<<<<< HEAD

<script lang="ts" setup>
import { RegisetrForm } from "~/types/pages/login/register";
import { Field, Form, ErrorMessage } from "vee-validate";
import { ExclamationCircleIcon } from "@heroicons/vue/20/solid";
import { httpRegiste } from "~/apis/user";
import { HttpRegisterReq } from "~/types/apis/user";
import * as yup from "yup";
import Swal from "sweetalert2";
import { useUserStore } from "~/store/user";
import { UserInfo } from "~/types/store/users";
import LoadingIcon from "~/assets/icons/loading.svg"

const router = useRouter();
const isLoading = ref(false);
const userSotre = useUserStore();
=======
  
  
<script lang="ts"  setup>
import { RegisetrForm } from "~/types/pages/login/register"
import { Field, Form, ErrorMessage } from 'vee-validate';
import { ExclamationCircleIcon } from '@heroicons/vue/20/solid'
import { httpRegiste } from "~/apis/user"
import { HttpRegisterReq } from '~/types/apis/user';
import * as yup from 'yup';

const router = useRouter()

>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
const form = reactive<RegisetrForm>({
    email: "",
    password: "",
    password_check: "",
<<<<<<< HEAD
});

const submit = async () => {
    try {
        isLoading.value = true;
        const response = await httpRegiste(<HttpRegisterReq>{
            email: form.email,
            password: form.password,
            password_check: form.password_check,
        });
        console.log(response);
        isLoading.value = false;
        Swal.fire({
            title: "注册成功",
            heightAuto: false,
            // confirmButtonColor: globalScss.colorButtonTheme,
            icon: "success",
        });
        userSotre.setUserInfo(<UserInfo>{
            id: response.data?.id,
            email: response.data?.email,
            token: response.data?.token,
        });
        router.push({ name: "hone" });
    } catch (err) {
        isLoading.value = false;
        Swal.fire({
            title: (err as Error).message ? (err as Error).message : (err as any).msg,
            heightAuto: false,
            // confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        });
    }
};

const schema = yup.object({
    email: yup.string().required("请输入您的邮箱").email("您的邮箱格式不正确"),
    password: yup.string().required("请输入您的密码").min(6, "密码长度最短为6位"),
    password_check: yup
        .string()
        .required("请输入重复您的密码")
        .oneOf([yup.ref("password")], "重复密码错误"),
});

const importClass = [
    "block",
    "w-full",
    "rounded-md",
    "border-0",
    "py-1.5",
    "text-gray-900",
    "shadow-sm",
    "ring-1",
    "ring-inset",
    "ring-gray-300",
    "placeholder:text-gray-400",
    "focus:ring-2",
    "focus:ring-inset",
    "focus:ring-indigo-600",
    "sm:text-sm",
    "sm:leading-6",
].reduce((acc, current) => {
    acc[current as string] = true;
    return acc;
}, {} as Record<string, boolean>);
</script>
=======

})

const { baseApi } = useRuntimeConfig();
console.log(baseApi)

const submit = async () => {
    const response = await httpRegiste(<HttpRegisterReq>{
        email: form.email,
        password: form.password,
        password_check: form.password_check
    })
    console.log(response);
}


const schema = yup.object({
    email: yup.string().required('请输入您的邮箱').email("您的邮箱格式不正确"),
    password: yup.string().required('请输入您的密码').min(6, "密码长度最短为6位"),
    password_check: yup.string().required("请输入重复您的密码").oneOf([yup.ref("password")], "重复密码错误"),
});

</script>
>>>>>>> bf70a4241f55b397fd46fd414aae75dadfd2966e
