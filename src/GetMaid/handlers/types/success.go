package types

type Success struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg, omitempty"`
}
