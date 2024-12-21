package main

import (
	"context"
	"flag"

	"github.com/xianlianghe0123/go-web-mvc-tpl/cmd/migration/wire"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/config"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
)

func main() {
	var envConf = flag.String("conf", "config/dev.yaml", "config path, eg: -conf ./config/dev.yaml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
