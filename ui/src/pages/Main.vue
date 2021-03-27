<template>
  <v-app class="main-app">
    <v-navigation-drawer absolute color="#E8E8E8" expand-on-hover clipped fixed app permanent>
      <v-list>
        <div v-for="(navItem, i) in navItems" :key="i">
          <v-list-item v-if="authorized ? true : !navItem.authOnly" :to="navItem.to" router exact>
            <v-list-item-action>
              <v-icon>{{ navItem.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title v-text="navItem.title" />
            </v-list-item-content>
          </v-list-item>
        </div>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar color="#495464" class="px-10" clipped-left fixed app>
      <v-img :src="mainIcon" max-width="200" position="bottom -20px left 10px"></v-img>
      <v-spacer />
      <v-btn to="/login" outlined dark>
        <v-icon color="#f4f4f2" left>mdi-login</v-icon>
        {{ authorized ? 'Sign out' : 'Sign in' }}
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import mainIcon from '@/assets/srak_logo.png';

export default {
  data() {
    return {
      mainIcon,
      navItems: [
        {
          icon: 'mdi-format-list-bulleted-square',
          title: 'Posts',
          to: '/',
          authOnly: false,
        },
        {
          icon: 'mdi-account',
          title: 'Profile',
          to: '/profile',
          authOnly: true,
        },
        {
          icon: 'mdi-information-outline',
          title: 'About',
          to: '/about',
          authOnly: false,
        },
      ],
    };
  },
  computed: {
    authorized() {
      return this.$store.getters.authState;
    },
  },
};
</script>

<style lang="scss" scoped>
.main-app {
  background: #f4f4f2;
}
</style>
