package hooks

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/twitchtv/twirp"
)

var ctxKey = new(int)

// LoggingHooks creates a new twirp.ServerHooks which logs requests as they are
// routed to Twirp, and logs responses (including response time) when they are
// sent.
//
// This is a demonstration showing how you can use context accessors with hooks.
func LoggingHooks(w io.Writer) *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			startTime := time.Now()
			ctx = context.WithValue(ctx, ctxKey, startTime)
			return ctx, nil
		},
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			svc, _ := twirp.ServiceName(ctx)
			meth, _ := twirp.MethodName(ctx)
			fmt.Fprintf(w, "received req svc=%q method=%q\n", svc, meth)
			return ctx, nil
		},
		ResponseSent: func(ctx context.Context) {
			startTime := ctx.Value(ctxKey).(time.Time)
			svc, _ := twirp.ServiceName(ctx)
			meth, _ := twirp.MethodName(ctx)
			fmt.Fprintf(w, "response sent svc=%q method=%q time=%q\n", svc, meth, time.Since(startTime))
		},
	}
}
