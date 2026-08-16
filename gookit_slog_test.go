package bench

import (
	"io"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

func newGookitSlog(w io.Writer) *slog.Logger {
	h := handler.NewSimpleHandler(w, slog.InfoLevel)
	h.SetFormatter(slog.NewJSONFormatter())

	l := slog.NewWithHandlers(h)
	l.DoNothingOnPanicFatal()
	l.CallerFlag = 0

	return l
}

func gookitSlogFields() slog.M {
	return slog.M{
		"bytes":           ctxBodyBytes,
		"request":         ctxRequest,
		"elapsed_time_ms": ctxTimeElapsedMs,
		"user":            ctxUser,
		"now":             ctxTime,
		"months":          ctxMonths,
		"primes":          ctxFirst10Primes,
		"users":           ctxUsers,
		"error":           ctxErr,
	}
}

type gookitSlogBench struct {
	l   *slog.Logger
	sub *slog.SubLogger
}

func (b *gookitSlogBench) new(w io.Writer) logBenchmark {
	return &gookitSlogBench{
		l: newGookitSlog(w),
	}
}

func (b *gookitSlogBench) newWithCtx(w io.Writer) logBenchmark {
	l := newGookitSlog(w)
	sub := l.NewSub().KeepFields(gookitSlogFields())

	return &gookitSlogBench{
		l:   l,
		sub: sub,
	}
}

func (b *gookitSlogBench) name() string {
	return "GookitSlog"
}

func (b *gookitSlogBench) logEvent(msg string) {
	if b.sub != nil {
		b.sub.Info(msg)
		return
	}
	b.l.Info(msg)
}

func (b *gookitSlogBench) logEventFmt(msg string, args ...any) {
	if b.sub != nil {
		b.sub.Infof(msg, args...)
		return
	}
	b.l.Infof(msg, args...)
}

func (b *gookitSlogBench) logEventCtx(msg string) {
	b.l.WithFields(gookitSlogFields()).Info(msg)
}

func (b *gookitSlogBench) logEventCtxWeak(msg string) {
	b.l.WithFields(gookitSlogFields()).Info(msg)
}

func (b *gookitSlogBench) logDisabled(msg string) {
	if b.sub != nil {
		b.sub.Debug(msg)
		return
	}
	b.l.Debug(msg)
}

func (b *gookitSlogBench) logDisabledFmt(msg string, args ...any) {
	if b.sub != nil {
		b.sub.Debugf(msg, args...)
		return
	}
	b.l.Debugf(msg, args...)
}

func (b *gookitSlogBench) logDisabledCtx(msg string) {
	b.l.WithFields(gookitSlogFields()).Debug(msg)
}

func (b *gookitSlogBench) logDisabledCtxWeak(msg string) {
	b.l.WithFields(gookitSlogFields()).Debug(msg)
}
