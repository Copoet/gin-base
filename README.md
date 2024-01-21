# Go Gin Base
An example of gin contains many useful features

## Installation
```
go get github.com/copoet/gin-base
```

## How to run

### Required

- Mysql
- Redis

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

### Run
```
$ cd $GOPATH/src/gin-base

$ go run ./cmd/main.go 
```
### Swagger
#### 新增api后需要执行下面命令 可生成swagger文档
```
swag init 
```