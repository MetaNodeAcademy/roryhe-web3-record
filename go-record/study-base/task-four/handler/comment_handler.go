package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rory7/task-four/service"
)

/**
评论功能
实现评论的创建功能，已认证的用户可以对文章发表评论。
实现评论的读取功能，支持获取某篇文章的所有评论列表。
*/

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		commentService: service.NewCommentService(),
	}
}

func (h CommentHandler) CreateComment(c *gin.Context) {}

func (h CommentHandler) GetByPostID(c *gin.Context) {}
