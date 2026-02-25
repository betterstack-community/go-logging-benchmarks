package bench

import (
	"io"
	"log/slog"

	"go.uber.org/zap/exp/zapslog"
)

type slogZapBench struct {
	slogBench
}

// slog frontend with Zap backend.
func newSlogZap(w io.Writer) *slog.Logger {
	l := newZap(w)

	return slog.New(zapslog.NewHandler(l.Core()))
}

func newSlogZapWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	l := newZap(w)

	h := zapslog.NewHandler(l.Core()).WithAttrs(attr)

	return slog.New(h)
}

func (b *slogZapBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogZap(w),
	}
}

func (b *slogZapBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogZapWithCtx(w, slogAttrs()),
	}
}

func (b *slogZapBench) name() string {
	return "SlogZap"
}
