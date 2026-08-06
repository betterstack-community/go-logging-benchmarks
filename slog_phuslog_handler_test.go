package bench

import (
	"io"
	"log/slog"

	"github.com/phuslu/log"
)

type slogPhuslogHandlerBench struct {
	slogBench
}

func newSlogPhuslogHandler(w io.Writer) *slog.Logger {
	h := log.SlogNewJSONHandler(w, &slog.HandlerOptions{})
	return slog.New(h)
}

func newSlogPhuslogHandlerWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	h := log.SlogNewJSONHandler(w, &slog.HandlerOptions{}).WithAttrs(attr)
	return slog.New(h)
}

func (b *slogPhuslogHandlerBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogPhuslogHandler(w),
	}
}

func (b *slogPhuslogHandlerBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogPhuslogHandlerWithCtx(w, slogAttrs()),
	}
}

func (b *slogPhuslogHandlerBench) name() string {
	return "SlogPhuslogHandler"
}
