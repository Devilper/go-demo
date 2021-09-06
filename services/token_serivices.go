package services

import (
	"go-demo/common"
	"go-demo/global"
	"time"
)

type TokenService interface {
	GetToken(id string) error
}

func GetToken(id string) (string, error) {

	//获取uuid
	uuid := common.GetUUid()
	// 设置redis token
	pl := global.Redis.Pipeline()
	pl.Set(id, uuid, time.Second*200)
	if _, err := pl.Exec(); err != nil {
		return "", err
	}
	return uuid, nil
}
