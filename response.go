package openrouter

import (
	"reflect"
)

type apiResponse interface {
	*ChatCompletionResponse
}

type ApiResponse[T apiResponse] struct {
	Data T `json:"data"`
}

func (a ApiResponse[T]) IsNil() bool {
	return reflect.ValueOf(a.Data).IsNil()
}
