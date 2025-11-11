package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rory7/task-four/models"
	"github.com/rory7/task-four/service"
	"github.com/rory7/task-four/utils"
)

/**
文章管理功能
实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
实现文章的更新功能，只有文章的作者才能更新自己的文章。
实现文章的删除功能，只有文章的作者才能删除自己的文章。
*/

type PostHandler struct {
	userService service.UserService
	postService service.PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		userService: service.NewUserService(),
		postService: service.NewPostService(),
	}
}

func (h PostHandler) CreatePost(c *gin.Context) {
	userID, existsId := c.Get("userID")

	if !existsId {
		utils.Unauthorized(c, "")
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if !utils.ValidateRequired(post.Title) {
		utils.BadRequest(c, "请输入文章标题")
		return
	}

	if !utils.ValidateRequired(post.Content) {
		utils.BadRequest(c, "请输入文章内容")
		return
	}

	post.UserId = userID.(uint)
	postID, err := h.postService.CreatePost(&post)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, postID)
}

func (h PostHandler) Get(c *gin.Context) {
	postId := c.Param("post_id")
	if !utils.ValidateRequired(postId) {
		utils.BadRequest(c, "post_id不能为空")
		return
	}

	id, _ := strconv.ParseUint(postId, 10, 64)
	post, err := h.postService.GetPostById(uint(id))
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, post)
}

func (h PostHandler) GetAll(c *gin.Context) {
	posts, err := h.postService.GetALLPosts()
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, posts)
}

func (h PostHandler) Put(c *gin.Context) {
	userID, existsId := c.Get("userID")

	if !existsId {
		utils.Unauthorized(c, "")
		return
	}

	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if post.ID == 0 {
		utils.BadRequest(c, "文章id不能为空")
		return
	}

	if !utils.ValidateRequired(post.Title) {
		utils.BadRequest(c, "请输入文章标题")
		return
	}

	if !utils.ValidateRequired(post.Content) {
		utils.BadRequest(c, "请输入文章内容")
		return
	}

	post.UserId = userID.(uint)
	err := h.postService.UpdatePost(&post)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, "")

}

func (h PostHandler) Delete(c *gin.Context) {

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err := h.postService.DeletePostById(post.ID)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, "")
}
