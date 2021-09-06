package services

import (
	"errors"
	"fmt"
	"go-demo/global"
	"go-demo/model"
	"go.uber.org/zap"
	"strconv"
)

type UserService interface {
	GetUsers() []interface{}
	GetUser(id string) map[string]interface{}
	SaveUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
}

func GetUsers() []interface{} {
	var users []model.User
	var UsersInfos []interface{}
	if err := global.Db.Find(&users).Error; err != nil {
		return UsersInfos
	}
	for _, v := range users {
		UserMap := make(map[string]interface{})
		UserMap["user_name"] = v.UserName
		UserMap["phone_num"] = v.PhoneNum
		UserMap["id"] = v.ID
		UsersInfos = append(UsersInfos, UserMap)
	}
	return UsersInfos
}
func SaveUser(user *model.User) error {
	if err := global.Db.Create(user).Error; err != nil {
		zap.S().Error(err)
		return err
	}
	return nil
}

func GetUser(id string) map[string]interface{} {
	var user model.User
	UserInfo := make(map[string]interface{})
	if err := global.Db.Where("ID=?", id).First(&user).Error; err != nil {
		return UserInfo
	}
	UserInfo["id"] = user.ID
	UserInfo["user_name"] = user.UserName
	UserInfo["phone_num"] = user.PhoneNum

	return UserInfo
}

func UpdateUser(user *model.User) error {
	if err := global.Db.Model(&user).
		Where("ID=?", user.ID).
		Updates(map[string]interface{}{"UserName": user.UserName,
			"PhoneNum": user.PhoneNum}).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteUser(id string) error {
	var user model.User
	if err := global.Db.Where("ID=?", id).Take(&user).Error; err != nil {
		return err
	}
	global.Db.Delete(&user)
	return nil
}

func UserLogin(json map[string]interface{}) (string, error) {
	var Users []model.User
	UserName := json["user_name"]
	if err := global.Db.Where("user_name=?", UserName).First(&Users).Error; err != nil {
		return "", err
	}
	fmt.Println(len(Users))
	if len(Users) == 1 {
		User := Users[0]
		if User.Password == json["password"] {
			ID := User.ID
			id := strconv.Itoa(int(ID))
			// 生成token
			token, err := GetToken(id)
			if err != nil {
				return "", err
			} else {
				return token, nil
			}
		}
	}
	return "", errors.New("密码错误")

}
