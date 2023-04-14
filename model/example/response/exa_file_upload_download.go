package response

import "github.com/Joword/chatbgi-manager/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
