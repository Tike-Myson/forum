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
      <v-form @submit.prevent="signIn" v-model="valid">
        <v-text-field
          v-model="username"
          :rules="nonemptyRules"
          type="text"
          :prepend-inner-icon="'mdi-account'"
          label="Username"
          filled
        ></v-text-field>
        <v-text-field
          v-model="password"
          :rules="nonemptyRules"
          type="password"
          :prepend-inner-icon="'mdi-lock'"
          label="Password"
          filled
        ></v-text-field>
        <v-card-actions class="d-flex justify-center flex-column">
          <v-btn type="submit" color="success" :disabled="!valid" large rounded width="200"
            >login</v-btn
          >
        </v-card-actions>
      </v-form>
    </v-card>
    <div class="mt-5">
      New to srakoverflow?
      <v-btn :to="'/signup'" class="ml-5" text outlined small>Create an account</v-btn>
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
      valid: false,
      nonemptyRules: [(v) => !!v || 'This field can not be empty'],
    };
  },
  methods: {
    signIn() {
      console.log('Form submitted');
    },
  },
};
</script>
