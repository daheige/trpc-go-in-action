package main

import (
	"context"
	"fmt"

	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/config"
	"trpc.group/trpc-go/trpc-go/log"

	"helloworld/pb"
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
	s := trpc.NewServer()
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
