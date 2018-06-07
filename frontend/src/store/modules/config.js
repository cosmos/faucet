const state = {
  bCamera: false // static camera?
};

const mutations = {
  setCamera(state, value) {
    state.bCamera = value;
  }
};

export default {
  state,
  mutations
};
