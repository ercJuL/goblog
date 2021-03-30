package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"goblog/config"
	"goblog/model/ent"
)

func Register(ctx context.Context, u *ent.User) (*ent.User, error) {
	hmac := hmac.New(sha512.New, []byte(config.ServerConfig.System.SecretKey))
	hmac.Write([]byte(u.Password))
	user, err := Client.User.Create().
		SetUserName(u.UserName).
		SetNickName(u.NickName).
		SetPassword(hex.EncodeToString(hmac.Sum(nil))).
		SetHeaderImg(u.HeaderImg).
		Save(ctx)
	if err != nil {
		log.Err(err).Stack().Msg("创建用户错误")
	}
	return user, err
}
