package config

import (
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

type Config struct {
	DB          *DBConfig
	ReadOnly    bool
	ServicePort int
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

func DefaultConfig() *Config {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("server initialize: could not read default config file")
	}
	path = path + "/config_template.toml"
	return ParseConfig(path)
}

func ParseConfig(cfgFile string) *Config {

	content, err := toml.LoadFile(cfgFile)
	if err != nil {
		log.Fatal("server initialize: could not read config file. Error: " + err.Error())
	}

	readonly := content.Get("service.read-only").(bool)
	serverPort := content.Get("service.port").(int64)

	dbHost := content.Get("database.host").(string)
	dbPort := content.Get("database.port").(int64)
	dbUsername := content.Get("database.username").(string)
	dbPassword := content.Get("database.password").(string)
	databaseName := content.Get("database.database-name").(string)

	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     dbHost,
			Port:     int(dbPort),
			Username: dbUsername,
			Password: dbPassword,
			Name:     databaseName,
			Charset:  "utf8",
		},
		ReadOnly:    readonly,
		ServicePort: int(serverPort),
	}
}
