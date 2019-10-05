# golang-web-template
## 说明
这个项目是一个golang的web开发模板，想着当你在用golang进行web开发的时候可以直接将项目拷贝然后作为脚手架，减少前期构建项目的工作。想法来源于之前写过的 spring-boot-template  

使用golang进行web开发现在当前还没有非常成熟的一站式解决方案，没有类似java中的spring，所以项目的构建也是见仁见智的，这里我只是根据之前java开发的经验来给到我认为合适的样子，适合小项目快速构建。不适用于巨大项目的构建。请斟酌参考。  

PS： 如果其中有写的不好的地方，或者有可以优化的地方也请点击上方直接提 issue 拜托🙏

## 使用的三方库
- gin web框架
- xorm 数据库
- viper 读取配置文件
- zap 日志
- validator.v9 验证器
- jinzhu/copier 数据映射拷贝
- pkg/errors 错误处理
- swagger 接口文档

使用时需要设置项目 GOPATH 为当前项目根目录，当前依赖使用govendor，后续等13普及会引入go mod

## 项目说明
### 目录说明
```
.
├── README.md
├── doc
│   ├── default-conf.yml 默认配置
│   ├── golang-web-template.postman_collection.json postman测试文件
│   └── golang_web_template.sql 数据库sql
└── src
    ├── base
    │   ├── config 配置文件
    │   ├── db 数据库
    │   ├── httper 请求统一定义处理器
    │   ├── logger 日志
    │   ├── router 路由
    │   └── validator 验证器
    ├── cmd 启动项
    │   ├── api 接口文档依赖
    │   ├── docs 接口文档文件
    │   └── main.go 启动文件
    ├── controller 
    ├── dao
    ├── model
    ├── myerr 自定义错误
    ├── service
    ├── util 工具类
    └── val 验证器
```

### 配置文件
- 项目启动之后会默认去读取doc目录下的配置文件 default-conf.yml 根据配置文件中的配置进行启动
- doc中还包含：数据库的sql文件，postman测试文件

### 日志说明
- 当前的日志使用zap日志框架，当前配置为在启动的根目录创建log文件夹并记录日志
- debug级别的日志不记录文件，info日志单独记录文件，error与warn日志单独记录文件
- 可以根据自己的需要进行调整或者切换，切换日志框架只需要实现具体接口和方法就可以了
- 打印日志方式在main.go中有案例

### 错误处理
- 错误处理使用`pkg/errors`进行封装，或者进行自定义错误在myerr中
- 错误统一交由一个地方进行处理和日志打印
- 当项目在进行过程中，如果错误需要统一进行处理的，那么就会自定义错误
- 在调用链比较长的业务中或者一些三方网络调用中需要主动手动去记录日志，精确定位

### 接口文档
- 使用swagger来自动生成接口文档 
https://github.com/go-swagger/go-swagger/ 
https://github.com/swaggo/swag
- 使用之前需要安装 go get -u github.com/swaggo/swag/cmd/swag 相关的库，查看上面两个库中的说明进行安装即可
- 在根目录使用命令 `swag init -g ./cmd/main.go -o ./cmd/docs` 会生成对应的接口文档json,yml到对应位置
- 在根目录使用命令 `swagger serve ./cmd/docs/swagger.yaml` 会在本地启动接口文档的服务进行查看
- 生成文档的时候会生成一个docs.go文件，在启动的main.go上面已经有了这个文件的依赖，所以直接启动项目访问：http://localhost:8080/swagger/index.html   也可以看到接口文档，并且可以直接使用其中的try进行测试
- 生成的命令记录在脚本 gen-swagger-doc.sh 中，也可以直接运行脚本查看接口文档
