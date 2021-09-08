package api

import (
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"go-demo/forms"
	"go-demo/global"
	"go-demo/model"
	"go-demo/services"
	"go.uber.org/zap"
	"net/http"
)

func GetUsers(c *gin.Context) {
	result := global.NewResult(c)
	UsersInfo := services.GetUsers()
	result.Success(UsersInfo)
	return
}

func SaveUser(c *gin.Context) {
	result := global.NewResult(c)
	var userRequest forms.SavaUserRequest
	var user model.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		zap.S().Error(err)
		result.Error(http.StatusNotFound, "参数错误")
		return
	}
	user.UserName = userRequest.UserName
	user.PhoneNum = userRequest.PhoneNum
	HashPasswd := common.EncryptPassword(userRequest.PassWord)
	user.Password = HashPasswd
	if err := services.SaveUser(&user); err != nil {
		result.Error(http.StatusInternalServerError, "创建失败")
		return
	}
	result.Success(nil)
	return
}

func GetOneUser(c *gin.Context) {
	result := global.NewResult(c)
	id := c.Param("id")
	UserInfo := services.GetUser(id)
	result.Success(UserInfo)
	return
}

func UpdateUser(c *gin.Context) {
	var user model.User
	result := global.NewResult(c)
	if err := c.ShouldBindJSON(&user); err != nil {
		result.Error(http.StatusBadRequest, "参数有误")
		return
	}
	if err := services.UpdateUser(&user); err != nil {
		result.Error(http.StatusBadRequest, "更新失败")
		return
	}
	result.Success(nil)
	return
}

func DeleteUser(c *gin.Context) {
	result := global.NewResult(c)
	id := c.Param("id")
	if err := services.DeleteUser(id); err != nil {
		result.Error(http.StatusBadRequest, "删除有误")
		return
	}
	result.Success(nil)
	return
}

func LoginUser(c *gin.Context) {
	result := global.NewResult(c)
	var json forms.UserLoginRequest
	if err := c.BindJSON(&json); err != nil {
		result.Error(http.StatusBadRequest, "参数错误")
		return
	}
	user, err := services.UserLogin(&json)
	if err != nil {
		result.Error(http.StatusBadRequest, "登入失败")
		return
	}
	data, err := services.GenerateToken(user)
	if err != nil {
		result.Error(http.StatusBadRequest, err.Error())
		return
	}
	result.Success(data)
	return
}
