package global

type AppConfig struct {
	Mysql *MySqlConfig `yaml:"mysql"`
	Jwt   *JwtConfig   `yaml:"jwt"`
}

// MySqlConfig 配置清单/**
type MySqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// JwtConfig 配置清单/**
type JwtConfig struct {
	Header  string `yaml:"header"`
	Issuer  string `yaml:"issuer"`
	Key     string `yaml:"key"`
	Context string `yaml:"context"`
}
