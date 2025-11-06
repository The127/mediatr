package mediatr

import (
	"context"
	"log/slog"
)

type loggerKey struct{}

func WithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

func Logger(ctx context.Context) *slog.Logger {
	l, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if ok && l != nil {
		return l
	}
	return slog.Default()
}
