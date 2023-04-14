package response

import (
	"github.com/Joword/chatbgi-manager/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
