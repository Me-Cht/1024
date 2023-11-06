<template>
  <div class="profile-card">
    <div class="profile-header">
      <img :src="user.avatar" alt="user avatar" class="avatar">  <!--  这里放用户头像  -->
      <div class="user-info">
        <h3 class="username">{{ user.name }}</h3>
        <p class="user-details">关注 {{ user.following }} | 粉丝 {{ user.followers }}</p>
      </div>
      <!--   其他信息和按钮等   -->
    </div>
    <!--   其他个人信息   -->
  </div>
</template>

<script setup>
import { ref } from 'vue';

const user = ref({
  name: localStorage.getItem('username'),
  avatar: 'path_to_avatar_image',
  followers: 216, //模拟粉丝数量
  following: 29, // 模拟关注数量
});
</script>

<style scoped>
.profile-card {
  position:relative;
  top:80px;
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.profile-header {
  display: flex;
  align-items: center;
}

.avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  margin-right: 20px;
}

.username {
  margin: 0;
  color: #333;
}

.user-details {
  margin: 0;
  font-size: 0.9rem;
  color: #666;
}

</style>
