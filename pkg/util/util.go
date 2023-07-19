package util

import "github.com/shinhagunn/todo-backend/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.Cfg.SecretKey)
}
