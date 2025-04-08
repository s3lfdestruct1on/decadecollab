package config

type Config struct{
	AppPort  int    `env:"PORT,required,omitempty"`
}

type PSQLConfig struct{
	Username string `env:"USER,required,omitempty"`
	Password string	`env:"PASS,required,omitempty"`
	Database string `env:"DB,required,omitempty"`
}