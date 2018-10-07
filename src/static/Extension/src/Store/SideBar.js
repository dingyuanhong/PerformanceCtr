export default {
	state:{
		options: [
			{name:'个人',icon:'el-icon-news',route:'/Self',index:'1'},
			{
				name:'订单',icon:'el-icon-menu',route:'/Order',index:'2',
				child:[
				{name:'入档',icon:'el-icon-edit-outline',route:'/OrderNew',index:'2-1'},
				{name:'查看',icon:'el-icon-document',route:'/OrderManage',index:'2-2'},
				]
			},
			{name:'汇总',icon:'el-icon-tickets',route:'/Summary',index:'3'},
		],
		position:'left',
		active:'',
	},
	getters:{
		Sidebar:(state)=>{
			return state
		},
		GetRoute:(state, getters) => (index) =>{
			let optionSelect = null
			let options = state.options
			for(var i = 0;i < options.length;i++){
				let option = options[i]
				if(option.child && option.child.length !== 0){
					let childs = option.child
					for(var j=0;j<childs.length;j++){
						let child = childs[j]
						if(child['index'] === index){
							optionSelect = child
							break
						}
					}
					if(optionSelect !== null){
						break
					}
				}else if(option['index'] === index){
					optionSelect = option
					break
				}
			}
			return optionSelect
		},
		GetIndex:(state, getters) => (route) =>{
			let optionSelect = null
			let options = state.options
			for(var i = 0;i < options.length;i++){
				let option = options[i]
				if(option.child && option.child.length !== 0){
					let childs = option.child
					for(var j=0;j<childs.length;j++){
						let child = childs[j]
						if(child['route'] === route){
							optionSelect = child
							break
						}
					}
					if(optionSelect !== null){
						break
					}
				}else if(option['route'] === route){
					optionSelect = option
					break
				}
			}
			return optionSelect
		}
	}
}