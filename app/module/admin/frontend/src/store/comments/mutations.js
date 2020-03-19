export default {
  setList(state, data) {
    state.pages = data;
  },
  delete(state, data) {
    state.pages.data.splice(data, 1);
  }
};
