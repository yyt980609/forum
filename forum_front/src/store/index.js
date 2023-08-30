import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const defaultLoginResult = {
    accessToken: null,
    userId: null,
    username: null,
}

export default new Vuex.Store({
    state: {
        isLogin: false,
        loginResult: defaultLoginResult,
    },
    mutations: {
        init(state) {
            let loginResult = JSON.parse(localStorage.getItem("loginResult"));
            if (loginResult != null) {
                state.loginResult = loginResult;
            }
        },
        login(state, loginResult) {          // 登录
            state.loginResult = loginResult;
        },
        logout(state) {                      // 退出
            localStorage.removeItem("loginResult");   // 将全局的loginResult删掉
            state.loginResult = defaultLoginResult;
        }
    },
    actions: {},
    getters: {
        isLogin: state => state.loginResult.userId !== null,
        userID: state => state.loginResult.userId,
        username: state => state.loginResult.username,
        accessToken: state => state.loginResult.accessToken,
    }
})
