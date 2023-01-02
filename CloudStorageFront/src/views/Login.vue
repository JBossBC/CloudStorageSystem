<template>
<div id="root">
   <div id="mainPanel" class="container">
     <div id="LoginForm" >
       <img src="../assets/logo.svg" class="m-2 pb-3" width="40" height="52">
       <label class="text-white m-2" style="font-size: 35px">登录</label>
         <div class="input-group mb-3">
           <label class="input-group-text">账号<span style="color: red">*</span></label>
         <input type="text" class="form-control" v-model="username" placeholder="请输入账号">
         </div>
         <div class="input-group mb-3">
           <label class="input-group-text">密码<span style="color: red">*</span></label>
         <input type="password" class="form-control" v-model="password" placeholder="请输入密码">
         </div>
<!--         <div class="btn-group mb-3">-->
         <button type="submit" class="w-100 btn btn-primary btn-lg"  @click="loginEvent">登录</button>
<!--         </div>-->
     </div>
   </div>
  <div class="container w-100 text-center  align-items-center" id="footDisplay">
    <footer class="w-100 d-flex flex-wrap justify-content-center align-items-center  py-3 my-4 border-top">
        <div class="right col-md-4   align-items-center">
           <img src="../assets/logo.svg " width="40" height="32" class="mb-3 me-2 mb-md-0 text-muted text-decoration-none lh-1">
          <span class="mb-3 mb-md-0 text-muted">© 2022 Xiyang</span>
        </div>
    </footer>
  </div>
</div>
</template>
<script>
import RouteUtil from "@/router/RouteUtil";
export default {
  name: "Login",
  data(){
    return {
      username: "",
      password: "",
    }
  },
  methods:{
    loginEvent:function(){
      if(this.username==undefined||this.password==undefined){
        alert("please input your username and passowrd");
        return;
      }
      RouteUtil.setCookie("isLogin",this.username,60);
      this.$message({
        message:"登录成功",
        type:"success"
      });
      this.$store.state.loginName=this.username;
      // console.log(this.$store.userName);
      // this.$root.$data.userName=this.$props.username;
      this.$router.push("/home");
    }
  }
}
</script>

<!--when logoff return the log page the background color become the white-->
<style>
#root{
  width: 100%;
  height: 100%;
  border: 0px;
  margin: 0px;
  background-color: RGBA(33, 37, 41, var(--bs-bg-opacity, 1)) !important;
}
 #LoginForm{
   position:absolute;
  top: 100px;
  right:300px;
   width:400px;
   heigth:500px;
}
 #mainPanel{
    width: 100%;
   height: 100%;
 }
 #footDisplay{
   position:absolute;
   bottom: 0px;
 }
</style>