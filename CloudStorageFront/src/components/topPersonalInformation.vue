
<template>
  <header class="p-3 text-bg-dark">
  <div class="container">
    <div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
      <a class="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-dark text-decoration-none">
        <img src="../assets/logo.svg" width="40" height="32">
        <span class="fs-4  text-white" >CloudStorage</span>
      </a>
      <template v-if="loginName">
      <div class="text-start">
        <a class="btn btn-info m-2"  style="width: 100px" @click="selectPage(3)" >我的主页</a>
      </div>
      <div class="text-muted">
        <a class="btn btn-info m-2"   style="width: 100px" @click="selectPage(1)" >上传文件</a>
      </div>
      <div class="text-muted">
        <a class="btn btn-info m-2"  style="width: 100px" @click="selectPage(2)"  >下载文件</a>
      </div>
      </template>
      <div class="text-end dropdown" >
    <a class="btn btn-warning m-2 " style="width: 100px;overflow: hidden"  @click='loginCheck'>{{loginNameDisplay}}</a>
        <div id="userOperationList " class="dropdown-content" v-if="userOperationDisplay">
        <ul  class="dropdown-menu dropdown-menu-dark position-static d-grid gap-1 p-2 rounded-3 mx-0 shadow dropdown-content">
          <li><a class="dropdown-item rounded-2 " @click="PersonalInfoDialog" style="text-decoration: none;width: 100px">个人信息</a></li>
          <li ><a class="dropdown-item rounded-2 " style="text-decoration: none;width: 100px"  @click.once="logoff">注销</a></li>
        </ul>
        </div>
      </div>
    </div>
  </div>
  </header>
  <!--个人信息模态框(尝试使用element-ui)-->
  <el-dialog :model-value="personInfoDisplay" @close="personInfoDisplay=false" title="个人信息" draggable>
    <div>
    <div class="input-group mb-3"  >
         <label class="input-group-text" >姓名:</label>
         <input class="form-control" v-model="personalInfo.name">
    </div>
      <div class="input-group mb-3">
        <label class="input-group-text " style="max-height: 30px">邮箱:</label>
   <el-select v-model="personalInfo.sex" >
<!--     <template #prefix>-->

<!--     </template>-->
          <el-option value="0">女</el-option>
          <el-option value="1">男</el-option>
        </el-select>
      </div>
      <div class="input-group mb-3">
        <label class="input-group-text">电话号码:</label>
        <input type="tel" class="form-control" v-model="personalInfo.phoneNumber">
      </div>
      <div class="input-group mb-3">
        <label class="input-group-text">邮箱:</label>
        <input type="email" class="form-control" v-model="personalInfo.Email">
      </div>
    </div>
    <template #footer>
       <button class="btn btn-prev" @click="personInfoDisplay=false" >取消</button>
      <button class="btn-primary btn" @click="updatePersonalInfo">保存</button>
    </template>
  </el-dialog>
</template>
<style>
.dropdown-content{
  display: none;
  position: absolute;
  min-width: 100px;
  height: 100px;

  /*width: 10px;*/
}
.dropdown{
  position:relative;
  display: inline-block;
}
.dropdown:hover .dropdown-content{
  display: block;
}

</style>


<script >
// import "element-ui/lib/tree.js";
import routeUtil from "@/router/RouteUtil";
import {ref} from "vue";

// let personInfoDisplay=ref(false)
export default{
  name:"topPersonalInformation",
  //TODO flush page can reset the loginName
  props:{
    personalInfo:{
      type:Object,
      default:()=>{
        let defaultInstance={}
        defaultInstance.name="";
            defaultInstance.phoneNumber="";
            defaultInstance.sex="";
         defaultInstance.Email="";
         this.personalInfo=defaultInstance;
      }
    },
    loginName: {
      type: String,
      required: true
    },
    pageNumber:{
      type: Number,
      required: true
    },
  },
  methods:{
    loginCheck:function(){
      if(!this.$store.state.loginName){
        window.location.href="/";
      }
    },
    selectPage:function(pageNumber){
      this.$emit('selectDisplayPage',pageNumber);
    },
    logoff:function(){
      this.$store.state.loginName="";
      routeUtil.deleteCookie("isLogin",this.$store.state.loginName);
      this.loginCheck();
    },
    PersonalInfoDialog:function (){
      this.personInfoDisplay=true;
    },
    //save the user info and exit the dialog
    updatePersonalInfo:async function(){
      let This=this;
       await this.$axios.get("/user/save",{params:JSON.stringify(this.personalInfo)}).then(function(response){
         let responseData=response.data;
         if(responseData.result!="result"){
                if(!responseData.message){
                  responseData.message="保存失败";
                }
                This.$message({
                  type:"error",
                  message:responseData.message
                })
           return;
         }
         This.$message({
           type:"success",
           message:"保存成功"
         })
         //exit the dialog
         this.personInfoDisplay=false;
       }).catch(function(reason){
          This.$message({
            type:"error",
            message:"服务器错误，请联系管理员"
          })
       })


    }
  },
  data(){
    return{
      userOperationDisplay:false,
      personInfoDisplay:false
    }
  },
  computed:{
    loginNameDisplay:function(){
      if(!this.loginName){
        return "请登录";
      }
      return this.loginName;
    }
  },
  created(){
    if(this.loginName){
      this.userOperationDisplay=true;
    }
  }
}

</script>