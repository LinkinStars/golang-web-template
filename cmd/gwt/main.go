package main

import (
	"flag"

	"github.com/LinkinStars/go-scaffold/contrib/log/zap"
	"github.com/LinkinStars/go-scaffold/contrib/static_config/file"
	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/golang-web-template/internal/base/conf"
	"github.com/gin-gonic/gin"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the project
	Name = "golang-web-template"
	// Version is the version of the project
	Version string
	// confFlag is the config flag.
	confFlag string
)

func init() {
	flag.StringVar(&confFlag, "c", "../../configs/config.yaml", "config path, eg: -c config.yaml")
}

type App struct {
	log    logger.Logger
	server *gin.Engine
}

func newApp(logger logger.Logger, server *gin.Engine) *App {
	return &App{
		log:    logger,
		server: server,
	}
}

// @title golang-web-template api documentation
// @version 1.0
// @host localhost:8080
// @BasePath /gwt/api/v1
func main() {
	flag.Parse()

	// init logger
	logger.SetLogger(zap.NewLogger(logger.LevelDebug, zap.WithName(Name), zap.WithCallerFullPath()))

	// init config
	c := &conf.AllConfig{}
	if err := file.NewStaticConfigParser(confFlag).LoadAndSet(c); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(c.Server, c.Data, logger.GetLogger())
	if err != nil {
		panic(err)
	}
	defer cleanup()
	logger.Infof("server start, listen on %s", c.Server.HTTP.Addr)
	err = app.server.Run(c.Server.HTTP.Addr)
	if err != nil {
		panic(err)
	}
}
