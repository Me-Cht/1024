import {createRouter, createWebHistory} from 'vue-router'
import Upload from "../common/upload/Upload.vue";
import Home from "../views/Home.vue";
import Videos from "../common/videoList/Videos.vue";
import AboutView from "@/views/AboutView.vue";
import LoginModule from "@/common/user/LoginModule.vue";
import VideoSearchResult from "@/common/search/VideoSearchResult.vue";
import UserInfo from "@/common/user/UserInfo.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView
    },
    {
      path: '/videos',
      name: 'videos',
      component: Videos
    },
    {
      path: '/test',
      name: 'test',
      component: Videos
    },
    {
      path: '/upload',
      name: 'upload',
      component: Upload
    },
    {
      path:'/login',
      name:'login',
      component:LoginModule,
    },
    {
      path:'/search',
      name:'search',
      component:VideoSearchResult,
    },
    {
      path:'/userinfo',
      name:'userinfo',
      component:UserInfo,
    }
  ]
})

export default router
