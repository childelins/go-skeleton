package v1

import (
	"github.com/childelins/go-skeleton/app/model"
	"github.com/childelins/go-skeleton/app/service"
	"github.com/childelins/go-skeleton/pkg/resp"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController

	userSrv *service.UserService
}

func NewUserController(userSrv *service.UserService) *UserController {
	return &UserController{
		userSrv: userSrv,
	}
}

func (u *UserController) Create(c *gin.Context) {
	userModel := &model.User{
		Nickname: "kangkang",
		Email:    "kang@gmail.com",
		Gender:   1,
	}

	ok := u.userSrv.Create(userModel)
	if ok {
		resp.Success(c, "创建成功")
	} else {
		resp.Error(c, "创建失败")
	}
}

func (u *UserController) List(c *gin.Context) {
	users := u.userSrv.List(1, 10)
	resp.Success(c, users)
}
