<template>
<div id='OrderShow'>
	订单号码:
	<el-input v-model="input" placeholder=""></el-input>
	发货商店:<el-input v-model="input2" placeholder=""></el-input>
	订单日期:
	<el-date-picker
      v-model="value"
      type="date"
      placeholder="选择日期"
      align="right"
      :picker-options="pickerOptions">
    </el-date-picker>
    销售渠道:
    <el-input v-model="input3" placeholder=""></el-input>
    订单类型:
    <el-input v-model="input4" placeholder=""></el-input>
    <div>
    	<el-table :data="tableData" class="table" stripe border>
          <el-table-column prop="number" :label="table.number">
          </el-table-column>
          <el-table-column prop="category" :label="table.category">
          </el-table-column>
          <el-table-column prop="name" :label="table.name">
          </el-table-column>
          <el-table-column prop="count" :label="table.count">
          </el-table-column>
          <el-table-column :label="table.price">
          	<template slot-scope="scope">
          		<span>{{(scope.row.price).toFixed(2)}}</span>
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
    	<el-input v-model="input" placeholder=""></el-input>
    	折扣:
    	<el-input v-model="input" placeholder=""></el-input>
    	捐赠金额:
    	<el-input v-model="input" placeholder=""></el-input>
    	配送费:
    	<el-input v-model="input" placeholder=""></el-input>
    	已付金额:
    	<el-input v-model="input" placeholder=""></el-input>
    </div>
</div>
</template>
<script>
	
</script>
<style>
	
</style>