<template>
	<div id="login">
		<div class='login'>
			<div class="content">
				<div class='tips' :style="tipsStyle">{{tips}}</div>
				<el-input v-model="name" class='input' :placeholder="Login.name">
				</el-input>
				<el-input v-model="password" class='input' type="password" :placeholder="Login.password">
				</el-input>
				<div>
				<el-checkbox class="auto" v-model="auto" :label='Login.auto'></el-checkbox>
				</div>
				<el-button v-on:click="btn_login" :label='Login.login'>{{Login.login}}</el-button>
			</div>
			<div class="bottom">
			<router-link to='/Forget' class='link'>{{Login.forget}}</router-link>
			<span class="dotted">|</span>
			<router-link to='/Registe' class='link'>{{Login.registe}}</router-link>
			<span class="dotted">|</span>
			<router-link to='/Opinion' class='link'>{{Login.opinion}}</router-link>
			</div>
		</div>
	</div>
</template>
<script>
import storage from 'good-storage'
import {mapGetters} from 'vuex'
export default {
	name:'Login',
	data(){
		return {
			name:'',
			password:'',
			auto:false,
			tips:'',
			tips_color:'#ff0000',
			key:'CACHE_LOGIN'
		}
	},
	beforeMount:function(){
		let datas = storage.get(this.key, [])
		if(datas.length === 0){
			return
		}
		let data = datas[0]
		if(data.auto){
			this.name = data.name
			this.password = data.password
			this.auto = data.auto
		}
	},
	computed:{
		...mapGetters(['LangConfig']),
		Login(){
			return this.LangConfig.Login
		},
		tipsStyle(){
			return {
				color:this.tips_color
			}
		},
	},
	methods:{
		save(){
			if(this.auto){
				storage.set(this.key, [{
					name:this.name,
					password:this.password,
					auto:this.auto
				}])
			}else{
				storage.set(this.key,[])
			}
		},
		btn_login(){
			if(this.name === ''){
				this.tips = this.Login.error.name
			}else if(this.password === ''){
				this.tips = this.Login.error.password
			}else{
				this.$http.post('/login',{})
				.then(res=>{
					this.save()
					this.$router.push('/Home')
					this.tips = ''
				})
				.catch(error=>{
					this.tips = error.message
					this.$router.push('/Home')
				})
			}
		}
	}
}
</script>
<style lang='less'>
#login{
	padding-top:40px;

	.login{
		width:300px;
		height:240px;
		border:1px solid #eee;

		margin:auto;

		.content{
			width:282px;
			height:200px;

			margin:auto;
			margin-top:10px;

			.tips{
				height: 20px;
				margin-left:4px;
				margin-bottom: 4px;
				font: 12px "lucida Grande",Verdana,"Microsoft YaHei";
			}

			.input{
				margin-bottom: 2px;
			}

			.auto{
				margin-top:4px;
				margin-bottom: 8px;
			}
		}

		.bottom{
			height: 16px;
		    margin-bottom: 8px;
		    margin-right:20px;
		    bottom: 0;
		    text-align: right;
		    font-size: 12px;

			.link{
				color:#225592;
			}

			a{
				text-decoration: none;
			}
			a:hover{
				text-decoration: underline;
			}
		}
		.dotted{
			color: #bfbfbf;
    		margin: 0 5px;
		}
	}
}
</style>
<style>
#login .el-button{
	width:100%;
	font-size:18px;
}

#login .auto .el-checkbox__label{
	font-size:12px;
	padding-left:4px;
}
</style>