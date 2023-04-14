package response

import "github.com/Joword/chatbgi-manager/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
