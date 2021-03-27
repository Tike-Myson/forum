<template>
  <v-card>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-text>{{ content }}</v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'Post',
  props: {
    title: {
      type: String,
      required: true,
    },
    content: {
      type: String,
      required: true,
    },
  },
  data() {
    return {};
  },
};
</script>
