# CloudStorageSystem

##  Structure(暂定)

+ language:golang(1.17.6)、java(jdk8)

+ backend:

|Environment|Version|
|--|--|
|golang|1.17.6|
|go-zero|1.4.2|
|go-stash|1.0.8|
|elasticsearch|8.5.3|
|filebeat|8.5.3|
|kibana|7.17.0|
|zookeeper|3.8.0|
|kafka|3.3.1|
|nginx|stable version|
|docker|20.10.21|
|docker-compose|2.13.0|
|kubernetes|暂定|
|DTM|暂定|
|jaeger|暂定|
|Prometheus|暂定|

+ frontend:

|Environment|Version|
|--|--|
|vue3||
|element-plus||
|bootstrap||
|axios||
     vue3、element-plus、bootstrap3、axios



## v1.0.0(2023.1.2-预期(2023.1.22))

### 功能 
  + 用户登录，注册(邮件注册)，文件上传、下载、预览，用户信息修改，不同权限登录(管理员、用户)，页面美化(axios前后的进度条，文件目录的树结构，文件下载界面，首页结构和组件需要改善)

  + 文件服务(fastdfs客户端、文件上传、下载、预览、session验证、文件元数据管理(mysql、redis))
   
  + 用户服务(登录、注册、修改、发邮件、不同权限的表设计、session、验证码)go-zero(自动降载、自动熔断、鉴权、数据记录、监控报警、超时控制、链路追踪...)
  + 要求:对于微服务系统，我们使用DTM进行分布式事务管理，同时集成日志管理系统(报警日志、链路追踪日志、服务日志、nginx网关日志)(filebeat、kafka、go-stash、elasticsearch、kibana、jeager、Prometheus)，服务数据方面使用传统的mysql、redis集群，
对于redis来说使用主从控制+哨兵模型来追求并发量，同时对于mysql来说，根据情况使用分库分表使用mysql集群。前期使用docker-compose进行容器化部署，使用nginx网关进行反向大力，后期升级为kubernetes做自动拆箱装箱，横向扩展，负载均衡