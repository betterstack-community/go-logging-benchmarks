package bench

import (
	"context"
	"io"
	"log/slog"
)

func slogAttrs() []slog.Attr {
	return []slog.Attr{
		slog.Int("bytes", ctxBodyBytes),
		slog.String("request", ctxRequest),
		slog.Float64("elapsed_time_ms", ctxTimeElapsedMs),
		slog.Any("user", ctxUser),
		slog.Time("now", ctxTime),
		slog.Any("months", ctxMonths),
		slog.Any("primes", ctxFirst10Primes),
		slog.Any("users", ctxUsers),
		slog.Any("error", ctxErr),
	}
}

func newSlog(w io.Writer) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func newSlogWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}).WithAttrs(attr))
}

type slogBench struct {
	l *slog.Logger
}

func (b *slogBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlog(w),
	}
}

func (b *slogBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogWithCtx(w, slogAttrs()),
	}
}

func (b *slogBench) name() string {
	return "Slog"
}

func (b *slogBench) logEventOnly(msg string) {
	b.l.Info(msg)
}

func (b *slogBench) logWithCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		msg,
		slogAttrs()...,
	)
}

func (b *slogBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *slogBench) logDisabledWithCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelDebug,
		msg,
		slogAttrs()...,
	)
}
