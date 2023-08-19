package message

import (
	"context"
	"douyin/common"
	"douyin/constant"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/message/messageservice"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
)

var messageClient messageservice.Client

func init() {
	// Etcd 服务发现
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	messageClient, err = messageservice.NewClient(
		constant.MessageServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.MessageServiceName}),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Action(ctx context.Context, c *app.RequestContext) {
	toUserIdStr := c.Query("to_user_id")
	toUserId, err := strconv.ParseUint(toUserIdStr, 10, 64)
	content := c.Query("content")
	fromUserId, err := common.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("Get user id from ctx", zap.Error(err))
		c.JSON(http.StatusOK, "Unauthorized operation.")
		return
	}
	resp, err := messageClient.MessageAction(ctx, &message.MessageActionRequest{
		FromUserId: int32(fromUserId),
		ToUserId:   int32(toUserId),
		ActionType: 1,
		Content:    content,
	})
	if err != nil {
		zap.L().Error("Message action error", zap.Error(err))
		c.JSON(http.StatusOK, &message.MessageActionResponse{
			StatusCode: 1,
			StatusMsg:  thrift.StringPtr("Server internal error"),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func Chat(ctx context.Context, c *app.RequestContext) {
	fromUserId, err := common.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("Get user id from ctx", zap.Error(err))
		c.JSON(http.StatusOK, "Unauthorized operation.")
		return
	}
	toUserIdStr := c.Query("to_user_id")
	preMsgTimeStr := c.Query("pre_msg_time")
	preMsgTime, err := strconv.ParseInt(preMsgTimeStr, 10, 64)
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		zap.L().Error("Parse param err", zap.Error(err))
		c.JSON(http.StatusOK, "Invalid param.")
		return
	}

	resp, err := messageClient.MessageChat(ctx, &message.MessageChatRequest{
		FromUserId: int32(fromUserId),
		ToUserId:   int32(toUserId),
		PreMsgTime: preMsgTime,
	})
	if err != nil {
		zap.L().Error("Message chat error", zap.Error(err))
		c.JSON(http.StatusOK, &message.MessageChatResponse{
			StatusCode: 1,
			StatusMsg:  thrift.StringPtr("Server internal error"),
		})
		return
	}
	c.JSON(http.StatusOK, resp)

}
