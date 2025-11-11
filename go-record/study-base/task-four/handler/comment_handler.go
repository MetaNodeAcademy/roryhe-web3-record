package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rory7/task-four/models"
	"github.com/rory7/task-four/service"
	"github.com/rory7/task-four/utils"
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

func (h CommentHandler) CreateComment(c *gin.Context) {
	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	ID, err := h.commentService.CreateComment(&comment)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, ID)
}

func (h CommentHandler) GetByPostID(c *gin.Context) {
	postID := c.Param("post_id")
	if !utils.ValidateRequired(postID) {
		utils.BadRequest(c, "post_id is required")
		return
	}

	id, _ := strconv.ParseUint(postID, 10, 64)
	comments, err := h.commentService.GetCommentByPostID(uint(id))
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, comments)

}
