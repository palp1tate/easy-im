package main

import (
	"flag"
	"fmt"

	"github.com/palp1tate/easy-im/apps/social/api/internal/config"
	"github.com/palp1tate/easy-im/apps/social/api/internal/handler"
	"github.com/palp1tate/easy-im/apps/social/api/internal/svc"
	"github.com/palp1tate/easy-im/pkg/resultx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/social.yaml", "the config file")

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
