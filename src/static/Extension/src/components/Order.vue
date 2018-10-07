<template>
<div>
	<div>订单信息</div>
	<table>
		<tr>
			<th v-for="(name,index) in order_columes" :key="index">{{name}}
			</th>
		</tr>
		<tr v-for="(order,index) in orders" :key="index">
			<td>{{order.id}}</td>
			<td>{{order.type}}</td>
			<td>{{order.name}}</td>
			<td>{{order.count}}</td>
			<td>{{order.price}}</td>
			<td>{{order.total}}</td>
		</tr>
	</table>
</div>
</template>
<script>
export default {
  name: 'Order',
  data () {
    return {
    }
  },
  computed: {
  	order_columes:()=>{
  		return ['产品编号','类别','产品名称','订货数量','销售价','小计']
  	},
  	orders:()=>{
  		return [{id:'1307',type:'P',name:'胶囊',count:3,price:365.00,total:1098.00}]
  	}
  }
}
</script>
<style scoped>
</style>
