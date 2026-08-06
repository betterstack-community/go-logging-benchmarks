package bench

import (
	"io"
	"log/slog"
)

type slogPhuslogLoggerBench struct {
	slogBench
}

func newSlogPhuslogLogger(w io.Writer) *slog.Logger {
	l := newPhusLog(w)
	return l.Slog()
}

func newSlogPhuslogLoggerWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	l := newPhusLog(w)
	return slog.New(l.Slog().Handler().WithAttrs(attr))
}

func (b *slogPhuslogLoggerBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogPhuslogLogger(w),
	}
}

func (b *slogPhuslogLoggerBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogPhuslogLoggerWithCtx(w, slogAttrs()),
	}
}

func (b *slogPhuslogLoggerBench) name() string {
	return "SlogPhuslogLoggerWrapper"
}
