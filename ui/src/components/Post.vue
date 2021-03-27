<template>
  <v-card class="my-6">
    <v-row justify="center">
      <v-col cols="1" class="ml-5 text-center">
        <v-btn icon @click="upvote" :disabled="loading">
          <v-icon :size="30">mdi-arrow-up-drop-circle{{ upvoted ? '' : '-outline' }}</v-icon>
        </v-btn>
        <h3>{{ rating }}</h3>
        <v-btn icon @click="downvote" :disabled="loading">
          <v-icon :size="30">mdi-arrow-down-drop-circle{{ downvoted ? '' : '-outline' }}</v-icon>
        </v-btn>
      </v-col>
      <v-col>
        <v-card-title>{{ post.title }}</v-card-title>
        <v-card-text>{{ post.content }}</v-card-text>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
export default {
  name: 'Post',
  props: {
    post: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      loading: false,
      upvoted: false,
      downvoted: false,
      // TODO: remove later
      rating: 0,
    };
  },
  methods: {
    upvote() {
      this.upvoted = !this.upvoted;
      this.downvoted = false;
      this.loading = true;
      // TODO: Send to server rating +1
      setTimeout(() => {
        if (this.downvoted) {
          this.rating += 2;
        } else if (this.upvoted) {
          this.rating -= 1;
        } else {
          this.rating += 1;
        }
        this.loading = false;
      }, 500);
    },
    downvote() {
      this.downvoted = !this.downvoted;
      this.upvoted = false;
      this.loading = true;
      // TODO: Send to server rating -1
      setTimeout(() => {
        this.loading = false;
      }, 500);
    },
  },
};
</script>
