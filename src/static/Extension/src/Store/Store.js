import vue from 'vue'
import vuex from 'vuex'
import Lang from './Lang.js'
import SideBar from './SideBar.js'
import Login from './Login.js'
import Resource from './Resource.js'

vue.use(vuex)

export default new vuex.Store({
	modules: {
		Lang:Lang,
		Login: Login,
		SideBar: SideBar,
		Resource:Resource,
	}
})