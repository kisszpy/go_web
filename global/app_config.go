package global

type AppConfig struct {
	Mysql *MySqlConfig `yaml:"mysql"`
	Jwt   *JwtConfig   `yaml:"jwt"`
}

type MySqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
type JwtConfig struct {
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
}
