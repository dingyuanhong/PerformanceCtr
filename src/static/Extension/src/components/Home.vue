<template>
  <div class="home">
    <app-head class='head'></app-head>
    <div class="content" :style="padding">
        <div :class="barClass">
          <side-bar v-bind:position='position'></side-bar>
        </div>
        <div class="view">
          <router-view/>
        </div>
    </div>
    <app-foot></app-foot>
  </div>
</template>

<script>
import AppHead from '../components/Head.vue'
import AppFoot from '../components/Foot.vue'
import SideBar from '../components/Sidebar.vue'
import {mapGetters} from 'vuex'
export default {
  name: 'Home',
  data(){
    return {
      position:'left',
      padding:{'padding':'0 0 0 64px'},
      barClass:'left',
    }
  },
  computed:{
    ...mapGetters(['Sidebar','GetIndex'])
  },
  mounted:function(){
    this.WatchResize()
  },
  created:function(){
    let route = '/Self'
    let option = this.GetIndex(route)
    if(option !== null){
      this.Sidebar.active = option['index']
    }
    this.$router.push(route)
  },
  components: {
    AppHead,
    SideBar,
    AppFoot
  },
  methods:{
    WatchResize:function(){
      let Height = `${document.documentElement.clientHeight}`
      let Width = `${document.documentElement.clientWidth}`
      this.Resize(Width,Height)

      const that = this
      window.onresize = () => {
        let Height = `${document.documentElement.clientHeight}`
        let Width = `${document.documentElement.clientWidth}`
        that.Resize(Width,Height)
      }
    },
    Resize:function(Width,Height){
      let width = parseInt(Width)
      let height = parseInt(Height)
      if(width >= height){
        if(this.Sidebar.position !== 'left'){
          this.Sidebar.position = 'left'
          this.ChangePosition('left')
        }
      }else{
        if(this.Sidebar.position !== 'bottom'){
          this.Sidebar.position = 'bottom'
          this.ChangePosition('bottom')
        }
      }
    },
    ChangePosition(postion){
      this.position = postion
      this.barClass = postion
      if(postion === 'left'){
        this.padding = {
          'padding':'0 0 0 64px'
        }
      }else{
        this.padding = {
          'padding':'0 0 50px 0'
        }
      }
    }
  }
}
</script>

<style lang="less">
.home {
  height:100%;

  box-sizing: border-box;
  position:relative;
  padding-top:50px;

  > .head{
    position:absolute;
    background-color: #fff;

    border-bottom: 1px solid #DDDDDDDD;
    width:100%;
    top:0;
  }

 > .content{
    height:100%;
    width :100%;
    position:relative;
    box-sizing: border-box;

    > .left{
      position:absolute;
      top :0;
      left:0;
      height:100%;
      width:64px;

      border-right:1px solid #DDDDDDDD;
    }

    > .view {
      width:100%;
      height:100%;
      margin:0;
      padding:0;
    }

    > .bottom{
      position:absolute;
      bottom: 0;

      width:100%;
      height:50px;

      box-sizing: border-box;
      border-top:1px solid #DDDDDDDD;
    }
  }
}
</style>
