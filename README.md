# GoCL_API

* 基于 Gin 的 API 框架
* go 版本要求 `>= 1.6`
* 目录结构
    ```
    ├── config                  //配置文件
    ├── controller              //业务逻辑层
    ├── doc                     //文档
    ├── library                 //拓展库
    │   ├── common
    │   ├── error
    │   └── system
    ├── middleware              //中间件
    ├── model                   //数据模型层
    ├── public                  //静态资源
    │   ├── css
    │   ├── image
    │   └── js
    ├── router                  //路由
    └── vendor                  //第三方依赖
    ```

## 快速开始

* 获取源码 (将代码放置于Go环境 src 目录下)

    `git clone git@github.com:GoCL/api.git`

* 入口文件

    `go run gocl.go`