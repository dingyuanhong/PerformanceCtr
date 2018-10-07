<template>
<div id='OrderNew'>
	订单号码:
	<el-input v-model="input"></el-input>
	发货商店:<el-input v-model="input2"></el-input>
	订单日期:
	<el-date-picker
      v-model="value"
      type="date"
      placeholder="选择日期"
      align="right"
      :picker-options="pickerOptions">
    </el-date-picker>
    销售渠道:
    <el-input v-model="input3"></el-input>
    订单类型:
    <el-input v-model="input4"></el-input>
    <div>
    	<el-table :data="tableData" class="table" stripe border>
          <el-table-column :label="table.number">
          	<template slot-scope="scope">
          		<el-autocomplete
			      class="inline-input"
			      v-model="scope.row.number"
			      :fetch-suggestions="querySearch"
			      value-key='number'
			      placeholder="请输入内容"
			      @select="handleSelect"
			    >
				  <template slot-scope="{ item }">
				    <div class="number">{{ item.number }}</div>
				    <span class="name">{{ item.name }}</span>
				  </template>
			    </el-autocomplete>
          	</template>
          </el-table-column>
          <el-table-column :label="table.category">
          	<template slot-scope="scope">
          	<el-input v-model="scope.row.category"></el-input>
          	</template>
          </el-table-column>
          <el-table-column :label="table.name">
          	<template slot-scope="scope">
          	<el-autocomplete
		      class="inline-input"
		      v-model="scope.row.name"
		      :fetch-suggestions="querySearch"
		      value-key='name'
		      placeholder="请输入内容"
		      @select="SelectProduce(scope)"
		    >
			  <template slot-scope="{ item }">
			    <div class="number">{{ item.number }}</div>
			    <span class="name">{{ item.name }}</span>
			  </template>
		    </el-autocomplete>
          	</template>
          </el-table-column>
          <el-table-column :label="table.count">
          	<template slot-scope="scope">
          	<el-input v-model="scope.row.count" type="number"></el-input>
          	</template>
          </el-table-column>
          <el-table-column :label="table.price">
          	<template slot-scope="scope">
          	<el-input v-model="scope.row.price" type="number"></el-input>
          	</template>
          </el-table-column>
          <el-table-column :label="table.subtotal">
          	<template slot-scope="scope">
          	<span>{{(scope.row.count*scope.row.price).toFixed(2)}}</span>
          	</template>
          </el-table-column>
          <el-table-column :label="table.operations">
            <template slot-scope="scope">
              <el-button size="mini">{{ table.new }}</el-button>
              <el-button size="mini" type="danger">{{ table.delete }}</el-button>
            </template>
          </el-table-column>
        </el-table>
    </div>
    <div>
    	总零售额:
    	<el-input v-model="input"></el-input>
    	折扣:
    	<el-input v-model="input"></el-input>
    	捐赠金额:
    	<el-input v-model="input"></el-input>
    	配送费:
    	<el-input v-model="input"></el-input>
    	已付金额:
    	<el-input v-model="input"></el-input>
    </div>
</div>
</template>
<script>
export default {
	name:'OrderNew',
	data(){
		return {
			pickerOptions: {
		      disabledDate(time) {
		        return time.getTime()>Date.now()
		      }
		  	},
	      	input:'888-55553454',
	      	input2:'3020 网购物流中心',
	      	value:'',
			input3:'网上订单',
			input4:'康宝莱订单',
			tableData:[
			{number:'1307',category:'P',name:'科迪赛胶囊',count:3,price:365.00},
			{number:'',category:'',name:'',count:0,price:0}
			],
			table:{
				number:'产品编号',
				category:'类别',
				name:'产品名称',
				count:'订货数量',
				price:'销售价',
				subtotal:'小计金额',
				operations:'操作',
				new:'新增',
				delete:'删除'
			},
			products: [
			{number:'1307',category:'P',name:'科迪赛胶囊',price:365.00},
			{number:'1324',category:'P',name:'强力大蒜素片',price:218.00}
			]
		}
	},
	computed:{
	},
	methods:{
		querySearch(queryString, cb) {
			var products = this.products
			var results = queryString ? products.filter(this.createFilter(queryString)) : products
			// 调用 callback 返回建议列表的数据
			cb(results)
		},
		createFilter(queryString) {
			return (restaurant) => {
			  return (restaurant.number.toLowerCase().indexOf(queryString.toLowerCase()) === 0)
			}
		},
		SelectProduce(scope,item){
			var thiz = scope
			console.log('SelectProduct')
			console.log(scope)
			console.log(item)
			return function(item){
				return this.handleSelect(item,thiz)
			}
		},
		handleSelect:(item,thiz)=>{
			console.log('handleSelect')
			console.log(item)
			console.log(this)
			console.log(thiz)
		}
	}
}
</script>
<style lang='less'>
#OrderNew {
	.table {
		.number {
			font-size:14px;
		}
		.name {
			font-size:10px;
		}
	}
}
</style>