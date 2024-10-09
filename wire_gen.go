// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"wire-go/handler"
	"wire-go/ioc"
	"wire-go/service"
)

// Injectors from wire.go:

func InitializeApp() *gin.Engine {
	iPostService := service.NewPostService()
	postHandler := handler.NewPostHandler(iPostService)
	engine := ioc.NewGinEngineAndRegisterRoute(postHandler)
	return engine
}

// 使用 wire.NewSet 函数将提供者进行分组，该函数返回一个 ProviderSet 结构体。
// 不仅如此，wire.NewSet 还能对多个 ProviderSet 进行分组 wire.NewSet(PostSet, XxxSet) 。
func InitializeAppV2() *gin.Engine {
	iPostService := service.NewPostService()
	postHandler := handler.NewPostHandler(iPostService)
	engine := ioc.NewGinEngineAndRegisterRoute(postHandler)
	return engine
}

func InitializeAppV3() *gin.Engine {
	postService := service.NewPostServiceV2()
	postHandler := handler.NewPostHandler(postService)
	engine := ioc.NewGinEngineAndRegisterRoute(postHandler)
	return engine
}

// 结构体提供者（Struct Providers）
func InitializeUser() *User {
	name := NewName()
	publicAccount := NewPublicAccount()
	user := &User{
		MyName:          name,
		MyPublicAccount: publicAccount,
	}
	return user
}

// wire.go:

type Name string

func NewName() Name {
	return "陈明勇"
}

type PublicAccount string

func NewPublicAccount() PublicAccount {
	return "公众号:Go技术干货"
}

type User struct {
	MyName          Name
	MyPublicAccount PublicAccount
}
