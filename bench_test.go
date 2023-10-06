package bench

import (
	"io"
	"testing"
)

// BenchmarkEvent tests the performance of logging a simple message with no
// contextual fields.
func BenchmarkEvent(b *testing.B) {
	b.Logf("Log a simple message without any contexual fields")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.new(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logEvent(logMsg)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in log write count. Expected: %d, Actual: %d",
					b.N,
					out.WriteCount(),
				)
			}
		})
	}
}

// BenchmarkDisabled tests the impact of logging at a disabled level for
// each library to determine how much overhead is incurred.
func BenchmarkDisabled(b *testing.B) {
	b.Logf("Log an event without any contexual fields")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			l := v.new(io.Discard)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logDisabled(logMsg)
				}
			})
		})
	}
}

// BenchmarkEventFmt tests the performance of logging a simple message with
// string formatting verbs.
func BenchmarkEventFmt(b *testing.B) {
	b.Logf("Log a simple message using string formatting verbs")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.new(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logEventFmt(logMsgFmt, logMsgArgs...)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in log write count. Expected: %d, Actual: %d",
					b.N,
					out.WriteCount(),
				)
			}
		})
	}
}

// BenchmarkDisabledFmt tests the performance of logging at a disabled level with
// string formatting verbs.
func BenchmarkDisabledFmt(b *testing.B) {
	b.Logf("Log at a disabled level with string formatting verbs")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			l := v.new(io.Discard)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logDisabledFmt(logMsgFmt, logMsgArgs...)
				}
			})
		})
	}
}

// BenchmarkEventCtx test the performance impact of each library when
// logging an event with several contextual fields.
func BenchmarkEventCtx(b *testing.B) {
	b.Logf("Log an event with several contextual fields")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.new(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logEventCtx(logMsg)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in log write count. Expected: %d, Actual: %d",
					b.N,
					out.WriteCount(),
				)
			}
		})
	}
}

// BenchmarkDisabledCtx tests the performance impact of logging an event
// at a disabled level with several contextual fields.
func BenchmarkDisabledCtx(b *testing.B) {
	b.Logf("Log a disabled event with several contextual fields")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			l := v.new(io.Discard)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logDisabledCtx(logMsg)
				}
			})
		})
	}
}

// BenchmarkEventCtxWeak tests the impact of logging an event with weakly typed
// contextual fields.
func BenchmarkEventCtxWeak(b *testing.B) {
	b.Logf("Log an event with weakly typed contextual fields")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.newWithCtx(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logEventCtxWeak(logMsg)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in log write count. Expected: %d, Actual: %d",
					b.N,
					out.WriteCount(),
				)
			}
		})
	}
}

// BenchmarkDisabledCtxWeak tests the impact of logging at a disabled level
// with weakly typed contextual fields.
func BenchmarkDisabledCtxWeak(b *testing.B) {
	b.Logf("Log at a disabled level with weakly typed contextual fields")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			l := v.newWithCtx(io.Discard)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logDisabledCtxWeak(logMsg)
				}
			})
		})
	}
}

// BenchmarkEventAccumulatedCtx tests the impact of creating a logger with
// accumulated context and using it to log events.
func BenchmarkEventAccumulatedCtx(b *testing.B) {
	b.Logf("Log an event with some accumulated context")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.newWithCtx(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logEvent(logMsg)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in log write count. Expected: %d, Actual: %d",
					b.N,
					out.WriteCount(),
				)
			}
		})
	}
}

// BenchmarkDisabledAccumulatedCtx creates a logger with accumulated context,
// but logs at a disabled level.
func BenchmarkDisabledAccumulatedCtx(b *testing.B) {
	b.Logf("Log a disabled event with some accumulated context")

	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			l := v.newWithCtx(io.Discard)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logDisabled(logMsg)
				}
			})
		})
	}
}
