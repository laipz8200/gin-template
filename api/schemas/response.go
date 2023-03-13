package schemas

type ErrorMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

var _ Response = Object{}
var _ Response = Message("")
var _ Response = List{}

type Response interface {
	ToResponse(int) any
}

type ObjectResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type ListResponse struct {
	Code  int `json:"code"`
	Total int `json:"total"`
	Data  any `json:"data"`
}

type MessageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Object struct {
	Data any
}

// ToResponse implements Response
func (o Object) ToResponse(code int) any {
	return ObjectResponse{
		Code: code,
		Data: o.Data,
	}
}

type Message string

// ToResponse implements Response
func (m Message) ToResponse(code int) any {
	return MessageResponse{
		Code:    code,
		Message: string(m),
	}
}

type List struct {
	Total int
	Data  any
}

// ToResponse implements Response
func (l List) ToResponse(code int) any {
	return ListResponse{
		Code:  code,
		Total: l.Total,
		Data:  l.Data,
	}
}
