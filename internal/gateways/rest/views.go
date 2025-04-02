package rest

type R struct {
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
}
