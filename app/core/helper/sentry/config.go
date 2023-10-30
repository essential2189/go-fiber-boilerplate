package sentry

import (
	"go-boilerplate/app/core/consts"
	"go-boilerplate/config"
	"net/http"
	"net/url"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func setScopeAndCapture(ctx *fiber.Ctx, err error, stackTrace string) *sentry.EventID {
	requestId := ctx.Locals("request_id").(string)
	var id *sentry.EventID
	h := sentry.CurrentHub().Clone()
	h.WithScope(func(scope *sentry.Scope) {
		scope = h.Scope()
		req := convertRequest(ctx.Request())
		scope.SetRequest(req)
		scope.SetRequestBody(ctx.Body())
		scope.SetTag("request_id", requestId)
		scope.AddEventProcessor(func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			event.Exception[0].Value = err.Error()
			if len(stackTrace) == 0 {
				event.Exception[0].Stacktrace = sentry.ExtractStacktrace(errors.Unwrap(err))
			} else {
				event.Exception[0].Stacktrace = sentry.ExtractStacktrace(errors.New(stackTrace))
			}
			event.Exception[0].Type = errors.Cause(err).Error()
			return event
		})

		id = h.CaptureException(errors.Cause(err))
	})
	return id
}

func getOptions(dsn string) sentry.ClientOptions {
	commitSha := os.Getenv(consts.CommitSHA)

	return sentry.ClientOptions{
		Dsn:              dsn,
		Environment:      config.Profile,
		Release:          commitSha,
		AttachStacktrace: true,
		TracesSampleRate: 0.5,
	}
}

func convertRequest(req *fasthttp.Request) *http.Request {
	httpReq := new(http.Request)

	httpReq.Method = string(req.Header.Method())

	httpReq.URL = &url.URL{
		Scheme: string(req.URI().Scheme()),
		Host:   string(req.URI().Host()),
		Path:   string(req.URI().Path()),
	}

	httpReq.Header = make(http.Header)
	req.Header.VisitAll(func(key, value []byte) {
		httpReq.Header.Add(string(key), string(value))
	})

	httpReq.URL.RawQuery = string(req.URI().QueryString())

	return httpReq
}
