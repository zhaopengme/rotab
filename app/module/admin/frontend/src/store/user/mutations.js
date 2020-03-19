export default {
  setUser(state, data) {
    state.token = data.token;
    state.username = data.username;
    localStorage.setItem("token", data.token);
    localStorage.setItem("username", data.token);
  },
  setUserinfo(state, data) {
    state.nickname = data.nickname;
    state.email = data.email;
    state.description = data.description;
    localStorage.setItem("nickname", data.nickname);
    localStorage.setItem("email", data.email);
    localStorage.setItem("description", data.description);
  }
};
