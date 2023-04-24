package auth_service

import "github.com/Joword/chatbgi-manager/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() []models.UserMessage {
	return models.CheckAuth(a.Username, a.Password)
}
