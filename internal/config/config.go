package config

type Config struct{
	AppPort  int    `env:"PORT,required"`
}

type PSQLConfig struct{
	Username string `env:"USER,required"`
	Password string	`env:"PASS,required"`
	Database string `env:"DB,required"`
}