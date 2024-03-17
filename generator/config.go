package generator

type Config struct {
	Mappers []TypeMapping `yaml:"mappers"`
}

type TypeMapping struct {
	DBType string            `yaml:"dbType"`
	GoType string            `yaml:"goType"`
	GoPkg  string            `yaml:"goPkg"`
	Tag    map[string]string `yaml:"tag"`
}
