<template>
    <ul class="flex flex-col items-start justify-center dark:bg-dark-bg-moderate-grey">
        <li class="w-full px-4 py-2 border-b-2 border-slate-100 dark:border-dark-border-primary text-slate-900">
            <div>
                <div class="text-sm dark:text-white">{{ userInfo.email }}</div>
                <div class="text-xs text-slate-500 dark:text-neutral-200"> Basic plan </div>
            </div>
        </li>
        <hr />
        <li class="w-full px-4 py-2 text-sm text-slate-900 dark:text-white hover:bg-slate-100 dark:hover:bg-dark-bg-grey">管理账户</li>
        <li class="w-full px-4 py-2  border-b-2 border-slate-100 dark:border-dark-border-primary text-sm text-slate-900 dark:text-white hover:bg-slate-100 dark:hover:bg-dark-bg-grey">购买字符量</li>
        <li class="w-full px-4 py-2 text-sm text-slate-900 flex items-center dark:text-white hover:bg-slate-100 dark:hover:bg-dark-bg-grey ">退出
            <ArrowLeftOnRectangleIcon class="w-6 h-4" />
        </li>
    </ul>
</template>

<script lang="ts" setup>
import { useUserStore } from "~/store/user";
import {ArrowLeftOnRectangleIcon} from "@heroicons/vue/24/outline";

const userInfo = useUserStore().userInfoData;
</script>