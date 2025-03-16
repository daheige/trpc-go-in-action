# tRPC-GO支持grpc协议
参考文档：https://github.com/trpc-ecosystem/go-codec/blob/main/grpc/README.zh_CN.md

设计思想如下：
- tRPC-GO通过对grpc协议实现codec封装，来达到支持grpc协议。
- 主要通过grpc server transport和编解码来支持grpc server处理以及grpc client的请求。

# 快速开始
1. 首先安装好golang(go version >= 1.18)以及trpc-cmdline相关工具，具体安装方式参考[trpc-cmdline工具和相关依赖安装](../README.md)
2. 执行如下命令，创建一个go项目
```shell
mkdir trpc-grpc-example
cd trpc-grpc-example
# 初始化go.mod文件
go mod init trpc-grpc-example
```
3. 在`trpc-grpc-example`目录中新建proto目录，并创建`helloworld.proto`文件，以及添加如下protobuf内容
```protobuf
syntax = "proto3";

// 这里的包名，可以根据实际情况更改，例如：package trpc.helloworld
package trpc.helloworld;

// 这里的pb前面的路径可以忽略，或者可以将生成的代码放在制定的git仓库中
option go_package="github.com/some-repo/examples/pb";

service Greeter {
  rpc Hello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string msg = 1;
}

message HelloReply {
  string msg = 1;
}
```
4. 在`trpc-grpc-example`目录中新建Makefile文件，并添加如下内容
```makefile
# 生成trpc-go代码
gen-pb:
	trpc create \
		--protocol=grpc \
		-p proto/helloworld.proto \
		-o pb -f \
		--rpconly \
		--nogomod=true \
		--mock=false

.PHONY: gen-pb
```
接着，执行`make gen-pb`命令生成pb代码（该命令会是用`trpc-cmdline`工具读取proto文件并生成pb文件，包含服务端代码和客户端代码）。

5. 在`trpc-grpc-example`目录中新建server目录，并在server目录中新增server.go文件，以及在main.go文件中添加如下代码
```go
package main

import (
	"context"
	"fmt"

	"trpc.group/trpc-go/trpc-codec/grpc"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/config"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/server"

	"trpc-grpc-example/pb"
)

// AppConfig 自定义的配置文件
type AppConfig struct {
	AppDebug bool   `yaml:"app_debug"`
	AppEnv   string `yaml:"app_env"`
	// 其他配置省略...
}

func main() {
	// 加载自定义配置文件
	c, err := config.Load("app.yaml", config.WithCodec("yaml"), config.WithProvider("file"))
	if err != nil {
		fmt.Println("load config error:", err)
		return
	}

	var appConfig AppConfig
	err = c.Unmarshal(&appConfig)
	if err != nil {
		fmt.Println("unmarshal config error:", err)
		return
	}

	fmt.Printf("appConfig:%v\n", appConfig)

	// 创建tRPC服务实例，启动服务
	s := trpc.NewServer(server.WithTransport(grpc.DefaultServerTransport))
	s.RegisterOnShutdown(func() {
		// 优雅退出执行的shutdown操作...
		fmt.Println("server shutdown")
	})

	pb.RegisterGreeterService(s, &Greeter{})
	if err := s.Serve(); err != nil {
		log.Errorf("serve error:%v", err)
	}
}

// Greeter 实现trpc服务
type Greeter struct{}

// Hello 实现hello方法
func (g Greeter) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Infof("got hello request: %s", req.Msg)
	return &pb.HelloReply{Msg: "Hello " + req.Msg + "!"}, nil
}
```
6. 执行`go mod tidy`获取相关的依赖包
7. 替换`go.mod`中的protobuf包为下面的内容
```
replace google.golang.org/protobuf => google.golang.org/protobuf v1.32.0
```
当替换后，我们继续执行`go mod tidy`获取相关的依赖包。

8. 添加配置文件`trpc_go.yaml`文件，内容如下：
```yaml
global:                             # global config.
  namespace: development            # environment type, two types: production and development.
  env_name: test                    # environment name, names of multiple environments in informal settings.

server:                                            # server configuration.
  app: helloworld                                  # business application name.
  server: helloworld                               # service process name.
  service:                                         # business service configuration，can have multiple.
    - name: trpc.helloworld                        # the route name of the service.
      ip: 127.0.0.1                                # the service listening ip address, can use the placeholder ${ip}, choose one of ip and nic, priority ip.
      port: 8080                                  # the service listening port, can use the placeholder ${port}.
      network: tcp                                 # the service listening network type,  tcp or udp.
      protocol: grpc                               # application layer protocol, trpc or http.
      timeout: 1000                                # maximum request processing time in milliseconds.
      idletime: 300000                             # connection idle time in milliseconds.

# 插件配置
plugins:                                           # plugin configuration.
  log:                                             # logging configuration.
    default:                                       # default logging configuration, supports multiple outputs.
      - writer: console                            # console standard output, default setting.
        level: info                                # log level of standard output.
      - writer: file                               # local file logging.
        level: debug                               # log level of local file rolling logs.
        formatter: json                            # log format for standard output.
        writer_config:
          filename: ./trpc-grpc-example.log                     # storage path of rolling continuous log files.
          max_size: 10                             # maximum size of local log files, in MB.
          max_backups: 10                          # maximum number of log files.
          max_age: 7                               # maximum number of days to keep log files.
          compress: false
```
这里，为了演示加载自定义配置，我这里在当前项目中新增了app.yaml文件，并在上述server/server.go文件中读取了相关配置。

9. 启动server代码
```shell
go run server/server.go
```
运行结果如下：
```ini
appConfig:{true test}
2025-03-16 21:40:13.590	INFO	server/service.go:168	process:10884, grpc service:trpc.helloworld launch success, tcp:127.0.0.1:8080, serving ...
```

# grpcurl工具安装
- grpcurl工具主要用于grpcurl请求，可以快速查看grpc proto定义以及调用grpc service定义的方法。 
- grpcurl安装，参考文档：https://github.com/fullstorydev/grpcurl
  如果你本地安装了golang，那可以直接运行如下命令，安装grpcurl工具

```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
查看服务列表：
```shell
grpcurl -plaintext 127.0.0.1:8080 list
```
输出结果如下：
```ini
grpc.reflection.v1alpha.ServerReflection
trpc.helloworld.Greeter
```
查看proto文件定义的所有方法
```shell
grpcurl -plaintext 127.0.0.1:8080 describe trpc.helloworld.Greeter
```
输出结果如下：
```ini
trpc.helloworld.Greeter is a service:
service Greeter {
  rpc Hello ( .trpc.helloworld.HelloRequest ) returns ( .trpc.helloworld.HelloReply );
}
```

# 验证grpc服务接口
通过grpcurl工具，请求如下方法：
```shell
grpcurl -d '{"msg":"daheige"}' -plaintext 127.0.0.1:8080 trpc.helloworld.Greeter.Hello
```
运行结果如下：
```json
{
  "msg": "Hello daheige!"
}
```
此时，服务端输出日志如下：
```ini
2025-03-16 21:52:09.560	INFO	server/server.go:58	got hello request: daheige
```
当然，你也可以使用其他工具验证，例如：postman软件请求grpc服务对应的相关方法即可。
