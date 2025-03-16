# tRPC官方
https://trpc.group

# tRPC命令工具安装
推荐使用linux系统或mac系统安装tRPC命令行工具，tRPC cmdline安装方式如下：
1. install protoc
- mac系统安装方式如下：
```shell
brew install protobuf
```
- linux系统安装方式如下：
```shell
# Reference: https://grpc.io/docs/protoc-installation/
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip -o protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
export PATH=~/.local/bin:$PATH # Add this to your `~/.bashrc`.
protoc --version
libprotoc 3.15.8
```
2. install flatc
- linux系统安装方式
```shell
# Reference: https://github.com/google/flatbuffers/releases
wget https://github.com/google/flatbuffers/releases/download/v23.5.26/Linux.flatc.binary.g++-10.zip
unzip -o Linux.flatc.binary.g++-10.zip -d $HOME/.bin
export PATH=~/.bin:$PATH # Add this to your `~/.bashrc`.

# check flatc version
flatc --version
flatc version 23.5.26
```
- mac系统安装方式
```shell
cd ~
wget https://github.com/google/flatbuffers/releases/download/v25.2.10/MacIntel.flatc.binary.zip
unzip -o MacIntel.flatc.binary.zip -d $HOME/.bin
# after you should add flatc to PATH
export PATH=~/.bin:$PATH # Add this to your ~/.bashrc.

# check flatc version
flatc --version
flatc version 25.2.10
```

3. install go tools
```shell
# install protoc-gen-go
# Reference: https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# install goimports
go install golang.org/x/tools/cmd/goimports@latest

# install mockgen
go install go.uber.org/mock/mockgen@latest

# install protoc-gen-validate and protoc-gen-validate-go
go install github.com/envoyproxy/protoc-gen-validate@latest
go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest
```
当tRPC命令行工具安装好后，就可以参考：https://github.com/trpc-group/trpc-cmdline/blob/main/README.zh_CN.md 创建一个微服务项目

# tRPC架构设计
https://trpc.group/zh/docs/what-is-trpc/archtecture_design/

# tRPC插件生态
https://trpc.group/zh/docs/what-is-trpc/plugin_ecosystem/

# tRPC-Go使用手册
https://github.com/trpc-group/trpc-go/blob/main/docs/README.zh_CN.md

# tRPC服务发现和注册
https://github.com/trpc-ecosystem/go-naming-consul
https://github.com/trpc-ecosystem/go-naming-consul/blob/main/README.zh_CN.md

# etcd配置管理
tRPC-Go etcd configuration plugin: https://github.com/trpc-ecosystem/go-config-etcd
