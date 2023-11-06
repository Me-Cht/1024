<template>
  <div class="home-page">
    <div class="top-nav">
      <router-link to="/" class="nav-link front_page">首页</router-link>
      <router-link to="/videos" class="nav-link">推荐</router-link>
      <div class="search-area">
        <input
          class="search-input"
          placeholder="搜索..."
          v-model="searchKey"
          @keyup.enter="search"
        />
        <button class="search-btn" @click="search">搜索</button>
      </div>
      <router-link to="/upload">
        <img src="../assets/icons/加号.svg" alt="上传" class="upload-product" />

      </router-link>
      <div class="user" @click="handleUserClick">
        <img src="../assets/icons/用户.svg" alt="用户图像" class="user-avatar" />
        <span class="user-text" >{{ username ? username : '登录' }}</span>
      <LoginModule :showModal="showModal"  @close="closeModal" @loggedIn="updateLoginStatus" />
      </div>
    </div>
    <router-view :videos="videoInfo"></router-view>
  </div>
</template>

<script  setup>
import { ref } from "vue";
import LoginModule from "@/common/user/LoginModule.vue";
import {searchVideo} from "@/api/search";
import { useRouter } from 'vue-router'

const router = useRouter()
const showModal = ref(false); //展示登录框
const username = localStorage.getItem("username")
const isLogged= ref(false)
const searchKey = ref(""); //搜索关键字
const offset = ref(0)
const videoInfo = ref([]); // 添加一个用于存放视频数据的响应式引用

// 搜索方法

async function search() {
  console.log('key:', searchKey.value);
  try {
    const response = await searchVideo({ key: searchKey.value,offset:offset.value }); // 假设 searchVideo 已正确接收对象参数
    console.log("输出的值为:",response.data.data.info)
    videoInfo.value = response.data.data.info;
    await router.push({
      name: 'search',
      params: {videos: videoInfo.value}
    });

  } catch (error) {
    console.error('Error fetching videos:', error);
  }
}



function updateLoginStatus(status){
  isLogged.value = status;
}
function handleUserClick(){
  if(isLogged.value){
    router.push('/userinfo');

  }else{
    showModal.value = true;
  }
}
function closeModal() {
  showModal.value = false;
}





</script>


<style scoped lang="scss">
.home-page {
  //display: flex;
  flex-direction: column;
  align-items: center;

  .top-nav {
    z-index: 1000;
    width: 100%;
    background-color: #181818;
    color: white;
    display: flex;
    justify-content: space-between; /* 使子元素在水平方向上均匀分布 */
    align-items: center; /* 使子元素在垂直方向上居中对齐 */
    padding: 10px;
    position: absolute;
    top: 0;
    font-size: 24px;
    text-decoration: none;

    .nav-link {
      color: white;
      text-decoration: none;
    }
    .front_page{
      margin-left: 20%;
    }
    .search-area {
      display: flex; // 使用Flex布局来排列搜索框和按钮
      align-items: center; // 垂直居中对齐搜索框和按钮
      border: 1px solid white; // 给整个搜索区域加上边框
      border-radius: 5px; // 边框圆角效果
      overflow: hidden; // 防止子元素溢出边框
      background-color: #181818; // 搜索区域背景色

    }

    .search-input {
      width: 10%;
      padding: 8px;
      border: none;
      color: white;
      border-radius: 5px;
      outline: none;
      font-size: 16px;
      flex-grow: 1; // 输入框会填充剩余空间
      margin: 0; // 移除外边距
      background-color: transparent; // 输入框背景透明


    }

    .search-btn {
      border: none; // 去除按钮自身的边框
      background-color: #3498db; // 按钮背景颜色
      color: white; // 按钮文字颜色
      padding: 8px 16px; // 按钮内边距
      margin: 0; // 移除外边距
      border-radius: 0; // 由于是在.search-area内，所以移除单独的圆角
      cursor: pointer; // 鼠标悬停时显示指针
      outline: none; // 去除聚焦时的边框
    }
    .search-btn:hover {
      background-color: #2980b9; // 按钮背景颜色变深
    }

    // 输入框聚焦时的样式

    .user {
      width: 120px;
      height: 50px;
      background-color: rgb(254, 44, 85);
      display: flex; /* 使用Flex布局 */
      align-items: center; /* 垂直居中 */
      border-radius: 15px;
      .user-avatar {
        width: 30px;
        height: 30px;
        object-fit: cover;
        border-radius: 50%;
        margin-left: 10px;
      }
      .user-text{
        font-size: 15px;
        margin-left: 10px;
      }
    }

    .user:hover{
      cursor: pointer; /* 当鼠标悬停在元素上时显示小手光标 */
    }


    .upload-product {
      width: 40px; /* 修改上传图像的宽度 */
      height: 40px; /* 修改上传图像的高度 */
      object-fit: cover;
      margin-right: 10px; /* 调整上传图像与输入框之间的间距 */
      background-color: rgb(254, 44, 85);
      /* 添加阴影 */
      box-shadow: 0 4px 8px rgba(254, 44, 85, 0.1); /* x偏移量, y偏移量, 模糊半径, 阴影颜色的透明度 */
    }

    .upload-product:hover {
      content: url("../assets/icons/加号1.svg");
      cursor: pointer; /* 当鼠标悬停在元素上时显示小手光标 */
      background-color: rgb(254, 44, 85);
      /* 添加阴影 */
      box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); /* x偏移量, y偏移量, 模糊半径, 阴影颜色的透明度 */
    }
  }
}
</style>
