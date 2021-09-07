package services

import (
	"github.com/dgrijalva/jwt-go"
	myjwt "go-demo/middleware/jwt"
	"go-demo/model"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type LoginResult struct {
	Token string `json:"token"`
	model.User
}
type TokenService interface {
	GenerateToken(user model.User) error
}

func GenerateToken(user model.User) (LoginResult, error) {
	j := &myjwt.JWT{
		SigningKey: []byte("devil"),
	}
	data := LoginResult{
		User:  user,
		Token: "",
	}
	claims := myjwt.CustomClaims{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.UserName,
		Phone: user.Password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), //过期时间 一小时
			Issuer:    "devil",                         //签名的发行者
		},
	}
	token, err := j.CreatToken(claims)
	if err != nil {
		return data, err
	}
	zap.S().Info(token)
	data.Token = token
	return data, nil
}
