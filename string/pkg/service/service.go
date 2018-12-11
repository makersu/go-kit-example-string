package service

import (
	"context"
	"errors"
	"strings"
)

// StringService describes the service.
type StringService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Uppercase(ctx context.Context, s string) (uppercase string, err error)
	Count(ctx context.Context, s string) (count int)
}

type basicStringService struct{}

// TODO implement the business logic of Uppercase
func (b *basicStringService) Uppercase(ctx context.Context, s string) (uppercase string, err error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return strings.ToUpper(s), nil
}

// TODO implement the business logic of Count
func (b *basicStringService) Count(ctx context.Context, s string) (count int) {
	return len(s)
}

// NewBasicStringService returns a naive, stateless implementation of StringService.
func NewBasicStringService() StringService {
	return &basicStringService{}
}

// New returns a StringService with all of the expected middleware wired in.
func New(middleware []Middleware) StringService {
	var svc StringService = NewBasicStringService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
