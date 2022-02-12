# pw_server

#### 介绍
完美世界服务端

#### 安装
go mod tidy

#### 使用
air

#### 编译Windows平台
go build -ldflags "-s -w -H windowsgui" main.go

#### 编译Linux平台
SET CGO_ENABLED=0\
SET GOARCH=amd64\
SET GOOS=linux\
go build -ldflags "-s -w" main.go