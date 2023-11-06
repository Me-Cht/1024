<template>
  <div v-if="props.showModal" class="modal" @click.stop="modalClick">
  <div class="modal-header">
    <span class="welcome-text">欢迎登录</span>
    <button class="close-btn" @click="closeModal">×</button>
  </div>
  <div class="modal-body">
    <p>欢迎</p>
    <input type="text" placeholder="请输入账号" v-model="loginState.username">
    <input type="password" placeholder="请输入密码" v-model="loginState.password">
    <div class="red-line"></div>

    <button class="login-btn" @click="debouncedLogin">登录</button>

    <div class="actions-row" >
      <button class="register-btn" >没有账号？<a href="#" @click="showRegister">去注册</a></button>
    </div>

  </div>
    <!-- 注册模态框 -->
    <RegisterModule :showRegisterModal="showRegisterModal" @closeRegister="closeBothModules" @switchToLogin="handleSwitchToLogin"></RegisterModule>

</div>
</template>




<script setup>
import {defineProps, defineEmits, onMounted, reactive, ref,watch} from "vue";
import { login, register } from "@/api/user";
import RegisterModule from "@/common/user/RegisterModule.vue";

const showRegisterModal = ref(false)
const props = defineProps({
  showModal: {
    type: Boolean,
    required: false
  },

});

const emit = defineEmits(["close"]);
const isLoggingIn = ref(false);

const loginState = reactive({
  username: "",
  password: ""
});

//登录防抖，防止用户连续点击登录按钮
function debounce(func,delay= 300){
  let timer = null
  return function(...args){
    if(timer) clearTimeout(timer)
    timer = setTimeout(()=>{
      func.apply(this,args)
    },delay)
  }
}
const debouncedLogin = debounce(submitLogin);

function handleSwitchToLogin() {
  showRegisterModal.value = false;  // 关闭注册模态框
}
function showRegister() {
  console.log("showRegister方法被调用了");
  showRegisterModal.value = true;
}
function modalClick(event) {
  //如果点击的是模态框本身，关闭模态框
  if (event.target.classList.contains("modal")) {
    emit("close");
  }
}

function closeModal() {
  console.log("Close button clicked");
  emit("close");
}


function submitLogin() {
  if (isLoggingIn.value) return;  // 如果正在登录，直接返回

  console.log(loginState.username, loginState.password);
  isLoggingIn.value = true;  // 设置标识为正在登录
  login(loginState.username, loginState.password)  // 登录逻辑
    .then(data => {
      console.log('登录成功', data);
      localStorage.setItem("username", loginState.username);
      emit('loggedIn', true);
      closeModal();
      alert("登录成功！");
      // 处理登录成功逻辑

    })
    .catch(error => {
      console.error('登录失败', error);
      // 处理登录失败逻辑
    }).finally(() => {
      //不管是不是登录成功，都标记为登录结束
    isLoggingIn.value = false;  // 重置标识
  });
}
function closeBothModules() {
  showRegisterModal.value = false;
  emit("close");  // 关闭登录框
}

// 监听 ESC 键按下事件，关闭登录框
onMounted(() => {
  window.addEventListener("keydown", (event) => {
    if (event.key === "Escape") {
      closeModal();
    }
  });
});


</script>


<style scoped lang="scss">
.modal {
  width: 400px;
  background-color: #fff;
  padding: 20px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.5);
  border-radius: 8px;
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 1000;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.welcome-text {
  color: #333;
  font-size: 20px;
  position: relative;
  top: 15px;
  left: 40%;
}

.modal-body input {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ccc;
  border-radius: 4px;
  cursor: text;
}

.red-line {
  height: 2px;
  background-color: #f00;
  margin: 10px 0;
}

.login-btn {
  width: 100%;
  padding: 10px;
  background-color: #f66;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: text;
  margin: 10px 0;
}

.actions-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.register-btn {
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  font-size: 12px;
  text-decoration: none;
  color: #f66;
}

.register-link {
  margin: 0;
  text-align: right;
  flex-grow: 1;
  color: gray;
  font-size: 12px;
}

.register-link a {
  font-size: 15px;
  color: #f66;
}

.modal, .modal * {
  cursor: default;
}
</style>
