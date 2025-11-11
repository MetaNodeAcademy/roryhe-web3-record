package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rory7/task-four/models"
	"github.com/rory7/task-four/service"
	"github.com/rory7/task-four/utils"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	//登陆参数validate
	if !utils.ValidateRequired(user.Name) || !utils.ValidateRequired(user.Password) {
		utils.BadRequest(c, "用户名或密码不能为空")
		return
	}

	loginUser, err := h.userService.GetUserByName(user.Name)
	if err != nil {
		utils.InternalServerError(c, "用户名或密码错误")
		return
	}

	expireDuration := time.Hour * 24
	token, err := utils.GenerateToken(loginUser.ID, loginUser.Name, expireDuration)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, token)
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body models.User true "用户信息"
// @Success 200 {object} utils.Response
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 参数验证
	if !utils.ValidateRequired(user.Name) {
		utils.BadRequest(c, "用户名不能为空")
		return
	}

	if !utils.ValidateRequired(user.Password) {
		utils.BadRequest(c, "密码不能为空")
		return
	}

	if !utils.ValidateRequired(user.Email) {
		utils.BadRequest(c, "邮箱不能为空")
		return
	}

	if !utils.ValidateEmail(user.Email) {
		utils.BadRequest(c, "邮箱格式不正确")
		return
	}

	//注册时对密码进行加密
	user.Password = utils.MD5String(user.Password)

	// 调用服务层创建用户
	if err := h.userService.CreateUser(&user); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.Success(c, "")
}

// GetAllUsers 获取所有用户
// @Summary 获取用户列表
// @Description 获取所有用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, users)
}
