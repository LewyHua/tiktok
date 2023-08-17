package pack

import (
	"douyin/dal/model"
	"douyin/kitex_gen/message"
)

func Messages(messageModels []*model.Message) []*message.Message {
	if messageModels == nil {
		return nil
	}
	messages := make([]*message.Message, 0, len(messageModels))
	for _, msg := range messageModels {
		messages = append(messages, Message(msg))
	}
	return messages

}

func Message(messageModel *model.Message) *message.Message {
	if messageModel == nil {
		return nil
	}
	return &message.Message{
		Id:         int32(messageModel.ID),
		ToUserId:   int32(messageModel.ToUserId),
		FromUserId: int32(messageModel.FromUserId),
		Content:    messageModel.Content,
		CreateTime: &messageModel.CreateTime,
	}
}
