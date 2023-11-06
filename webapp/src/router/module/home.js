import Home from '@/views/Home.vue';

const home =[
  {
    path: '/',
    name: 'Home',
    component: Home,
    children: [
      {
        path: 'videos',
        name: 'videos',
        component: () => import('../../common/videoList/Videos.vue'),
      },

    ],
  },

]
