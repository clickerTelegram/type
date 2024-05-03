package db

type ConfigEnv struct {
	PostgresUrl        string `env:"POSTGRES_URL"`
	PostgresPort       int64  `env:"POSTGRES_PORT"`
	PostgresDb         string `env:"POSTGRES_DB"`
	PostgresUsername   string `env:"POSTGRES_USERNAME"`
	PostgresPassword   string `env:"POSTGRES_PASSWORD"`
	SslMode            string `env:"SSL_MODE"`
	TimeZone           string `env:"TIME_ZONE"`
	SetMaxIdleConns    int    `env:"SET_MAX_IDLE_CONNS"`
	SetMaxOpenConns    int    `env:"SET_MAX_OPEN_CONNS"`
	SetConnMaxLifetime int    `env:"SET_CONN_MAX_LIFETIME"`
}
