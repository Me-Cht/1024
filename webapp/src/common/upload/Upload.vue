<template>
  <div class="page">
    <div class="top-banner">
      <div class="title">短片精选保障</div>
      <div class="subtitle">收录率高达100%</div>
      <!-- 其他信息 -->
    </div>
    <div class="upload-section">

      <div class="input-section">
        <input v-model="subtitled" placeholder="输入视频标题"/>
        <input v-model="tagsInput" placeholder="输入标签，用逗号分隔"/>
        <input v-model="uploaderId" placeholder="输入上传者ID"/>

      </div>
      <div class="progress-container" v-if="progress < 100">
        <div class="progress-bar" :style="{ width: progress + '%' }"></div>
      </div>
      <div class="upload-box">
        <input type="file" @change="handleFileChange" id="file-input"/>
        <label for="file-input" class="upload-label">选择文件</label>
        <button @click="uploadFile" class="upload-btn">上传</button>
      </div>
      <div v-if="!file" class="file-hint">
        请上传您的视频以确保不会超时，平均每40秒的视频处理时间约2分钟
      </div>
    </div>
  </div>
</template>

<script>
import * as qiniu from "qiniu-js";
import {callBackUpload, getUploadTokenAndKey} from '@/api/upload';


export default {

  name: "Upload",
  data() {
    return {
      progress:0,
      file: null,
      subtitled: "1",     //   用户输入的标题
      tagsInput: "1",     //    用户输入的标签（逗号分隔）
      uploaderId: parseInt(localStorage.getItem('uid')),

    };
  },
  methods: {
    handleFileChange(event) {
      this.file = event.target.files[0];
    },
    uploadFile() {
      this.progress = 0;  // 重置进度条
      if (!this.file) {
        console.log("请选择文件");
        return;
      }

      // 获取上传 token 和 key
      getUploadTokenAndKey(this.uploaderId, this.subtitled, this.tags)
        .then(response => {
          const token = response.data.data.token;
          const key = response.data.data.vkey;
          console.log("开始startUpload")

          // 接下来的上传逻辑
          this.startUpload(token, key);
          console.log("完成startUpload")
        })
        .catch(error => {
          console.error('获取上传凭证出现错误:', error.message);
          console.error('详细错误:', error)
        });
    },
    startUpload(token, key) {
      console.log("startUpload方法被调用了");
      // 定义七牛云上传配置
      const config = {
        // useCdnDomain: true,
        // region: qiniu.region.z2
      };

      const putExtra = {
        // params: {},
        // mimeType: "video/mp4",
      };


      // 文件上传
      const observable = qiniu.upload(this.file, key, token, putExtra, config); // 确保使用正确的key
      const observer = {
        next: (res) => {
          this.progress = res.total.percent.toFixed(2);
          console.log("上传进度:", this.progress);
        },
        error(err) {
          console.error('上传错误: ', err);
        },
        complete(res) {
          console.log('上传完成: ', res);
          const vid = parseInt(localStorage.getItem('vid'))
          callBackUpload(vid).then(response=>{
            console.log("服务器确认上传成功！",response)
          })
            .catch(error=>{
              console.log("服务器确认上传失败！",error)
            })
        }
      };

      // 上传开始x
      const subscription = observable.subscribe(observer);

    }
  },


};
</script>

<style scoped lang="scss">

 //上传进度条
.progress-container {
  height: 20px;
  width: 100%;
  background-color: #e0e0e0;
  margin-top: 15px;
  border-radius: 10px;
}

.progress-bar {
  height: 100%;
  background-color: #4caf50;
  transition: width 0.4s;
}
.page {
  font-family: 'Arial', sans-serif;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.top-banner {
  background: #222;
  padding: 20px 40px;
  color: #fff;
  border-radius: 10px;
  margin-top: 50px;
  width: 60%;

  .title {
    font-size: 1.8em;
    font-weight: bold;
    margin-bottom: 10px;
  }

  .subtitle {
    font-size: 1.3em;
  }
}

.upload-section {
  margin-top: 50px;
  width: 60%;

  .input-section {
    display: flex;
    flex-direction: column;
    gap: 10px;
    margin-bottom: 20px;

    input {
      padding: 8px 12px;
      border: 1px solid #ccc;
      border-radius: 5px;
    }

    .upload-box {
      display: flex;
      justify-content: space-between;
      align-items: center;
      border: 1px solid #ccc;
      padding: 15px;
      border-radius: 8px;

      input[type="file"] {
        display: none;
      }

      .upload-label {
        background-color: #333;
        color: #fff;
        padding: 10px 15px;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.2s;

        &:hover {
          background-color: #444;
        }
      }

      .upload-btn {
        background-color: #ff7f00;
        color: #fff;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.2s;

        &:hover {
          background-color: #ff9f30;
        }
      }
    }

    .file-hint {
      margin-top: 20px;
      color: red;
    }
  }
}
</style>
