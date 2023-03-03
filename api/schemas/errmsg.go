package schemas

type ErrMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Resp struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
