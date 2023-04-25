package auth_service

import "github.com/Joword/chatbgi-manager/models"

type Auth struct {
	Username    string
	Password    string
	IsSuperuser int
}

func (a *Auth) Check() bool {
	a.IsSuperuser = 1
	return models.CheckAuth(a.Username, a.IsSuperuser)
}
