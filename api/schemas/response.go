package schemas

type ErrorMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
