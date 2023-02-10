package model

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	return new(Response)
}

func (m *Response) WithMessage(message string) *Response {
	m.Message = message
	return m
}

func (m *Response) WithData(data any) *Response {
	m.Data = data
	return m
}
