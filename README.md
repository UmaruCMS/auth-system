# Umaru CMS - Auth System

## Notice

### dep

依赖发生变化时要执行 `dep ensure`

### grpc

`protoc -I rpc/ rpc/base/auth/main.proto --go_out=plugins=grpc:rpc`

## TODO List

### 架构

* [x] MySQL + gorm 集成
* [x] Redis 集成
* [x] gin 集成
* [x] grpc 集成

### 功能

* [x] 用户注册 (HTTP)
* [x] 用户登录 (HTTP)
* [ ] 获取用户信息 (HTTP / RPC)
* [ ] Token 校验 (RPC)
