package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RpcServer zrpc.RpcServerConf // rpc服务端
	RestConf  rest.RestConf      // rest服务端
}
