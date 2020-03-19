export default {
  setList(state, data) {
    state.pages = data;
  },
  setData(state, data) {
    state.item = data;
  },
  delete(state, data) {
    state.pages.data.splice(data, 1);
  }
};
