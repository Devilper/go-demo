### go-demo 项目目录结构
go-demo
├── README.md
├── api
│   └── user.go
├── common
├── config
│   └── config.go
├── config-dev.yaml
├── forms
│   └── user.go
├── global
│   ├── global.go
│   ├── response
│   │   └── user.go
│   └── result.go
├── go.mod
├── go.sum
├── initialize
│   ├── config.go
│   ├── database.go
│   ├── logger.go
│   ├── router.go
│   └── validator.go
├── log
│   └── user.log
├── main.go
├── middleware
│   └── jwt
│       └── jwt.go
├── model
│   ├── main.go
│   └── user.go
├── routers
│   └── user.go
├── services
│   ├── token_serivices.go
│   └── user_services.go
└── validator
    └── validators.go

项目修改自启动方式
1、下载依赖
    go get github.com/pilu/fresh
2、在跟目录下启动
    fresh
