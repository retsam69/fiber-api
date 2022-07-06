package api_response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type SetDefaulter interface {
	Default()
}

type APIResponse struct {
	OK     bool   `json:"ok"`     // Response is Error
	Msg    string `json:"msg"`    // Success Message
	Result any    `json:"result"` // Result Data
} // @name APISuccess
func NewAPIResponse() *APIResponse {
	res := APIResponse{}
	res.Default()
	return &res
}
func (a *APIResponse) Default() {
	a.Msg = ""
	a.OK = true
	a.Result = nil
}

type APIError struct {
	APIResponse
	Detail interface{} `json:"detail,omitempty"` // Eror Detail or ETC.
} // @name APIError

func NewAPIError() *APIError {
	res := APIError{}
	res.Default()
	return &res
}

func (a *APIError) Default() {
	a.APIResponse.Default()
	a.OK = false
}

// * Interface request for set fiber.ErrorHendler
func (a *APIError) SetError(c *fiber.Ctx, code int, err error) {
	if err == nil {
		return
	}
	a.APIResponse.Default()
	a.OK = false
	a.Msg = err.Error()
	a.Detail = fmt.Sprintf("Error code: %d, Message: %s", code, err.Error())
}
