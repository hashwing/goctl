# goctl 

golang 后台脚手架，beego + jwt

## 安装

```
go install  git.gzsunrun.cn/sunruniaas/goctl

```

## 使用

```

./goctl -app demo -dir demo -mod github.com/hashwing/demo -mongo=true

```

`app`: 应用名称，默认:server

`dir`: 代码文件夹，默认: 当前目录

`mod`: go mod，默认：server

`mongo` 初始化mongo 作为数据存储，默认: false


## 开发编译

生成文件放在template文件夹，文件开头注释文件生成路径`//pkg/test/test.go`

```

make gen

make bin

```

