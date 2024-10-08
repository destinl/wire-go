//go:build wireinject

package main

import (
	"wire-go/handler"
	"wire-go/ioc"
	"wire-go/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(
		handler.NewPostHandler,
		service.NewPostService,
		ioc.NewGinEngineAndRegisterRoute,
	)
	return &gin.Engine{}
}

// 使用 wire.NewSet 函数将提供者进行分组，该函数返回一个 ProviderSet 结构体。
// 不仅如此，wire.NewSet 还能对多个 ProviderSet 进行分组 wire.NewSet(PostSet, XxxSet) 。
func InitializeAppV2() *gin.Engine {
	wire.Build(
		handler.PostSet,
		ioc.NewGinEngineAndRegisterRoute,
	)
	return &gin.Engine{}
}

func InitializeAppV3() *gin.Engine {
	wire.Build(
		handler.NewPostHandler,
		service.NewPostServiceV2,
		ioc.NewGinEngineAndRegisterRoute,
		wire.Bind(new(service.IPostService), new(*service.PostService)),
	)
	return &gin.Engine{}
}