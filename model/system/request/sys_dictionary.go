package request

import (
	"github.com/Joword/chatbgi-manager/model/common/request"
	"github.com/Joword/chatbgi-manager/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}