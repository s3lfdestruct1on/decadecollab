package config

type Config struct{
	AppPort  int    `env:"PORT,required"`
}

type PSQLConfig struct{
	Port string `env:"PORT,required"`
	Username string `env:"USER,required"`
	Password string	`env:"PASS,required"`
	Database string `env:"DB,required"`
}