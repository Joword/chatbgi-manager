package request

import (
	"github.com/Joword/chatbgi-manager/model/common/request"
	"github.com/Joword/chatbgi-manager/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
