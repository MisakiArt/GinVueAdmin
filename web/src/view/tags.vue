<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                            
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增tags表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="name字段" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="mid字段" prop="mid" width="120"></el-table-column> 
    
    <el-table-column label="categoryId字段" prop="categoryId" width="120"></el-table-column> 
    
    <el-table-column label="status字段" prop="status" width="120"></el-table-column> 
    
    <el-table-column label="isDelete字段" prop="isDelete" width="120"></el-table-column> 
    
    <el-table-column label="tagFollowCount字段" prop="tagFollowCount" width="120"></el-table-column> 
    
    <el-table-column label="latestTaggedDate字段" prop="latestTaggedDate" width="120"></el-table-column> 
    
    <el-table-column label="isDefault字段" prop="isDefault" width="120">
         <template slot-scope="scope">{{scope.row.isDefault|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="createTime字段" prop="createTime" width="120"></el-table-column> 
    
    <el-table-column label="tag被打的次数" prop="tagCount" width="120"></el-table-column> 
    
    <el-table-column label="updateTime字段" prop="updateTime" width="120"></el-table-column> 
    
    <el-table-column label="菜单层级" prop="level" width="120">
         <template slot-scope="scope">{{scope.row.level|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="父标签id" prop="parentTagId" width="120"></el-table-column> 
    
    <el-table-column label="标签渠道来源" prop="userSourceId" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateTags(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteTags(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="name字段:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="mid字段:"><el-input v-model.number="formData.mid" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="categoryId字段:"><el-input v-model.number="formData.categoryId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="status字段:"><el-input v-model.number="formData.status" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="isDelete字段:"><el-input v-model.number="formData.isDelete" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="tagFollowCount字段:"><el-input v-model.number="formData.tagFollowCount" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="latestTaggedDate字段:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.latestTaggedDate" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="isDefault字段:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isDefault" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="createTime字段:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.createTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="tag被打的次数:"><el-input v-model.number="formData.tagCount" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="updateTime字段:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.updateTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="菜单层级:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.level" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="父标签id:"><el-input v-model.number="formData.parentTagId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="标签渠道来源:"><el-input v-model.number="formData.userSourceId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createTags,
    deleteTags,
    deleteTagsByIds,
    updateTags,
    findTags,
    getTagsList
} from "@/api/tags";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Tags",
  mixins: [infoList],
  data() {
    return {
      listApi: getTagsList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            mid:0,
            categoryId:0,
            status:0,
            isDelete:0,
            tagFollowCount:0,
            latestTaggedDate:new Date(),
            isDefault:false,
            createTime:new Date(),
            tagCount:0,
            updateTime:new Date(),
            level:false,
            parentTagId:0,
            userSourceId:0,
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10             
        if (this.searchInfo.isDefault==""){
          this.searchInfo.isDefault=null
        }           
        if (this.searchInfo.level==""){
          this.searchInfo.level=null
        }        
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteTagsByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateTags(row) {
      const res = await findTags({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.retags;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          name:"",
          mid:0,
          categoryId:0,
          status:0,
          isDelete:0,
          tagFollowCount:0,
          latestTaggedDate:new Date(),
          isDefault:false,
          createTime:new Date(),
          tagCount:0,
          updateTime:new Date(),
          level:false,
          parentTagId:0,
          userSourceId:0,
          
      };
    },
    async deleteTags(row) {
      this.visible = false;
      const res = await deleteTags({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createTags(this.formData);
          break;
        case "update":
          res = await updateTags(this.formData);
          break;
        default:
          res = await createTags(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>