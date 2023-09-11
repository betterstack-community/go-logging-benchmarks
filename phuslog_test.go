package bench

import (
	"io"
	"time"

	"github.com/phuslu/log"
)

func (u user) MarshalObject(e *log.Entry) {
	e.Str("name", u.Name).
		Int("age", u.Age).
		Time("dob", u.DOB)
}

func phusFields(e *log.Entry) *log.Entry {
	e.
		Int("bytes", ctxBodyBytes).
		Str("request", ctxRequest).
		Float64("elapsed_time_ms", ctxTimeElapsedMs).
		Object("user", ctxUser).
		Time("now", ctxTime).
		Strs("months", ctxMonths).
		Ints("primes", ctxFirst10Primes).
		Any("users", ctxUsers).
		Err(ctxErr)

	return e
}

func newPhusLog(w io.Writer) log.Logger {
	l := log.Logger{
		Level:      log.InfoLevel,
		Caller:     0,
		TimeField:  "time",
		TimeFormat: time.RFC3339Nano,
		Writer:     &log.IOWriter{w},
	}

	return l
}

type phusLogBench struct {
	l log.Logger
}

func (b *phusLogBench) new(w io.Writer) logBenchmark {
	return &phusLogBench{
		l: newPhusLog(w),
	}
}

func (b *phusLogBench) newWithCtx(w io.Writer) logBenchmark {
	l := newPhusLog(w)
	l.Context = phusFields(log.NewContext(nil)).Value()

	return &phusLogBench{
		l,
	}
}

func (b *phusLogBench) name() string {
	return "Phuslog"
}

func (b *phusLogBench) logEventOnly(msg string) {
	b.l.Info().Msg(msg)
}

func (b *phusLogBench) logWithCtx(msg string) {
	phusFields(b.l.Info()).Msg(msg)
}

func (b *phusLogBench) logDisabled(msg string) {
	b.l.Debug().Msg(msg)
}

func (b *phusLogBench) logDisabledWithCtx(msg string) {
	phusFields(b.l.Debug()).Msg(msg)
}
