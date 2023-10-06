package bench

import (
	"io"

	"github.com/sirupsen/logrus"
)

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
		l: newLogrus(w).WithFields(mapFields()).Logger,
	}
}

func (b *logrusBench) name() string {
	return "Logrus"
}

func (b *logrusBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *logrusBench) logEventFmt(msg string, args ...any) {
	b.l.Infof(msg, args...)
}

func (b *logrusBench) logEventCtx(msg string) {
	b.l.WithFields(mapFields()).Info(msg)
}

func (b *logrusBench) logEventCtxWeak(msg string) {
	b.logEventCtx(msg)
}

func (b *logrusBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *logrusBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debugf(msg, args...)
}

func (b *logrusBench) logDisabledCtx(msg string) {
	b.l.WithFields(mapFields()).Debug(msg)
}

func (b *logrusBench) logDisabledCtxWeak(msg string) {
	b.logDisabledCtx(msg)
}
