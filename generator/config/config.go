package config

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mappings []TypeMapping `yaml:"mappings"`
}

type TypeMapping struct {
	DBType string `yaml:"dbType"`
	GoType string `yaml:"goType"`
	GoPkg  string `yaml:"goPkg"`
}

func Parse(cfgFileName string) (*Config, error) {
	cfgFile, err := os.ReadFile(cfgFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(cfgFile, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &cfg, nil
}

type contextKey struct{}

func With(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, contextKey{}, cfg)
}

func From(ctx context.Context) (*Config, bool) {
	cfg, ok := ctx.Value(contextKey{}).(*Config)
	return cfg, ok
}
