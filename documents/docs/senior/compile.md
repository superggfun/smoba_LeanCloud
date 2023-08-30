---
title: 编译
lang: 'zh-CN'
permalink: /senior/compile
---

## 下载源码
* [Github厂库地址](https://github.com/superggfun/smoba)

## 编译使用
* 电脑安装了Go语言环境
* Vscode/Goland

### Linux/华为云函数
```sh
go mod tidy
set GOOS=linux
set GOARCH=amd64
go build -o handler main.go
zip main.zip handler config.go static
```

### Windows
```sh
go mod tidy
set GOOS=windows
set GOARCH=386
go build -o handler main.go
```