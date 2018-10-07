<template>
  <div id="sidebar">
    <div class="menu" :style="position==='left'?'':'width:323px;'">
      <el-menu :mode="position==='left'?'vertical':'horizontal'" :collapse="position==='left'?true:false" :default-active="active_index" @select="menu_select">
        <template v-for="(option,index) in options">
          <el-menu-item v-if="!option.child || option.child.length === 0" :key="index" :index="option.index">
            <i :class="option.icon"></i>
            <span slot="title">{{option.name}}</span>
          </el-menu-item>
          <template v-else>
            <el-submenu :key="index" :index="option.index">
              <template slot="title"><i :class="option.icon"></i>
              <span slot="title">{{option.name}}</span></template>
              <template v-for="(subOption,subIndex) in option.child">
                <el-menu-item :key="subIndex" :index="subOption.index">
                  <i :class="subOption.icon"></i>
                  <span slot="title">{{subOption.name}}</span>
                </el-menu-item>
              </template>
            </el-submenu>
          </template>
        </template>
  		</el-menu>
    </div>
	</div>
</template>
<script>
import {mapGetters} from 'vuex'
export default {
  name: 'SideBar',
  data () {
    return {
    }
  },
  props:['position'],
  computed: {
    ...mapGetters(['Sidebar','GetRoute']),
    options () {
      // return this.$store.state.SideBar.options
      // return this.$store.getters.getOptions
      return this.Sidebar.options
    },
    active_index(){
      return this.Sidebar.active
    }
  },
  methods: {
  	selectPage: function (page) {
      // console.log(page)
  		// this.$router.push('/Order')
      this.$router.replace(page)
  	},
    menu_select:function(index){
      let option = this.GetRoute(index)
      this.Sidebar.active = index
      if(option != null){
        this.selectPage(option.route)
      }
    }
  }
}
</script>
<style lang='less'>
	#sidebar {
    text-align: center;
    height:100%;

    > .menu{
      margin:auto;
    }
	}
</style>
<style>
#sidebar .el-menu--horizontal>.el-menu-item{
  height:48px;
  line-height:48px;
}
#sidebar .el-menu--horizontal>.el-submenu .el-submenu__title{
  height:48px;
  line-height:48px;
}
#sidebar .el-menu{
  border-right:0;
}
</style>