//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/golang-web-template/internal/base/conf"
	"github.com/LinkinStars/golang-web-template/internal/base/server"
	"github.com/LinkinStars/golang-web-template/internal/biz"
	"github.com/LinkinStars/golang-web-template/internal/controller"
	"github.com/LinkinStars/golang-web-template/internal/data"
	"github.com/LinkinStars/golang-web-template/internal/router"
	"github.com/google/wire"
)

// initApp init admin application.
func initApp(*conf.Server, *conf.Data, logger.Logger) (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, router.ProviderSet, controller.ProviderSet,
		biz.ProviderSet, data.ProviderSet, newApp))
}
