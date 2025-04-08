package config

type Config struct{
	AppHostPort  string    `env:"HOSTPORT,required"`
}

type PSQLConfig struct{
	Username string `env:"USER,required"`
	Password string	`env:"PASS,required"`
	Database string `env:"DB,required"`
}