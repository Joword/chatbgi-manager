package system

import (
	"github.com/Joword/chatbgi-manager/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
