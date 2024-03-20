package generator

import (
	"log/slog"
	"os"
	"testing"

	"github.com/lmittmann/tint"
)

func TestToUpperCamel(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.New(tint.NewHandler(
		os.Stdout,
		&tint.Options{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	))

	tests := []struct {
		field Field
		want  string
	}{
		{"hello", "Hello"},
		{"he_llo", "HeLlo"},
		{"hello_world", "HelloWorld"},
		{"hello_world_foo", "HelloWorldFoo"},
		{"_PrismaTable", "PrismaTable"},
		{"id", "ID"},
		{"id_token", "IDToken"},
		{"uuid_is_not_url", "UUIDIsNotURL"},
		{"userId", "UserID"},
		{"UpperCamel", "UpperCamel"},
		{"lowerCamel", "LowerCamel"},
	}
	for _, tt := range tests {
		t.Run(string(tt.field), func(t *testing.T) {
			if got := tt.field.ToUpperCamel(); got != tt.want {
				t.Errorf("Field.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
