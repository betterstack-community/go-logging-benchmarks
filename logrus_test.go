package bench

import (
	"io"

	"github.com/sirupsen/logrus"
)

func logrusFields() logrus.Fields {
	return logrus.Fields{
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

func newLogrus(w io.Writer) *logrus.Logger {
	l := logrus.New()
	l.Out = w
	l.Level = logrus.InfoLevel
	l.SetFormatter(&logrus.JSONFormatter{})

	return l
}

type logrusBench struct {
	l *logrus.Logger
}

func (b *logrusBench) new(w io.Writer) logBenchmark {
	return &logrusBench{
		l: newLogrus(w),
	}
}

func (b *logrusBench) newWithCtx(w io.Writer) logBenchmark {
	return &logrusBench{
		l: newLogrus(w).WithFields(logrusFields()).Logger,
	}
}

func (b *logrusBench) name() string {
	return "Logrus"
}

func (b *logrusBench) logEventOnly(msg string) {
	b.l.Info(msg)
}

func (b *logrusBench) logWithCtx(msg string) {
	b.l.WithFields(logrusFields()).Info(msg)
}

func (b *logrusBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *logrusBench) logDisabledWithCtx(msg string) {
	b.l.WithFields(logrusFields()).Debug(msg)
}
