// Code generated by hertz generator.

package main

import (
	"douyin/service/api/biz/comment"
	"douyin/service/api/biz/message"
	"douyin/service/api/biz/user"
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
	}
	userGroup.GET("", user.UserInfo)
	// video service

	// comment service
	commentGroup := douyin.Group("/comment")
	{
		commentGroup.POST("/action", comment.Action)
		commentGroup.GET("/list", comment.List)
	}

	// favorite service

	// relation service

	// message service
	messageGroup := douyin.Group("/message")
	{
		messageGroup.POST("/action", message.Action)
		messageGroup.GET("/chat", message.Chat)
	}
}
