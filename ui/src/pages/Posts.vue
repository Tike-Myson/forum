<template>
  <v-container>
    <post-form @updatePost="addPost" />
    <post v-for="post in posts" :key="post.title" :post="post" />
  </v-container>
</template>

<script>
import { fetchPosts, createPost } from '@/api/api';
import PostForm from '@/components/PostForm.vue';
import Post from '@/components/Post.vue';

export default {
  name: 'Posts',
  components: {
    PostForm,
    Post,
  },
  data() {
    return {
      posts: [
        {
          title: 'Some title',
          content: 'Some content',
        },
        {
          title: 'Some title',
          content: 'Some content',
        },
      ],
    };
  },
  methods: {
    addPost(postData) {
      createPost(postData).then(() => {
        this.posts = [...this.posts, postData];
      });
    },
  },
  async created() {
    await fetchPosts().then((allPosts) => {
      this.posts = allPosts || [];
    });
  },
};
</script>
