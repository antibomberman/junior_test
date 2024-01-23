package config

type Config struct {
}

func MustLoad() *Config {
	//path := "./config/local.yaml"
	var cfg Config
	//if err := cleanenv.ReadConfig(path, &cfg); err != nil {
	//	panic("failed to read config: " + err.Error())
	//}
	return &cfg
}
