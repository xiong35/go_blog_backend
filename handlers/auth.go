package handlers

import (
	"go_blog/database"
	"go_blog/model"
	"time"
)

type Token struct {
	Password string `form:"password"`
	Token    string `form:"token"`
}

// Check :check for authorization
func (token *Token) Check() (rtToken string, permission int32, new bool, err error) {
	var authModel model.Auth

	if token.Token != "" {
		oneDayAgo, _ := time.ParseDuration("-24h")
		yesterday := time.Now().Add(oneDayAgo)
		database.DB.Table("auths").
			Where("token = ? AND created_at > ? ", token.Token, yesterday).
			First(&authModel)
	} else if token.Password == "" {
		return
	}
	if authModel.Permission != 0 {
		rtToken = token.Token
		permission = authModel.Permission
		return
	}

	authModel.Token = authModel.GenToken()
	authModel.Permission = 1
	if authModel.IsValidPassord(token.Password) {
		authModel.Permission = 10
	}

	permission = authModel.Permission
	rtToken = authModel.Token
	new = true

	result := database.DB.Create(&authModel)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
