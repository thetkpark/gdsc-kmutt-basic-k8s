package config

type DB struct {
	Host     string `env:"DB_HOST,notEmpty"`
	Port     string `env:"DB_PORT,notEmpty"`
	User     string `env:"DB_USER,notEmpty"`
	Password string `env:"DB_PASSWORD,notEmpty"`
	Database string `env:"DB_DATABASE,notEmpty"`
}
