package handler

import (
	"net/http"
	_ "net/http/pprof"

	"dist-encoder/app/manager/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{http.MethodGet, "/hello", NilHandler(serverCtx)},
		},
	)

}

func NilHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("nil function"))
	}
}
