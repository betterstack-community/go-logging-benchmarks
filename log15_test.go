package bench

import (
	"io"

	"github.com/inconshreveable/log15"
)

type log15Bench struct {
	l log15.Logger
}

func newLog15(w io.Writer) log15.Logger {
	l := log15.New()
	h := log15.StreamHandler(w, log15.JsonFormat())
	l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, h))

	return l
}

func (b *log15Bench) new(w io.Writer) logBenchmark {
	return &log15Bench{
		l: newLog15(w),
	}
}

func (b *log15Bench) newWithCtx(w io.Writer) logBenchmark {
	return &log15Bench{
		l: newLog15(w).New(alternatingKeyValuePairs()...),
	}
}

func (b *log15Bench) name() string {
	return "Log15"
}

func (b *log15Bench) logEventOnly(msg string) {
	b.l.Info(msg)
}

func (b *log15Bench) logWithCtx(msg string) {
	b.l.Info(msg, alternatingKeyValuePairs()...)
}

func (b *log15Bench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *log15Bench) logDisabledWithCtx(msg string) {
	b.l.Debug(msg, alternatingKeyValuePairs()...)
}
