package api

import (
	"github.com/gin-gonic/gin"
	"go-demo/global"
	"go-demo/model"
	"go-demo/services"
	"net/http"

	"go.uber.org/zap"
)

func GetUsers(c *gin.Context) {
	result := global.NewResult(c)
	UsersInfo := services.GetUsers()
	result.Success(UsersInfo)
	return
}

func SaveUser(c *gin.Context) {
	result := global.NewResult(c)
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.S().Error(err)
		result.Error(http.StatusNotFound, "参数错误")
		return
	}
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
