package config

type Config struct {
	DB       *DBConfig
	ReadOnly bool
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "dbuser",
			Password: "Password1234!",
			Name:     "eventus_api_database",
			Charset:  "utf8",
		},
		ReadOnly: false,
	}
}
