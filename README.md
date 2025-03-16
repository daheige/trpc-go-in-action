# tRPC官方文档
https://trpc.group

# trpc-cmdline工具和相关依赖安装
参考文档：https://github.com/trpc-group/trpc-cmdline/blob/main/README.zh_CN.md

这里，我推荐您使用linux系统或mac系统安装trpc-cmdline和trpc-go相关依赖，具体步骤如下：
1. 配置`~/.gitconfig`文件：
```yaml
# 添加如下配置
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```
2. 执行如下命令安装`trpc-cmdline`工具
```shell
go install trpc.group/trpc-go/trpc-cmdline/trpc@latest
```
查看trpc-go版本
```shell
trpc version
# 输出结果：trpc-group/trpc-cmdline version: v1.0.8
```
接下来，我们还需要继续执行3～5步，来安装trpc-go相关依赖。

3. install protoc
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
4. install flatc
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

5. install go tools
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
到这里，tRPC命令行工具安装完毕，此时我们就可以参考：https://github.com/trpc-group/trpc-cmdline/blob/main/README.zh_CN.md 创建一个微服务项目

# tRPC架构设计
https://trpc.group/zh/docs/what-is-trpc/archtecture_design/

# tRPC插件生态
https://trpc.group/zh/docs/what-is-trpc/plugin_ecosystem/

# tRPC-Go使用手册
https://github.com/trpc-group/trpc-go/blob/main/docs/README.zh_CN.md

# tRPC-database各种数据库快速接入
支持主流的各种数据库快速接入，参考文档：https://github.com/trpc-ecosystem/go-database/blob/main/README.zh_CN.md

- 在日常的开发过程中，开发者经常会访问 MySQL、Redis、Kafka 等存储进行数据库的读写。直接使用开源的 SDK 虽然可以满足访问数据库的需求，但是用户需要自己负责路由寻址、监控上报、配置的开发。
- 考虑到这些原因，tRPC-Go 提供了多种多样的路由寻址、监控上报、配置管理的插件，同时还封装了开源的 SDK，复用 tRPC-Go 插件的能力，减少重复代码。
- 同时，tRPC-Go 还提供了部分开源 SDK 的封装，可以直接复用 tRPC-Go 的路由寻址、监控上报等功能。

# tRPC-GO如何支持gRPC协议接入
参考文档：https://github.com/trpc-ecosystem/go-codec/blob/main/grpc/README.zh_CN.md

# tRPC-GO如何支持http协议接入
- 参考文档：https://github.com/trpc-group/trpc-go/blob/main/http/README.zh_CN.md
- 可以根据实际情况接入http标准库或者第三方http框架，例如：gin、gorilla/mux、go-chi/chi等框架

# tRPC服务发现和注册
名字服务参考：https://github.com/trpc-group/trpc-go/blob/main/docs/developer_guide/develop_plugins/naming.zh_CN.md

consul服务发现和注册使用参考如下文档：
- https://github.com/trpc-ecosystem/go-naming-consul
- https://github.com/trpc-ecosystem/go-naming-consul/blob/main/README.zh_CN.md
- 自定义服务发现和注册参考：https://github.com/trpc-group/trpc-go/blob/main/naming/README.zh_CN.md

# tRPC-GO服务可观测性接入
- prometheus接入： https://github.com/trpc-ecosystem/go-metrics-prometheus/blob/main/README.zh_CN.md
- opentelemetry接入：https://github.com/trpc-ecosystem/go-opentelemetry/blob/main/README.zh_CN.md

# tRPC-GO log日志包接入
参考文档：https://github.com/trpc-group/trpc-go/blob/main/log/README.zh_CN.md

你在代码中先直接引入log包，然后使用log包上的相关方法即可
```go
import (
    "trpc.group/trpc-go/trpc-go/log"
)

// ... 省略其他代码...
name := "abc"
index := 1
log.Infof("name: %s, index: %d", name, index)
```

# tRPC-GO recover组件接入
参考文档：https://github.com/trpc-ecosystem/go-filter/tree/main/recovery

# tRPC-GO 参数自动校验接入
参考文档：https://github.com/trpc-ecosystem/go-filter/tree/main/validation

# tRPC-GO go-gateway接入
参考文档：https://github.com/trpc-ecosystem/go-gateway

# tRPC-GO 配置插件接入
- 参考文档：https://github.com/trpc-group/trpc-go/blob/main/docs/developer_guide/develop_plugins/config.zh_CN.md
- tRPC-Go etcd configuration plugin: https://github.com/trpc-ecosystem/go-config-etcd
