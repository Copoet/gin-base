# Go Gin Example [![rcard](https://goreportcard.com/badge/github.com/EDDYCJY/go-gin-example)](https://goreportcard.com/report/github.com/EDDYCJY/go-gin-example) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/EDDYCJY/go-gin-example) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/EDDYCJY/go-gin-example/master/LICENSE)

An example of gin contains many useful features

## Installation
```
go get github.com/EDDYCJY/go-gin-example
```

## How to run

### Required

- Mysql
- Redis

### Ready

Create a **blog database** and import [SQL](https://github.com/EDDYCJY/go-gin-example/blob/master/docs/sql/blog.sql)

### Conf

You should modify `conf/app.ini`

```
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
...
```

### Run
```
$ cd $GOPATH/src/go-gin-example

$ go run main.go 
```

Project information and existing API
