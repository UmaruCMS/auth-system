# Umaru CMS - Auth System

## Notice

### dep

依赖发生变化时要执行 `dep ensure`

### grpc

* `protoc -I rpc/ rpc/base/auth/main.proto --go_out=plugins=grpc:rpc`
* `protoc -I rpc/ rpc/base/user/main.proto --go_out=plugins=grpc:rpc`

## TODO List

### 架构

* [x] MySQL + gorm 集成
* [x] Redis 集成
* [x] gin 集成
* [x] grpc 集成

### 功能

* [x] 用户注册 (HTTP) （隔离普通用户和管理员 -> 暂不支持普通用户 -> 添加「新建用户」接口）
* [x] 用户登录 (HTTP)
* [ ] 获取用户列表
* [x] 获取用户信息 (HTTP / RPC)
* [x] Token 校验 (RPC)
