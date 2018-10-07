import ZH_CN from '../lang/zh-CN.js'

export default {
	state:{
		'type':'zh-CN',
		'zh-CN':ZH_CN
	},
	mutations:{
		LangType(state,type){
			state['type'] = type
		}
	},
	actions:{
		LangType(context,args){
			context.commit('LangType',args)
		}
	},
	getters:{
		Lang:(state)=>{
			return state
		},
		LangConfig:(state)=>{
			return state[state.type]
		}
	}
}
