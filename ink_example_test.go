package ink_test

import (
	"context"
	"os"
	"time"

	"github.com/hemantjadon/ink/inkio"

	"github.com/hemantjadon/ink"
)

func ExampleLogger_message() {
	sink := inkio.NewSink(os.Stdout)
	logger := ink.NewLogger(ink.WithSink(sink))

	logger.Info("detected system health")

	// Output:
	// INFO detected system health
}

func ExampleLogger_message_with_fields() {
	sink := inkio.NewSink(os.Stdout)
	logger := ink.NewLogger(ink.WithSink(sink))

	logger.Info(
		"detected system health",
		ink.String("name", "ink"),
		ink.Uint64("uptime_days", 9),
		ink.Int64("temperature_celsius", -25),
		ink.Float64("load_factor", 0.275),
		ink.Bool("is_active", true),
		ink.Duration("timeout", 5*time.Second),
		ink.Strings("traits", []string{"primary", "master"}),
	)

	// Output:
	// INFO detected system health name=ink uptime_days=9 temperature_celsius=-25 load_factor=0.275 is_active=true timeout=5s traits=[primary master]
}

func ExampleLogger_levels() {
	sink := inkio.NewSink(os.Stdout)
	logger := ink.NewLogger(ink.WithSink(sink))

	logger.Info("detected system health")
	logger.Debug("detected system health")
	logger.Error("detected system health")

	// Output:
	// INFO detected system health
	// DEBUG detected system health
	// ERROR detected system health
}

func ExampleLogger_logger_level_fields() {
	sink := inkio.NewSink(os.Stdout)
	logger := ink.NewLogger(ink.WithSink(sink), ink.WithFields(ink.String("package", "ink"), ink.String("source", "server")))

	logger.Info("detected system health",
		ink.String("name", "ink"),
		ink.Uint64("uptime_days", 9),
		ink.Int64("temperature_celsius", -25),
		ink.Float64("load_factor", 0.275),
		ink.Bool("is_active", true),
		ink.Duration("timeout", 5*time.Second),
		ink.Strings("traits", []string{"primary", "master"}),
	)

	// Output:
	// INFO detected system health package=ink source=server name=ink uptime_days=9 temperature_celsius=-25 load_factor=0.275 is_active=true timeout=5s traits=[primary master]
}

func ExampleLogger_logger_name() {
	sink := inkio.NewSink(os.Stdout)
	logger := ink.NewLogger(ink.WithSink(sink), ink.WithName("httpserver"))

	logger.Info("detected system health",
		ink.String("name", "ink"),
		ink.Uint64("uptime_days", 9),
		ink.Int64("temperature_celsius", -25),
		ink.Float64("load_factor", 0.275),
		ink.Bool("is_active", true),
		ink.Duration("timeout", 5*time.Second),
		ink.Strings("traits", []string{"primary", "master"}),
	)

	// Output:
	// INFO detected system health logger=httpserver name=ink uptime_days=9 temperature_celsius=-25 load_factor=0.275 is_active=true timeout=5s traits=[primary master]
}

func ExampleLogger_context_fields() {
	sink := inkio.NewSink(os.Stdout)
	logger := ink.NewLogger(ink.WithSink(sink))

	ctx := context.Background()
	ctx = ink.ContextWithFields(ctx, ink.String("source", "server"), ink.Int("id", 1))

	logger.Info("detected system health", append(
		ink.ContextFields(ctx),
		ink.String("name", "ink"),
		ink.Uint64("uptime_days", 9),
		ink.Int64("temperature_celsius", -25),
		ink.Float64("load_factor", 0.275),
		ink.Bool("is_active", true),
		ink.Duration("timeout", 5*time.Second),
		ink.Strings("traits", []string{"primary", "master"}),
	)...)

	// Output:
	// INFO detected system health source=server id=1 name=ink uptime_days=9 temperature_celsius=-25 load_factor=0.275 is_active=true timeout=5s traits=[primary master]
}
