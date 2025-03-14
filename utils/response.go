package utils

type Response[T any] struct {
	Message   string `json:"message"`
	Data      T      `json:"data"`
	IsSuccess bool   `json:"isSuccess"`
}

func NewResponse[T any](message string, data T, isSuccess bool) *Response[T] {
	return &Response[T]{
		Message:   message,
		Data:      data,
		IsSuccess: isSuccess,
	}
}
