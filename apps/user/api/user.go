package main

import (
	"flag"
	"fmt"

	"github.com/palp1tate/easy-im/pkg/resultx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/palp1tate/easy-im/apps/user/api/internal/config"
	"github.com/palp1tate/easy-im/apps/user/api/internal/handler"
	"github.com/palp1tate/easy-im/apps/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
	httpx.SetOkHandler(resultx.OkHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
