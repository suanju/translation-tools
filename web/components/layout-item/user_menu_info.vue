<template>
    <ul class="flex flex-col items-start justify-center">
        <li class="w-full px-4 py-2 border-b-2 border-slate-100 text-slate-900">
            <div>
                <div class="text-sm">{{ userInfo.email }}</div>
                <div class="text-xs text-slate-500"> Basic plan </div>
            </div>
        </li>
        <hr />
        <li class="w-full px-4 py-2 text-sm text-slate-900">管理账户</li>
        <li class="w-full px-4 py-2  border-b-2 border-slate-100 text-sm text-slate-900">购买字符量</li>
        <li class="w-full px-4 py-2 text-sm text-slate-900 flex items-center ">退出
            <ArrowLeftOnRectangleIcon class="w-6 h-4" />
        </li>
    </ul>
</template>

<script lang="ts" setup>
import { useUserStore } from "~/store/user";

const userInfo = useUserStore().userInfoData;
</script>