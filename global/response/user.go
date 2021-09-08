package response

import "go-demo/model"

type LoginResult struct {
	Token string `json:"token"`
	model.User
}

type UserResult struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	PhoneNum string `json:"phone_num"`
}
