package bench

import (
	"io"
	"log/slog"

	"github.com/rs/zerolog"
)

type slogZerologBench struct {
	slogBench
}

func newSlogZerolog(w io.Writer) *slog.Logger {
	l := newZerolog(w)
	h := zerolog.NewSlogHandler(l)
	return slog.New(h)
}

func newSlogZerologWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	l := newZerolog(w)
	h := zerolog.NewSlogHandler(l).WithAttrs(attr)
	return slog.New(h)
}

func (b *slogZerologBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogZerolog(w),
	}
}

func (b *slogZerologBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogZerologWithCtx(w, slogAttrs()),
	}
}

func (b *slogZerologBench) name() string {
	return "SlogZerolog"
}
