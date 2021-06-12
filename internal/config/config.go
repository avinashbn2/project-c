package config

type Config struct {
	DSN                 string `yaml:DSN`
	GAUTH_CLIENT_ID     string `yaml:"GAUTH_CLIENT_ID"`
	GAUTH_CLIENT_SECRET string `yaml:"GAUTH_CLIENT_SECRET"`
	JWTKEY              string `yaml:"JWTKEY"`
	CLIENT_URL          string `yaml:"CLIENT_URL"`
}
