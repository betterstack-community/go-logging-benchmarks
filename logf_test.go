package bench

import (
	"fmt"
	"io"
	"time"

	"github.com/zerodha/logf"
)

type logfBench struct {
	l logf.Logger
}

func newLogf(w io.Writer) logf.Logger {
	l := logf.New(logf.Opts{
		Writer:          w,
		Level:           logf.InfoLevel,
		TimestampFormat: time.RFC3339Nano,
	})

	return l
}

func (b *logfBench) new(w io.Writer) logBenchmark {
	return &logfBench{
		l: newLogf(w),
	}
}

func (b *logfBench) newWithCtx(w io.Writer) logBenchmark {
	l := newLogf(w)
	l.DefaultFields = alternatingKeyValuePairs()

	return &logfBench{
		l,
	}
}

func (b *logfBench) name() string {
	return "Logf"
}

func (b *logfBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *logfBench) logEventFmt(msg string, args ...any) {
	b.l.Info(fmt.Sprintf(msg, args...))
}

func (b *logfBench) logEventCtx(msg string) {
	b.l.Info(msg, alternatingKeyValuePairs()...)
}

func (b *logfBench) logEventCtxWeak(msg string) {
	b.logEventCtx(msg)
}

func (b *logfBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *logfBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debug(fmt.Sprintf(msg, args...))
}

func (b *logfBench) logDisabledCtx(msg string) {
	b.l.Debug(msg, alternatingKeyValuePairs()...)
}

func (b *logfBench) logDisabledCtxWeak(msg string) {
	b.logDisabledCtx(msg)
}
