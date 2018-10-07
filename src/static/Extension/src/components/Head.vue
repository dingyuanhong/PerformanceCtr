<template>
  <div id="Head">
    <div class='logo' :style="logo_image"></div>
    <div class='head'>
      <el-badge :value="message_count" :max="99" class="item">
        <el-button size="small" icon='el-icon-message' circle @click="btn_message"></el-button>
      </el-badge>
    </div>
    <div class="self">
      <img v-if="login.image !== ''" :src="login.image" class="image float"/>
      <img v-else :src="this.Resource.user" class="image float"/>
      <div>
        <a class='name' @click="btn_Self">{{login.name}}</a>
        <div class="id" >{{login.id}}</div>
      </div>
      <div>
        <el-button v-on:click="btn_logout"><i class="logout" :style="logout_image"></i></el-button>
      </div>
    </div>
  </div>
</template>
<script>
import {mapGetters} from 'vuex'
export default {
  name: 'Head',
  data() {
    return {
      message_count:0
    }
  },
  computed: {
    ...mapGetters(['Login','Resource','Sidebar']),
    login(){
      return this.Login
    },
    user_image(){
      return this.Resource.user
    },
    logo_image(){
        return {
          backgroundImage: 'url('+ this.Resource.logo + ')'
        }
    },
    logout_image(){
        return {
          backgroundImage: 'url('+ this.Resource.logout + ')'
        }
    }
  },
  methods:{
    btn_logout:function(){
      console.log('logout')
    },
    btn_message:function(){
      console.log('message')
    },
    btn_Self:function(){
      this.Sidebar.active = '0'
    }
  }
}
</script>
<style lang="less">
#Head{
  height:50px;
  position:relation;
  box-sizing: border-box;
  padding-left:100px;

  .logo{
    position:absolute;
    left:0;

    height:100%;
    width:100px;
    background-repeat: no-repeat;
    background-size: 40px 40px;
    background-position: center;
  }

  .head{
    padding-top:12px;
    padding-right:160px;
    text-align:right;
  }

  .self {
    position:absolute;
    top:0;
    right:0;
    width:100px;
    padding-left:40px;

    .float {
      position:absolute;
      left:0;
      top:5px;
    }

    > div {
      display:inline-block;
      vertical-align: middle;

      min-width:10px;
      height:40px;

      margin-left: 5px;
      padding-top: 5px;
    }

    a{
      text-decoration: none;
    }
    a:hover{
      text-decoration: underline;
      cursor:pointer;
    }
  }

  .image{
    width:40px;
    height:40px;
  }
  .name {
    color:#888888;

    line-height:30px;
    font-size: 22px;

    height:30px;
    width:70px;
  }
  .id {
    color:#888888;

    font-size:8px;
    line-height:8px;

    height:10px;
    width:70px;
  }

  .logout{
    width:10px;

    background-repeat: no-repeat;
    background-size: 10px 40px;
    background-position: center;
  }
}
</style>
<style>
#Head .self .el-button {
  padding:0;
  width:10px;
  height:40px;
}

#Head .head .el-badge__content.is-fixed{
  top:0px;
  right:20px;
}
</style>
