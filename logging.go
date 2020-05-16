package main

import (
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"net/http"
	"time"
)

func LogHttp(log zerolog.Logger) alice.Chain {
	c := alice.New()

	// Install the logger handler with default output on the console
	c = c.Append(hlog.NewHandler(log))

	// Install some provided extra handler to set some request's context fields.
	// Thanks to that handler, all our logs will come with some prepopulated fields.
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))

	return c
}
