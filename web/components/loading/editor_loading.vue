<template>
    <div class="editor-loading-box flex items-center justify-center">
        <div class="loadingio">
            <div class="ldio">
                <div></div>
                <div></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>

</script>

<style lang="scss" scoped>
.editor-loading-box {
    min-height: calc(100vh - 9.2rem);
}


@keyframes ldio {
    0% {
        top: 96px;
        left: 96px;
        width: 0;
        height: 0;
        opacity: 1;
    }

    100% {
        top: 18px;
        left: 18px;
        width: 156px;
        height: 156px;
        opacity: 0;
    }
}

.ldio div {
    position: absolute;
    border-width: 4px;
    border-style: solid;
    opacity: 1;
    border-radius: 50%;
    animation: ldio 1s cubic-bezier(0, 0.2, 0.8, 1) infinite;
}

.ldio div:nth-child(1) {
    border-color: #7241ff;
    animation-delay: 0s;
}

.ldio div:nth-child(2) {
    border-color: #4f23d6;
    animation-delay: -0.5s;
}

.loadingio {
    width: 14rem;
    height: 14rem;
    display: inline-block;
    overflow: hidden;
}

.ldio {
    width: 100%;
    height: 100%;
    position: relative;
    transform: translateZ(0) scale(1);
    backface-visibility: hidden;
    transform-origin: 0 0;
}

.ldio div {
    box-sizing: content-box;
}
</style> 