package main

import (
	"flag"
	"fmt"

	"github.com/yl3715a/testrpc/rpc/internal/config"
	"github.com/yl3715a/testrpc/rpc/internal/server"
	"github.com/yl3715a/testrpc/rpc/internal/svc"
	"github.com/yl3715a/testrpc/rpc/types/restrpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/restrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		restrpc.RegisterRestrpcServer(grpcServer, server.NewRestrpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
