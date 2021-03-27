<template>
  <v-container
    :style="{ height: '80vh' }"
    class="flex-column"
    fill-height
    justify-center
    align-center
    fluid
  >
    <v-card width="500px" class="pa-10" raised rounded="20">
      <v-form @submit.prevent="signUp" v-model="valid">
        <v-text-field
          v-model="userData.username"
          :rules="usernameRules"
          type="text"
          :prepend-icon="'mdi-account'"
          label="Username"
        ></v-text-field>
        <v-text-field
          v-model="userData.email"
          :rules="emailRules"
          type="text"
          :prepend-icon="'mdi-email'"
          label="Email"
        ></v-text-field>
        <v-text-field
          v-model="userData.password"
          :rules="passwordRules"
          :type="hidden ? 'password' : 'text'"
          :prepend-icon="'mdi-lock'"
          :append-icon="hidden ? 'mdi-eye-off' : 'mdi-eye'"
          @click:append="hidden = !hidden"
          label="Password"
        ></v-text-field>
        <v-text-field
          v-model="passwordRepeat"
          :rules="repeatPasswordRules"
          :type="hidden ? 'password' : 'text'"
          :prepend-icon="'mdi-lock-check'"
          :append-icon="hidden ? 'mdi-eye-off' : 'mdi-eye'"
          @click:append="hidden = !hidden"
          label="Repeat password"
        ></v-text-field>
        <v-card-actions class="d-flex justify-center flex-column">
          <v-btn type="submit" color="success" :disabled="!valid" large rounded width="200"
            >Sign up</v-btn
          >
        </v-card-actions>
      </v-form>
    </v-card>
    <div class="mt-5">
      Already have an acoount?
      <v-btn :to="'/login'" class="ml-5" text outlined small>Sign in</v-btn>
    </div>
  </v-container>
</template>

<script>
import { addUser } from '@/api/api';

export default {
  data() {
    return {
      userData: {
        username: '',
        email: '',
        password: '',
      },
      usernameRules: [(v) => !!v || 'This field can not be empty'],
      passwordRules: [
        (v) => !!v || 'This field can not be empty',
        (v) => v.length >= 7 || 'Password length can not be less than 7 characters',
      ],
      passwordRepeat: '',
      hidden: true,
      emailRules: [
        (v) => !!v || 'This field can not be empty',
        (v) => {
          const validEmailFormat = /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v);
          return !v || validEmailFormat || 'E-mail must be valid';
        },
      ],
      valid: false,
    };
  },
  methods: {
    signUp() {
      addUser(this.userData).then(() => {
        console.log(this.$store.state);
        this.$store.commit('toggleAuthState');
        console.log(this.$store.state);
        this.$router.push('/');
      });
    },
  },
  computed: {
    repeatPasswordRules() {
      return [
        (v) => !!v || 'This field can not be empty',
        (v) => v === this.userData.password || 'Passwords should match',
      ];
    },
  },
};
</script>

<style lang="scss"></style>
