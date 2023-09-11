package bench

import (
	"io"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (u user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.Name)
	enc.AddInt("age", u.Age)
	enc.AddTime("dob", u.DOB)

	return nil
}

func (uu users) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	var err error
	for i := range uu {
		err = multierr.Append(err, arr.AppendObject(uu[i]))
	}

	return err
}

func zapFields() []zap.Field {
	return []zap.Field{
		zap.Int("bytes", ctxBodyBytes),
		zap.String("request", ctxRequest),
		zap.Float64("elapsed_time_ms", ctxTimeElapsedMs),
		zap.Object("user", ctxUser),
		zap.Time("now", ctxTime),
		zap.Strings("months", ctxMonths),
		zap.Ints("primes", ctxFirst10Primes),
		zap.Array("users", ctxUsers),
		zap.Error(ctxErr),
	}
}

func newZap(w io.Writer) *zap.Logger {
	stdout := zapcore.AddSync(w)

	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "time"
	productionCfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, stdout, level),
	)

	return zap.New(core)
}

type zapBench struct {
	l *zap.Logger
}

func (b *zapBench) new(w io.Writer) logBenchmark {
	return &zapBench{
		l: newZap(w),
	}
}

func (b *zapBench) newWithCtx(w io.Writer) logBenchmark {
	return &zapBench{
		l: newZap(w).With(zapFields()...),
	}
}

func (b *zapBench) name() string {
	return "Zap"
}

func (b *zapBench) logEventOnly(msg string) {
	b.l.Info(msg)
}

func (b *zapBench) logWithCtx(msg string) {
	b.l.Info(msg, zapFields()...)
}

func (b *zapBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *zapBench) logDisabledWithCtx(msg string) {
	b.l.Debug(msg, zapFields()...)
}

type zapSugarBench struct {
	l *zap.SugaredLogger
}

func (b *zapSugarBench) new(w io.Writer) logBenchmark {
	return &zapSugarBench{
		l: newZap(w).Sugar(),
	}
}

func (b *zapSugarBench) newWithCtx(w io.Writer) logBenchmark {
	return &zapSugarBench{
		l: newZap(w).Sugar().With(alternatingKeyValuePairs()...),
	}
}

func (b *zapSugarBench) name() string {
	return "ZapSugar"
}

func (b *zapSugarBench) logEventOnly(msg string) {
	b.l.Info(msg)
}

func (b *zapSugarBench) logWithCtx(msg string) {
	b.l.Infow(msg, alternatingKeyValuePairs()...)
}

func (b *zapSugarBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *zapSugarBench) logDisabledWithCtx(msg string) {
	b.l.Debugw(msg, alternatingKeyValuePairs()...)
}
