package request

import (
	"github.com/Joword/chatbgi-manager/model/common/request"
	"github.com/Joword/chatbgi-manager/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
