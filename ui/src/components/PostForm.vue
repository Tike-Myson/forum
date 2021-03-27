<template>
  <v-card justify-center align-center>
    <v-form @submit.prevent="emitPostData">
      <v-card-title>Create Post</v-card-title>
      <v-row class="px-5">
        <v-col cols="5">
          <v-text-field v-model="postData.title" label="Title" outlined></v-text-field>
          <v-textarea v-model="postData.content" label="Content" outlined></v-textarea>
        </v-col>
      </v-row>
      <v-card-actions>
        <v-btn type="submit">
          Add post
        </v-btn>
      </v-card-actions>
    </v-form>
  </v-card>
</template>

<script>
export default {
  name: 'PostForm',
  data() {
    return {
      postData: {
        title: '',
        content: '',
      },
    };
  },
  methods: {
    emitPostData() {
      this.$emit('updatePost', { ...this.postData });
      this.postData.title = '';
      this.postData.content = '';
    },
  },
};
</script>
