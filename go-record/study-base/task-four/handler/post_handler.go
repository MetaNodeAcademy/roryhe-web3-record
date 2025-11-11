package handler

import "github.com/rory7/task-four/service"

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
