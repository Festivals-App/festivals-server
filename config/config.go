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
			Host:     "192.168.8.206",
			Port:     3306,
			Username: "festivals.api.writer",
			Password: "we4711",
			Name:     "festivals_api_database",
			Charset:  "utf8",
		},
		ReadOnly: false,
	}
}
