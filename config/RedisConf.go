package config

type RedisConfig struct {
	Password            string `yaml:"password"`
	Host                string `yaml:"host"`
	Port                int    `yaml:"port"`
	Db                  int    `yaml:"db"`
	ReadTimeout         string `yaml:"readTimeout"`
	WriteTimeout        string `yaml:"writeTimeout"`
	TlsHandshakeTimeout string `yaml:"tlsHandshakeTimeout"`
}
