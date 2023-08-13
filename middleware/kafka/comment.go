package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"project/dao/mysql"
	"project/models"
)

type CommentMQ struct {
	Topic    string
	GroupId  string
	Producer *kafka.Writer
	Consumer *kafka.Reader
}

var (
	CommentMQInstance *CommentMQ
)

func InitCommentKafka() {
	CommentMQInstance = &CommentMQ{
		Topic:   "comments",
		GroupId: "comment_group",
	}

	// 创建 Comment 业务的生产者和消费者实例
	CommentMQInstance.Producer = kafkaManager.NewProducer(CommentMQInstance.Topic)
	CommentMQInstance.Consumer = kafkaManager.NewConsumer(CommentMQInstance.Topic, CommentMQInstance.GroupId)

	go CommentMQInstance.Consume()
}

// ProduceAddCommentMsg 发布添加评论的消息, 向mysql中添加评论时, 调用此方法
func (m *CommentMQ) ProduceAddCommentMsg(message *models.Comment) {
	err := kafkaManager.ProduceMessage(m.Producer, message)
	if err != nil {
		fmt.Println("kafka发送添加评论的消息失败：", err)
		return
	}
}

// ProduceDelCommentMsg 发布删除评论的消息, mysql删除评论时, 调用此方法
func (m *CommentMQ) ProduceDelCommentMsg(commentId uint) {
	err := kafkaManager.ProduceMessage(m.Producer, commentId)
	if err != nil {
		fmt.Println("kafka发送删除评论的消息失败：", err)
		return
	}
}

// Consume 消费添加或者删除评论的消息
func (m *CommentMQ) Consume() {
	for {
		msg, err := m.Consumer.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("[CommentMQ]从消息队列中读取消息失败:", err)
		}

		fmt.Printf("[CommentMQ]从消息队列中读取到消息: %v\n", msg)

		// 发送确认
		err = m.Consumer.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println("Failed to commit message:", err)
		}

		var result json.RawMessage
		err = json.Unmarshal(msg.Value, &result)
		if err != nil {
			fmt.Println("[CommentMQ]解析消息失败:", err)
			continue
		}

		// 解析消息, 消息类型可能为models.Comment, 也可能为CommentId, 如果是前者, 则添加评论, 如果是后者, 则删除评论
		// 解析为models.Comment, 则向数据库中添加评论
		message := new(models.Comment)
		err = json.Unmarshal(result, message)
		if err == nil {
			_, err = mysql.AddComment(message)
			if err != nil {
				fmt.Println("[CommentMQ]向mysql中添加评论失败:", err)
			}
			fmt.Println("[CommentMQ]向mysql中添加评论成功")
			continue
		}

		// 解析为整型, 即CommentId, 则从数据库中删除评论
		commentId := new(uint)
		err = json.Unmarshal(result, commentId)
		if err == nil {
			err = mysql.DeleteCommentById(*commentId)
			if err != nil {
				fmt.Println("[CommentMQ]从mysql中删除评论失败:", err)
			}
			fmt.Println("[CommentMQ]从mysql中删除评论成功")
			continue
		}
		fmt.Println("[CommentMQ]解析消息失败:", result)
	}
}
