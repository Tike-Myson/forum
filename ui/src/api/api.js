import Axios from 'axios';

const api = Axios.create({
  timeout: 30000,
});

api.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded';

export function fetchPosts() {
  return api
    .get('http://localhost:8080/api/posts')
    .then((response) => response.data)
    .catch((err) => {
      throw new Error(err);
    });
}

export function createPost(postData) {
  return api
    .post('http://localhost:8080/api/create-post', postData)
    .then((response) => response.data)
    .catch((err) => {
      throw new Error(err);
    });
}

export function addUser(userData) {
  return api
    .post('http://localhost:8080/api/auth/signup', userData)
    .then((response) => response.data)
    .catch((err) => {
      throw new Error(err);
    });
}
