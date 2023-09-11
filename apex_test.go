package bench

import (
	"io"

	apex "github.com/apex/log"
	"github.com/apex/log/handlers/json"
)

func apexFields() apex.Fields {
	return apex.Fields{
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

func newApex(w io.Writer) *apex.Logger {
	writer := json.New(w)

	return &apex.Logger{
		Handler: writer,
		Level:   apex.InfoLevel,
	}
}

type apexBench struct {
	l apex.Interface
}

func (b *apexBench) new(w io.Writer) logBenchmark {
	return &apexBench{
		l: newApex(w),
	}
}

func (b *apexBench) newWithCtx(w io.Writer) logBenchmark {
	return &apexBench{
		l: newApex(w).WithFields(apexFields()),
	}
}

func (b *apexBench) name() string {
	return "Apex"
}

func (b *apexBench) logEventOnly(msg string) {
	b.l.Info(msg)
}

func (b *apexBench) logWithCtx(msg string) {
	b.l.WithFields(apexFields()).Info(msg)
}

func (b *apexBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *apexBench) logDisabledWithCtx(msg string) {
	b.l.WithFields(apexFields()).Debug(msg)
}
