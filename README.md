# goctl 

golang 后台脚手架，beego + jwt + [mongodb + mysql]

## 安装

```
go install  git.gzsunrun.cn/sunruniaas/goctl

```

## 使用

```

./goctl -app demo -dir demo -mod github.com/hashwing/demo -mongo=true

```

`app`: 应用名称，默认:server

`dir`: 代码文件夹，默认: 与app相同

`mod`: go mod，默认：server

`mongo` 初始化mongo 作为数据存储，默认: false

`mysql` 初始化mysql 作为数据存储，默认: false

### 生成目录结构


```
|- cmd/{app-name}/
   |- main.go 函数入口
   |- command/
      |- cmd.go 启动命令定义
      |- run.go 启动函数，实例化各类接口（按照自己需求增删）
|- core/ 各类数据结构（struct）、接口（interface）、常量（constant）定义
   |- config.go 配置定义，采用的是环境变量导入
   |- api.go 定义api manager 结构体、auth 接口
   |- store.go 定义存储接口
|- hack/ 各类脚本，例如Dockerfile、启动动脚本、代码生成脚本
|- pkg/ 包含各类业务实现包
   |- auth/ 认证
      |- jwt.go jwt 认证
   |- store/ 数据存储
   |- config/ 配置读取
   |- server/ 服务器接口实现
      |- apis/ 各个版本api 控制器
         |- base/ 基本controller 定义实现，其他cotroller 都会继承（参考beego 文档）
         |- v1/
            |- type/ 接口数据结构定义
            |- login.go 登录controller（根据具体业务更改实现）
            |- test.go 测试例子（可以按需移除）
      |- routers/ 接口路由 （参考beego 文档）
      |- server.go 启动函数
            

```

### 启动程序

1、默认配置文件 `.env`

```
go run cmd/{app}/main.go

```

2、指定配置文件

```
go run cmd/{app}/main.go --cfg=test.env

```

### 配置文件

参考 `github.com/kelseyhightower/envconfig`

```
SERVER_PORT=8090

```

### 日志

使用 `github.com/sirupsen/logrus`

### 认证

使用了 jwt 认证，实现接口

```golang
type Auth interface {
    //JwtAuthFilter 实现beego 认证中间件，通过 beego.InsertFilter("/api/*", beego.BeforeRouter, apiMgr.Auth.JwtAuthFilter) 使用 
    JwtAuthFilter(ctx *context.Context)
    
    //CreateToken 生成token，用于登录成功后生成token 返回给用户
    CreateToken(info TokenInfo) string
    
    //ParseFromRequestToken 通过传入*http.Request验证token，用于自定义 handler 验证
	ParseFromRequestToken(req *http.Request) (TokenInfo, error)
}

```

### 数据存储

- mongodb

`pkg/mongo/db.go`

```golang
...
    s := &store{dbc}
    //注册你的store
    //_ = s
	store := &core.DBStore{
        User: s,
    }
...

```

## 开发规范

[Go 语言编码规范](docs/go-style-guide.md)



## 开发编译

生成文件放在template文件夹，文件开头注释文件生成路径`//pkg/test/test.go`

```

make gen

make bin

```

