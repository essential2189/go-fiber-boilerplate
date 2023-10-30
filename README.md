# go-boilerplate v1.0.0

go version - 1.2.1

### Must Changed Before Use

- The `module` name in `go.mod` file
- The `AppName` in [app.go](app/app.go) file
- The `historyType` and `version` in [logger.go](app/core/helper/logger/logger.go)

### How to Run

#### Install required packages

- `go mod tidy` or `go mod vendor`
- `go install github.com/cosmtrek/air@latest`
- `go install github.com/swaggo/swag/cmd/swag@latest`

#### Run

env : `APP_HOME`, `DEBUG`, `TEST`, `wavve_env`, `wavve_port`

- `wavve_env=dev air` - Run the server with hot reload
- `swag init` - Generate swagger docs
- `docker build -t go-boilerplate .` - Build docker image
- `docker run -e "wavve_env=dev" -p 8080:8080 go-boilerplate` - Run docker image

### Library

- Fiber - https://github.com/gofiber/fiber
- Validator - https://github.com/go-playground/validator
- DI - https://github.com/uber-go/fx
- Logger - https://github.com/uber-go/zap
- Errors - https://github.com/cockroachdb/errors

### Resty Example

```go
url := config.Apis + consts.GetProductDetailUrl + strconv.Itoa(paymentID)

requestInfo := resty.RequestInfo{
    Uri:          url,
    Method:       resty.MethodGET,
    Headers:      map[string]string{"Content-Type": "application/json"},
    Query:        map[string]string{},
    Body:         nil,
    Timeout:      0,
    RetryCount:   0,
    RetryBackOff: 0,
    IsSkipSSL:    false,
}

res, err := svc.resty.Request(requestInfo)
if err != nil {
    return errors.Wrap(err, "failed to get product detail")
}

err = json.Unmarshal(res.Body, &productDetail)
if err != nil {
    return errors.Wrap(err, "failed to unmarshal message")
}
```
