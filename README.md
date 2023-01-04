# CloudStorageSystem

##  Structure(暂定)

+ language:golang(1.17.6)、java(jdk8)

+ backend:

    nginx、docker(OCI)、kubernetes、go-zero、kafka、go-stash、hadoop、elasticsearch、mysql、redis

+ frontend:

     vue3、element-plus、bootstrap3、axios



## v1.0.0(2023.1.2-预期(2023.1.22))

### 功能 
  + 用户登录，注册(邮件注册)，文件上传、下载、预览，用户信息修改，不同权限登录(管理员、用户)，页面美化(axios前后的进度条，文件目录的树结构，文件下载界面，首页结构和组件需要改善)

   + 文件服务(fastdfs客户端、文件上传、下载、预览、session验证、文件元数据管理(mysql、redis))
   
   + 用户服务(登录、注册、修改、发邮件、不同权限的表设计、session、验证码)go-zero(链路追踪，服务降载、负载均衡、鉴权、服务降级)