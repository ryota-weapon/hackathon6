export default {
    state: {
        userId: null,
    },
    mutations: {
        setAuthenticated(state, id) {
            state.userId = id
        }
    },
    getters: {
        isAuthenticated(state) {
            return state.userId != null
        }
    },
}