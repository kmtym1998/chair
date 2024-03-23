package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func ConfigMock() Config {
	yamlFileContent := `pkgName: 'pkgname'
output: 'pkgname/output.go'
postgres:
  schema: 'public'
mappings:
  - dbType: 'timestamp without time zone'
    goType: 'Time'
    goPkg: 'github.com/guregu/null'
    isNullable: true
`

	var cfg Config
	if err := yaml.Unmarshal([]byte(yamlFileContent), &cfg); err != nil {
		panic(fmt.Sprintf("failed to unmarshal config file: %v", err))
	}

	return cfg
}
