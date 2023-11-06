<template>
  <div class="video-container" @wheel="handleWheel">
    <div class="buttons">
      <button class="btns" @click="switchToVideo('prev')">上一个</button>
      <button class="btns" @click="switchToVideo('next')">下一个</button>
    </div>
    <transition name="fade" mode="out-in">
      <video class="myVideo" :key="currentVideo" :src="currentVideoSource" type="video/mp4" @ended="switchVideo"></video>
    </transition>
  </div>
</template>

<script setup>
import { ref } from "vue";

const videos = ['./tmp/video2.mp4', './tmp/video3.mp4'];
const currentVideoIndex = ref(0);

const currentVideoSource = computed(() => {
  return videos[currentVideoIndex.value];
});

const switchToVideo = (direction) => {
  if (direction === 'next') {
    currentVideoIndex.value = (currentVideoIndex.value + 1) % videos.length;
  } else if (direction === 'prev') {
    currentVideoIndex.value = (currentVideoIndex.value - 1 + videos.length) % videos.length;
  }
};

const switchVideo = () => {
  switchToVideo('next'); // 视频播放结束后切换到下一个视频
};

const handleWheel = (event) => {
  if (event.deltaY > 0) {
    switchToVideo('next'); // 滚轮向下滚动，切换到下一个视频
  } else {
    switchToVideo('prev'); // 滚轮向上滚动，切换到上一个视频
  }
  event.preventDefault(); // 阻止滚轮事件的默认行为
};
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
.video-container {
  height: 50%;
  width: 100vw;
  position: relative;
  overflow: hidden;
}

.buttons {
  position: absolute;
  top: 80px;
  left: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  z-index: 999;
}

.myVideo {
  width: 80%;
  height: 100%;
}

.btns {
  width: 90px;
  height: 40px;
  background-color: #4caf50;
  border: none;
  color: white;
  cursor: pointer;
}

.btns:hover {
  background-color: #45a049;
}
</style>
