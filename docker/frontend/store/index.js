export const actions = {
  onAuthStateChangedAction ({ commit }, { authUser, claims }) {
    if (!authUser) {
      if (location.pathname !== '/signin') {
        location.href = '/signin'
      }
    }
    commit('ON_AUTH_STATE_CHANGED_MUTATION', { authUser, claims })
  }
}

export const mutations = {
  ON_AUTH_STATE_CHANGED_MUTATION: (state, { authUser, claims }) => {
    if (!authUser) {
      state.user = null
    } else {
      const { uid, email, emailVerified } = authUser
      state.user = { uid, email, emailVerified }
    }
  }
}
