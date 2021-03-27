import Vue from 'vue';
import VueRouter from 'vue-router';
import Main from './pages/Main.vue';
import Posts from './pages/Posts.vue';
import Profile from './pages/Profile.vue';
import About from './pages/About.vue';
import Login from './pages/Login.vue';
import SignUp from './pages/SignUp.vue';

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'main',
      component: Main,
      children: [
        {
          path: '/',
          name: 'posts',
          component: Posts,
        },
        {
          path: '/login/',
          name: 'login',
          component: Login,
        },
        {
          path: '/signup/',
          name: 'signup',
          component: SignUp,
        },
        {
          path: '/profile/',
          name: 'profile',
          component: Profile,
        },
        {
          path: '/about/',
          name: 'about',
          component: About,
        },
      ],
    },
  ],
});

export default router;
