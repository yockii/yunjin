package domain

import "github.com/yockii/yunjin/internal/model"

type LoginRespData struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}
