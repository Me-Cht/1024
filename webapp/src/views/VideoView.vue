<template>
    <div class="video-container">
        <div class="video-warpper">
            <video class="myvideo" controls>
                <source v-bind:src="url" type="video/mp4">
            </video>
        </div>
        <div class="switch-button-container">
            <button class="floating-button-up" id="prev" @click='btnUpClicked'></button>
            <button class="floating-button-down" id="next" @click="btnDownClicked"></button>
        </div>
    </div>
</template>

<style scoped>
.video-container {
    width: 400px;
    height: 750px;
    position: relative;
}

.switch-button-container {
    position: absolute;
    top: 50%;
    right: 0;
    display: flex;
    flex-direction: column;
    align-items: flex-end;
}

.floating-button-up {
    width: 50px;
    height: 50px;
    text-align: center;
    line-height: 50px;
    /* 居中按钮文字 */
    margin: 5px;
    border-radius: 50%;
    /* 将按钮设置为圆形 */
    background-color: white;
    /* 设置按钮背景颜色 */
    background-repeat: no-repeat;
    /* 防止图片重复 */
    background-image: url('../assets/arrow-up.png');
    /* 设置箭头图片的路径 */
    background-position: center;
    /* 将图片置于按钮中央 */
    background-size: 20px;
    /* 自动缩放背景图片以适应按钮 */
    border: none;
    /* 去掉按钮的边框 */
}

.floating-button-down {
    width: 50px;
    height: 50px;
    text-align: center;
    line-height: 50px;
    /* 居中按钮文字 */
    margin: 5px;
    border-radius: 50%;
    /* 将按钮设置为圆形 */
    background-color: white;
    /* 设置按钮背景颜色 */
    background-repeat: no-repeat;
    /* 防止图片重复 */
    background-image: url('../assets/arrow-down.png');
    /* 设置箭头图片的路径 */
    background-position: center;
    /* 将图片置于按钮中央 */
    background-size: 20px;
    /* 自动缩放背景图片以适应按钮 */
    border: none;
    /* 去掉按钮的边框 */
}

.myvideo {
    width: 400px;
    height: 750px;
}

.video-warpper {
    width: 400px;
    height: 750px;
}
</style>

<script setup>
import {ref} from 'vue'
const url = ref('./tmp/video1.mp4')
const record=[]
function btnUpClicked() {
    url.value = "./tmp/video1.mp4"
    const video = document.querySelector('.myvideo')
    video.load()
}

function btnDownClicked() {
    url.value = "./tmp/video2.mp4"
    const video = document.querySelector('.myvideo')
    video.load()
}
</script>
