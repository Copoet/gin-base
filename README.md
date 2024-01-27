# Go Gin Base
An example of gin contains many useful features

## Installation
```
go get github.com/copoet/gin-base
```
## Project Dir Desc
```
project_name/
├── conf/              # 配置文件
  ├──app.ini           # 配置文件
├── docs/              # Swagger文档
├── internal           # 项目内部文件
  ├── middleware/        # 中间件
  ├── models/            # 数据库模型
    ├──user.go           # user模型
  ├── controllers/       # 控制器
  ├── routes/            # 路由目录
    ├──api.go            # 路由文件
  ├── services/          # 服务类
├── pkg/                 # 工具包
└── main.go              # 应用入口
```
### Conf
You should set  `conf/app.ini`

```
[app]
Port = 8080
PrefixUrl = http://127.0.0.1:8091
RunMode = debug
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200

```
## How to run

### Required
- Mysql
- Redis

### Run
```
$ cd $GOPATH/src/gin-base

$ go run main.go 
```
### Swagger
#### 新增api后需要执行下面命令 可生成swagger文档
```
swag init 
```