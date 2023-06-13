package main

import (
	"flag"
	"fmt"

	"dist-encoder/app/manager/internal/config"
	"dist-encoder/app/manager/internal/handler"
	"dist-encoder/app/manager/internal/server"
	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/manager.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	group := service.NewServiceGroup()

	s := zrpc.MustNewServer(c.RpcServer, func(grpcServer *grpc.Server) { // rpc服务
		distribute.RegisterDistributeServer(grpcServer, server.NewDistributeServer(ctx))

		if c.RpcServer.Mode == service.DevMode || c.RpcServer.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	r := rest.MustNewServer(c.RestConf) // rest服务
	handler.RegisterHandlers(r, ctx)    // 路由注册

	group.Add(s) // rpc服务注册
	group.Add(r) // rest服务注册

	defer group.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.RpcServer.ListenOn)
	fmt.Printf("Starting server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)

	group.Start()
}
