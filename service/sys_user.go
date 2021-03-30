package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"goblog/config"
	"goblog/model/ent"
	entUser "goblog/model/ent/user"
)

func Register(ctx context.Context, u *ent.User) (*ent.User, error) {
	hMac := hmac.New(sha512.New, []byte(config.ServerConfig.System.SecretKey))
	hMac.Write([]byte(u.Password))
	user, err := Client.User.Create().
		SetUserName(u.UserName).
		SetNickName(u.NickName).
		SetPassword(hex.EncodeToString(hMac.Sum(nil))).
		SetHeaderImg(u.HeaderImg).
		Save(ctx)
	if err != nil {
		log.Err(err).Stack().Msg("创建用户错误")
	}
	return user, err
}

func Login(ctx context.Context, userName, passwd string) (*ent.User, error) {
	hMac := hmac.New(sha512.New, []byte(config.ServerConfig.System.SecretKey))
	hMac.Write([]byte(passwd))
	user, err := Client.User.Query().Where(
		entUser.And(
			entUser.UserName(userName),
			entUser.Password(hex.EncodeToString(hMac.Sum(nil))),
		),
	).First(ctx)
	if err != nil {
		log.Err(err).Stack().Msg("登录验证错误")
	}
	return user, err
}
