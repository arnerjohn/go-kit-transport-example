package main

import (
	"context"
	"errors"
	"fmt"
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

func main() {
	var svc ServiceInterface
	svc = Service{}

	upperOut, err := svc.Uppercase(context.Background(), "hello")

	if err != nil {
		panic(err)
	}

	fmt.Println(upperOut)

	countOut := svc.Count(context.Background(), "hello")

	fmt.Println(countOut)
}
