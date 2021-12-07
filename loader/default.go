package loader

var CONFIG_INI_DEFAULT string = `
[app]
dev=1
baseurl=https://localhost:8888
listen=127.0.0.1
[logger]
log=/logs/log.log
error=./logs/error.log
`

type APIResponse struct {
	IsError bool   `json:"error"` // Response is Error
	Msg     string `json:"msg"`   // Success Message
} // @name APISuccess

type APIError struct {
	APIResponse
	Detail interface{} `json:"detail,omitempty"` // Eror Detail or ETC.
} // @name APIError
