<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
// import AppHome from './components/Home.vue'
export default {
  name: 'App',
  components: {
  }
}
</script>

<style lang="less">
html,body {
  margin:0;
  height:100%;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  color: #2c3e50;
  height:100%;
}
</style>
