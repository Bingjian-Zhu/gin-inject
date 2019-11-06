# gin-inject
gin，以依赖注入方式把项目解耦。[相关教程]()

### 项目结构
<pre><code>
├── cmd  程序入口
├── common 通用模块代码
├── config 配置文件
├── controller API控制器
├── docs 数据库文件
├── models 数据表实体
├── page 页面数据返回实体
├── repository 数据访问层
├── router 路由
├── service 业务逻辑层
├── vue-admin Vue前端页面代码
</code></pre>

### 下载安装项目
`go get -x github.com/bingjian-zhu/gin-inject`

### 运行方式
1. 修改config中数据库配置
2. 执行`go run main.go`


> 此demo仅为介绍gin中如何使用依赖注入，具体实战项目请参考[gin-vue-admin](https://github.com/Bingjian-Zhu/gin-vue-admin)
