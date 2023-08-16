// Code generated by hertz generator.

package main

import (
	"douyin/service/api/biz/comment"
	"douyin/service/api/biz/message"
	"douyin/service/api/biz/user"
	"douyin/service/api/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	douyin := r.Group("/douyin")
	// user service
	userGroup := douyin.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.GET("", mw.AuthWithoutLogin(), user.Info)
	}

	// video service

	// comment service
	commentGroup := douyin.Group("/comment")
	{
		commentGroup.POST("/action", mw.Auth(), comment.Action)
		commentGroup.GET("/list", mw.AuthWithoutLogin(), comment.List)
	}

	// favorite service

	// relation service

	// message service
	messageGroup := douyin.Group("/message")
	{
		messageGroup.POST("/action", mw.Auth(), message.Action)
		messageGroup.GET("/chat", mw.Auth(), message.Chat)
	}
}
