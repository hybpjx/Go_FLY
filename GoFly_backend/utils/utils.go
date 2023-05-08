package utils

import "fmt"

// AppendError 判断是否有错误 如果 无错，则把新的错抛出去
func AppendError(existError, newError error) error {
	if existError == nil {
		return newError
	}
	return fmt.Errorf("%v;%w", existError, newError)
}
