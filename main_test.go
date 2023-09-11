package bench

import (
	"errors"
	"io"
	"sync/atomic"
	"time"
)

type user struct {
	DOB  time.Time
	Name string
	Age  int
}

type users []user

var (
	ctxBodyBytes     = 123456789
	ctxRequest       = "GET /icons/ubuntu-logo.png HTTP/1.1"
	ctxTimeElapsedMs = 11.398466
	ctxUser          = user{
		Name: "John Doe",
		Age:  23,
		DOB:  time.Date(2000, 9, 9, 0, 0, 0, 0, time.UTC),
	}
	ctxUsers = users{
		ctxUser,
		ctxUser,
		ctxUser,
		ctxUser,
		ctxUser,
		ctxUser,
		ctxUser,
		ctxUser,
		ctxUser,
		ctxUser,
	}
	ctxTime   = time.Now()
	ctxMonths = []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	ctxFirst10Primes = []int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31}
	ctxErr           = errors.New("failed to open file: /home/dev/new.txt")
)

func alternatingKeyValuePairs() []any {
	return []any{
		"bytes", ctxBodyBytes,
		"request", ctxRequest,
		"elapsed_time_ms", ctxTimeElapsedMs,
		"user", ctxUser,
		"now", ctxTime,
		"months", ctxMonths,
		"primes", ctxFirst10Primes,
		"users", ctxUsers,
		"error", ctxErr,
	}
}

var logMsg = "The quick brown fox jumps over the lazy dog"

var loggers = []logBenchmark{
	&zerologBench{},
	&phusLogBench{},
	&zapBench{},
	&zapSugarBench{},
	&slogBench{},
	&apexBench{},
	&logrusBench{},
	&log15Bench{},
	&logfBench{},
}

type blackhole struct {
	count uint64
}

func (s *blackhole) WriteCount() uint64 {
	return atomic.LoadUint64(&s.count)
}

func (s *blackhole) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.count, 1)
	return len(p), nil
}

type logBenchmark interface {
	new(w io.Writer) logBenchmark
	newWithCtx(w io.Writer) logBenchmark
	name() string
	logEventOnly(msg string)
	logWithCtx(msg string)
	logDisabled(msg string)
	logDisabledWithCtx(msg string)
}
