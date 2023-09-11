package bench

import (
	"io"
	"testing"
)

// Log an event with no contextual fields.
func BenchmarkEvent(b *testing.B) {
	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.new(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logEventOnly(logMsg)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf("Log write count")
			}
		})
	}
}

// Log an event at a disabled level with no contextual fields,.
func BenchmarkEventDisabled(b *testing.B) {
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

// Log an event with several contextual fields.
func BenchmarkEventWithCtx(b *testing.B) {
	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.new(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logWithCtx(logMsg)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf("Log write count")
			}
		})
	}
}

// Log an event at a disabled level with several contextual fields,.
func BenchmarkEventDisabledWithCtx(b *testing.B) {
	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			l := v.new(io.Discard)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logDisabledWithCtx(logMsg)
				}
			})
		})
	}
}

// Create a logger with accumulated context.
func BenchmarkAccumulatedCtx(b *testing.B) {
	for _, v := range loggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.newWithCtx(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logEventOnly(logMsg)
				}
			})

			if out.WriteCount() != uint64(b.N) {
				b.Fatalf("Log write count")
			}
		})
	}
}

// Create a logger with accumulated context, but log at a disabled level.
func BenchmarkAccumulatedCtxDisabled(b *testing.B) {
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
