<template>
<div id="root">
    <topPersonalInformation @selectDisplayPage="updatePageNumber" :personal-info="personalInfo" :login-name="loginName" :page-number="0"></topPersonalInformation>
  <div class="container d-flex justify-content-center align-items-center  w-100 h-100" v-if="currentPage!=3">
    <div v-show="currentPage==0">
    <h1 style="font-size: 50px">Welcome to the CloudStorage System</h1>
    </div>
    <!--upload page -->
    <div v-show="currentPage==1" style="min-width: 500px">
         <div class="form-floating">
           <input id="fileNameInput" type="text" class="form-control" placeholder="请输入文件名">
           <label for="fileNameInput">存入云盘的位置 <span style="color: red">*</span></label>
         </div>
<!--        <div class="input-group mb-1">-->
<!--          <input type="file" class="form-text btn btn-info" placeholder="请选择要上传的文件">-->
<!--        </div>-->
      <el-upload action=""   ref="uploadRef" :on-exceed="limitFileNumber" :before-remove="fileRemoveHook" v-model:file-list="fileList"  :drag="true"  multiple :limit="3" :auto-upload="false">
        <el-icon  class="el-icon--upload"><upload-filled/></el-icon>
        <template #tip>
        <div class="el-upload__text">
             最多只能上传三个文件
        </div>
        </template>
        <template #trigger>
          <el-button type="primary" class="w-25 btn btn-primary" >选择文件</el-button>
        </template>
        <el-button class="w-100 btn btn-primary" type="success"  @click="uploadFile">上传</el-button>
      </el-upload>
    </div>

    <!--download page -->
    <div v-show="currentPage==2">
      <form>

      </form>
    </div>
  </div>
  <!--info page -->
  <div  class=" d-flex  w-100 h-100"   v-show="currentPage==3">
    <div class="w-25 " style="background-color:darkturquoise " id="FilesList">
      <div class="h6 m-0">
        <label class="btn btn-primary w-100 " style="border-radius: 0px!important;">云盘目录</label>
      </div>
      <el-tree style="background-color:darkturquoise "  :data="filesLocation" :props="defaultProps" @node-click="handleFileNodeClick"></el-tree>
    </div>
  </div>
</div>
</template>

<script>
import {UploadFilled} from "@element-plus/icons-vue";
import topPersonalInformation from "../components/topPersonalInformation.vue";
import RouteUtil from "../router/RouteUtil";

export default {
  name: "Home",
  components: {topPersonalInformation,UploadFilled},
  data(){
    return {
      personalInfo:{
        name:"析样",
        phoneNumber:"18080705675",
        sex:"1",
        Email:"1577002722@qq.com"
      },
      fileList:[],
      loginName: "",
      currentPage:0,
      filesLocation:[{
        label:"xiyang",
        children: [{
          label:"hello",
          children: [{
            label:"helloTwo"
          }]
        },{
          label:"world",
          children: [{
            label: "worldTwo"
          }]
        }]
    },{
        label:"xiyang2"
      }],
      defaultProps:{
        children:'children',
        label:'label'
      }
    }
  },
  methods:{
    updatePageNumber:function(pageNumber){
              this.currentPage=pageNumber;
    },
    limitFileNumber:function(){
      this.$message({
        type:"error",
        message:"最多只能上传三个文件"
      })
    },

    fileRemoveHook:async function(removeFile,FileList){
      let This=this;
      let isContinue=true;
     await This.$confirm(
             '确定要取消删除'+removeFile.name+"文件吗",
             '注意',
             {
               confirmButtonText:'确定',
               cancelButtonText:'取消',
               type:"info",
             }
         ).then(function(){
         This.$msgbox({
               type:"info",
              message:"删除"+removeFile.name+"成功"
            })
         }).catch(function(){
           This.$msgbox({
             type:"info",
             message:"取消成功"
           })
       isContinue=false;
      })
      return isContinue;
    },
    //custom click thing to uploadFile
    uploadFile:function(){
      let This=this
      // this.$refs.uploadRef.submit();
      this.$axios.post("/file/upload",this.fileList).then(function(response){
           let responseData=response.data;
           if(responseData.result!="true"){
             if(!responseData.message){
               responseData.message="服务器错误";
             }
             This.$message({
               type:"error",
               message:"上传失败::"+responseData.message
             })
             return;
           }
           This.$message({
             type:"success",
             message:"上传成功"
           })
        This.fileList.splice(0,This.fileList.length);
      }).catch(function (reason){
         This.$message({
           type:"error",
           message:"网络出现异常::"+reason
         })
      })
    },

    // In life cycle ,Init viewable data to vuex
    initRelativeInfo:function(){
      this.$store.state.loginName=RouteUtil.getCookie("isLogin");
    },

    handleFileNodeClick:function(){

    }
  },
  watch:{
    loginName:function(newVar,oldVar){
       if (newVar&&newVar!=oldVar){
         this.$store.state.loginName=newVar;
       }
    }
  },
  // the created hook function execute after the computed ,the method and the data init completely
  created() {
    this.initRelativeInfo();
    if (!this.loginName){
      if(this.$store.state.loginName){
        this.loginName=this.$store.state.loginName;
      }
    }
  }
}
</script>

<style>
#root{
  width: 100%;
  height: 100%;
  border: 0px;
  margin: 0px;
  background-color: aliceblue !important;
}

<!--the filelist display animation has some problem need to repair-->
#FilesList div[display="none"]{
  opacity: 0;
  transform: translateX(200px);
  transition:3s;
}
#FilesList div[display="block"]{
  opacity:1;
  transform: translateX(0px);
}
</style>