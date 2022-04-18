package global

type AppConfig struct {
	Mysql   *MySqlConfig   `yaml:"mysql"`
	Jwt     *JwtConfig     `yaml:"jwt"`
	Redis   *RedisConfig   `yaml:"redis"`
	MongoDb *MongoDbConfig `yaml:"mongo"`
}

type MongoDbConfig struct {
	Uri string `yaml:"uri"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   int    `yaml:"db"`
	Pass string `yaml:"pass"`
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
