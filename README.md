# go-kit-example-string

## PartI. Create Stringsvc1

## Test Stringsvc1
```
github.com/makersu/go-kit-example-string/stringsvc1 > go run main.go

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO, WORLD"}
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}
```

## PartI. Create Stringsvc2
## Test Stringsvc2
```
~/go/src/github.com/makersu/go-kit-example-string/stringsvc2> ./stringsvc2 

```

## Install kit
```
go-kit-example-hello> go get github.com/go-kit/kit
go-kit-example-hello> go get github.com/kujtimiihoxha/kit
```

## Create a new service by kit
```
# kit new service string
> kit n s string
```

## Define service (string/pkg/service/service.go)
```
type StringService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Uppercase(ctx context.Context, s string) (uppercase string, err error)
	Count(ctx context.Context, s string) (count int)
}
```

## Generate service
```
# kit generate service string
> kit g s string
```

## Implement service
```
// TODO implement the business logic of Uppercase
func (b *basicStringService) Uppercase(ctx context.Context, s string) (uppercase string, err error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return strings.ToUpper(s), nil
}
```

```
// TODO implement the business logic of Count
func (b *basicStringService) Count(ctx context.Context, s string) (count int) {
	return len(s)
}
```

## Run service
```
> go run string/cmd/main.go

ts=2018-12-11T15:36:27.211584Z caller=service.go:77 tracer=none
ts=2018-12-11T15:36:27.211952Z caller=service.go:99 transport=HTTP addr=:8081
ts=2018-12-11T15:36:27.211975Z caller=service.go:125 transport=debug/HTTP addr=:8080
```

## Test service
```
> curl -XPOST -d'{"s":"hello, world"}' localhost:8081/uppercase
{"uppercase":"HELLO, WORLD","err":null}

> curl -XPOST -d'{"s":"hello, world"}' localhost:8081/count
{"count":12}
```

```
.
├── README.md
└── string
    ├── cmd
    │   ├── main.go
    │   └── service
    │       ├── service.go
    │       └── service_gen.go
    └── pkg
        ├── endpoint
        │   ├── endpoint.go
        │   └── endpoint_gen.go
        ├── http
        │   ├── handler.go
        │   └── handler_gen.go
        └── service
            ├── middleware.go
            └── service.go
```