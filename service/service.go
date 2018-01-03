package service

import (
	"context"
	"errors"
	"strings"
)

var ErrEmpty = errors.New("Empty string")

type ServiceInterface interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type Service struct{}

func (Service) Uppercase(_ context.Context, input string) (output string, err error) {
	if input == "" {
		output = ""
		err = ErrEmpty
	} else {
		output = strings.ToUpper(input)
	}

	return
}

func (Service) Count(_ context.Context, input string) (output int) {
	return len(input)
}
