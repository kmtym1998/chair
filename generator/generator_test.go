package generator_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/kmtym1998/chair/generator"
	"github.com/kmtym1998/chair/generator/config"
	"github.com/kmtym1998/chair/postgres"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRun_PostgreSQL(t *testing.T) {
	mockLdr := generator.NewSchemaLoaderMock()
	cfg := config.ConfigMock()

	cfg.Output = "./golden_testing/got/01_output.go"

	gen := generator.New(&cfg, postgres.DefaultMappers(), mockLdr)
	if err := gen.Run(context.Background()); err != nil {
		t.Fatalf("failed to generate go file: %v", err)
	}

	t.Run("assert generated code is correct", func(t *testing.T) {
		assertGoldenFile := func(t *testing.T, gotFileName, wantFileName string) bool {
			t.Helper()
			got, err := os.ReadFile("./golden_testing/got/" + gotFileName)
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}

			want, err := os.ReadFile("./golden_testing/want/" + wantFileName)
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}

			return assert.Equal(t, string(want), string(got))
		}

		assertGoldenFile(t, filepath.Base(cfg.Output), "01_output.go")
	})
}
