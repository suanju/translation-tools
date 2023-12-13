package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"translation/internal/middleware"

	"translation/internal/config"
	"translation/internal/handler"
	"translation/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

//go:embed etc/lang.json
var JsonLang []byte

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	langInfo := &svc.LangInfo{}
	logx.MustSetup(c.Log)
	err := json.Unmarshal(JsonLang, langInfo)
	if err != nil {
		log.Fatalf("lang.json parse error")
	}

	server := rest.MustNewServer(c.RestConf, middleware.Cors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c, *langInfo)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}
