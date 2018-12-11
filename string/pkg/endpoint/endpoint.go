package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/makersu/go-kit-example-string/string/pkg/service"
)

// UppercaseRequest collects the request parameters for the Uppercase method.
type UppercaseRequest struct {
	S string `json:"s"`
}

// UppercaseResponse collects the response parameters for the Uppercase method.
type UppercaseResponse struct {
	Uppercase string `json:"uppercase"`
	Err       error  `json:"err"`
}

// MakeUppercaseEndpoint returns an endpoint that invokes Uppercase on the service.
func MakeUppercaseEndpoint(s service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		uppercase, err := s.Uppercase(ctx, req.S)
		return UppercaseResponse{
			Err:       err,
			Uppercase: uppercase,
		}, nil
	}
}

// Failed implements Failer.
func (r UppercaseResponse) Failed() error {
	return r.Err
}

// CountRequest collects the request parameters for the Count method.
type CountRequest struct {
	S string `json:"s"`
}

// CountResponse collects the response parameters for the Count method.
type CountResponse struct {
	Count int `json:"count"`
}

// MakeCountEndpoint returns an endpoint that invokes Count on the service.
func MakeCountEndpoint(s service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		count := s.Count(ctx, req.S)
		return CountResponse{Count: count}, nil
	}
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Uppercase implements Service. Primarily useful in a client.
func (e Endpoints) Uppercase(ctx context.Context, s string) (uppercase string, err error) {
	request := UppercaseRequest{S: s}
	response, err := e.UppercaseEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UppercaseResponse).Uppercase, response.(UppercaseResponse).Err
}

// Count implements Service. Primarily useful in a client.
func (e Endpoints) Count(ctx context.Context, s string) (count int) {
	request := CountRequest{S: s}
	response, err := e.CountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CountResponse).Count
}
