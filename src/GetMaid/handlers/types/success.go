package types

type Success struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg,omitempty"`
}

type TokenUnverified struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg,omitempty"`
	Secret  string      `json:"secret,omitempty"`
}
