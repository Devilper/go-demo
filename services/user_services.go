package services

import (
	"errors"
	"go-demo/common"
	"go-demo/forms"
	"go-demo/global"
	"go-demo/model"
	"go.uber.org/zap"

	uresponse "go-demo/global/response"
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

func GetUser(id string) uresponse.UserResult {
	var user model.User
	var UserInfo = uresponse.UserResult{}
	if err := global.Db.Where("ID=?", id).First(&user).Error; err != nil {
		return UserInfo
	}
	UserInfo = uresponse.UserResult{
		ID:       user.ID,
		UserName: user.UserName,
		PhoneNum: user.PhoneNum,
	}
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

func UserLogin(json *forms.UserLoginRequest) (model.User, error) {
	var Users []model.User
	var user model.User
	UserName := json.UserName
	if err := global.Db.Where("user_name=?", UserName).Take(&Users).Error; err != nil {
		return user, err
	}
	if len(Users) == 1 {
		user = Users[0]
		isTrue := common.DecryptPassword(user.Password, json.Password)
		if isTrue {
			return user, nil
		}
	}
	return user, errors.New("????????????")

}
