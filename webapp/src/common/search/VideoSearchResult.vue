<template>
  <div class="video-gallery">
    <div v-for="video in videos" :key="video.vid" class="video-card">
      <img :src="video.thumbnail" alt="视频缩略图" class="thumbnail" />
      <div class="video-info">
        <div class="video-title" v-html="video.highlight_subtitled"></div>
        <div class="video-details">
          <span class="likes">{{ video.likes }} ❤️</span>
          <span class="upload-time">{{ formatUploadTime(video.upload_time) }}前</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {useRoute} from "vue-router";

export default {
  name: "VideoSearchResult",
  setup() {
    const route = useRoute();
    return {
      videos: route.params.videos,
    };
  },
  props: {
    videos: {
      type: Array,
      required: true
    },
  },
  methods: {
    formatUploadTime(timestamp) {
      const date = new Date(timestamp);
      const now = new Date();
      const diff = now - date;
      // 这只是一个简单的示例，你需要根据需要实现逻辑
      const hours = Math.floor(diff / 3600000);
      return `${hours} 小时`;
    },
  },
};
</script>

<style scoped>
.video-gallery {
  position: absolute;
  width: 60%;
  top: 120px;
  left: 20%;
  background-color: #f7f7f7; /* 更轻柔的背景颜色 */
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); /* 响应式布局 */
  gap: 20px;
  padding: 20px; /* 添加内边距 */
}

.video-card {
  position: relative;
  overflow: hidden;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1); /* 添加阴影效果 */
}

.thumbnail {
  width: 100%;
  height: auto;
  display: block; /* 避免底部空白 */
}

.video-info {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  padding: 10px;
  background-color: rgba(0, 0, 0, 0.5); /* 提高透明度 */
  color: #ffffff;
  text-shadow: 0 1px 3px rgba(0,0,0,0.1); /* 文本阴影 */
}

.video-title {
  font-size: 16px; /* 增大字体 */
  margin-bottom: 5px;
}

.video-details {
  display: flex;
  justify-content: space-between;
  font-size: 14px; /* 调整字体大小 */
}
</style>


