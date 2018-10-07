export default {
	state:{
		name:'梁菲',
		id:'CN3756780',
		image:'',
		token:'',
		isValid:false,
	},
	getters:{
		Login:(state)=>{
			return state
		},
		LoginIsValid:(state)=>{
			return state.isValid
		}
	}
}
