package bench

import (
	"io"
	"log/slog"

	phuslog "github.com/phuslu/log"
)

type slogPhuslogBench struct {
	slogBench
}

// slog frontend with phuslog backend.
func newSlogPhuslog(w io.Writer) *slog.Logger {
	return slog.New(phuslog.SlogNewJSONHandler(w, nil))
}

func newSlogPhuslogWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	return slog.New(phuslog.SlogNewJSONHandler(w, nil).WithAttrs(attr))
}

func (b *slogPhuslogBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogPhuslog(w),
	}
}

func (b *slogPhuslogBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogPhuslogWithCtx(w, slogAttrs()),
	}
}

func (b *slogPhuslogBench) name() string {
	return "SlogPhuslog"
}
