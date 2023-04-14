package response

import "github.com/Joword/chatbgi-manager/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
