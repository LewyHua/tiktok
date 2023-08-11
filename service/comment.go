package service

import (
	"fmt"
	goredis "github.com/redis/go-redis/v9"
	"log"
	"project/dao/mysql"
	"project/dao/redis"
	"project/models"
	"strconv"
	"time"
)

func AddComment(videoId, userId uint, content string) (models.CommentResponse, error) {
	// 评论信息
	commentData := models.Comment{
		VideoId: videoId,
		UserId:  userId,
		Content: content,
	}
	// 新增评论，返回评论id
	_, err := mysql.AddComment(&commentData)
	if err != nil {
		return models.CommentResponse{}, err
	}
	log.Println("===========CommentID: " + strconv.Itoa(int(commentData.ID)))

	go func() {
		// 如果当前video的commentCount为0，不确定是没有评论，还是评论刚刚过期，所以不能直接+1
		// 所以需要先去看一下redis，如果有key，直接+1
		// 如果没key，更新commentCount再+1
		_, err := redis.GetCommentCountByVideoId(videoId)
		// 如果redis不存在key
		if err == goredis.Nil {
			// 获取最新commentCount
			cnt, err := mysql.GetCommentCnt(videoId)
			if err != nil {
				log.Println("mysql获取评论数失败", err)
				return
			}
			// 设置最新commentCount
			err = redis.SetCommentCountByVideoId(videoId, cnt)
			if err != nil {
				log.Println("redis更新评论数失败", err)
				return
			}
		}
		// 更新commentCount
		err = redis.IncrementCommentCountByVideoId(videoId)
		if err != nil {
			log.Printf("更新videoId为%v的评论数失败  %v\n", videoId, err.Error())
		}
	}()

	// 查询user
	user, exist := GetUserInfoByUserId(uint(userId))
	if !exist {
		fmt.Println("根据评论中的user_id找用户失败, 评论ID为：", commentData.ID)
		return models.CommentResponse{}, err
	}

	// 封装返回数据
	var commentResp models.CommentResponse
	commentResp.Id = int64(commentData.ID)
	commentResp.User = user
	commentResp.Content = content
	commentResp.CreateDate = models.TranslateTime(commentData.CreatedAt.Unix(), time.Now().Unix())

	return commentResp, nil
}

func GetCommentList(videoId uint) ([]models.CommentResponse, error) {
	// 1、根据videoId查询数据库，获取comments信息
	comments, err := mysql.FindCommentsByVideoId(videoId)
	if err != nil {
		fmt.Println("根据视频ID取评论失败")
		return nil, err
	}

	commentList := make([]models.CommentResponse, 0)
	for _, comment := range comments {
		user, exist := mysql.FindUserByID(comment.UserId)
		if !exist {
			fmt.Println("根据评论中的user_id找用户失败")
		}
		userResp := models.UserResponse{Id: user.Id, Name: user.Name}
		commentResp := models.CommentResponse{
			Id:         int64(comment.ID),
			User:       userResp,
			Content:    comment.Content,
			CreateDate: models.TranslateTime(comment.CreatedAt.Unix(), time.Now().Unix()),
		}
		commentList = append(commentList, commentResp)
	}

	return commentList, nil
}

func DeleteComment(videoId, userId, commentId uint) (models.CommentResponse, error) {

	// 查询冗余字段
	// 查询comment
	comment, err := mysql.FindCommentById(commentId)
	if err != nil {
		fmt.Println("查询评论失败")
		return models.CommentResponse{}, err
	}

	// 查询user
	user, exist := GetUserInfoByUserId(comment.UserId)
	if !exist {
		log.Println("根据评论中的user_id找用户失败")
	}

	// 封装返回数据
	var commentResp models.CommentResponse
	commentResp.Id = int64(comment.ID)
	commentResp.User = user
	commentResp.Content = comment.Content
	commentResp.CreateDate = models.TranslateTime(comment.CreatedAt.Unix(), time.Now().Unix())

	// 1、 redis评论数-1
	err = redis.DecrementCommentCountByVideoId(videoId)
	if err != nil {
		log.Println("redis评论数-1失败")
		return models.CommentResponse{}, err
	}

	// 2、 mysql删除comment
	err = mysql.DeleteCommentById(commentId)
	if err != nil {
		fmt.Println("删除Comment失败")
		return commentResp, err
	}

	return commentResp, nil
}

// GetCommentCount 根据视频ID获取视频的评论数
func GetCommentCount(videoId uint) (int64, error) {
	// 从redis中获取评论数
	count, err := redis.GetCommentCountByVideoId(videoId)
	if err != nil {
		log.Println("从redis中获取评论数失败：", err)
		return 0, err
	}
	// 1. 缓存中有数据, 直接返回
	if err != nil {
		return 0, err
	}
	if count > 0 {
		log.Println("从redis中获取评论数成功：", count)
		return count, nil
	}

	// 2. 缓存中没有数据，从数据库中获取
	num, err := mysql.GetCommentCnt(videoId)
	if err != nil {
		log.Println("从数据库中获取评论数失败：", err.Error())
		return 0, nil
	}
	log.Println("从数据库中获取评论数成功：", num)
	// 将评论数写入redis
	go func() {
		err = redis.SetCommentCountByVideoId(videoId, num)
		if err != nil {
			log.Println("将评论数写入redis失败：", err.Error())
		}
	}()
	return num, nil
}
