package config

type Config struct{
	AppHostPort  string    `env:"HOSTPORT,required"`
}

type PSQLConfig struct{
	Port string `env:"PORT,required"`
	Username string `env:"USERDB,required"`
	Password string	`env:"PASS,required"`
	Database string `env:"DB,required"`
}