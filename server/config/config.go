package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	DB                 *DBConfig
	ReadOnly           bool
	ServiceBindAddress string
	ServicePort        int
	APIKeys            []string
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

	/// TODO Add support for config from environment variable
	/*
		httpPort := os.Getenv("HTTP_PORT")
		if httpPort == "" {
			httpPort = "8080"
		}
	*/

	// first we try to parse the config at the global configuration path
	if fileExists("/etc/festivals-server.conf") {
		config := ParseConfig("/etc/festivals-server.conf")
		if config != nil {
			return config
		}
	}

	// if there is no global configuration check the current folder for the template config file
	// this is mostly so the application will run in development environment
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
		log.Fatal("server initialize: could not read config file at '" + cfgFile + "'. Error: " + err.Error())
	}

	readonly := content.Get("service.read-only").(bool)
	serverBindAdress := content.Get("service.bind-address").(string)
	serverPort := content.Get("service.port").(int64)

	keyValues := content.Get("authentication.api-keys").([]interface{})
	keys := make([]string, len(keyValues))
	for i, v := range keyValues {
		keys[i] = v.(string)
	}

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
		ReadOnly:           readonly,
		ServiceBindAddress: serverBindAdress,
		ServicePort:        int(serverPort),
		APIKeys:            keys,
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
// see: https://golangcode.com/check-if-a-file-exists/
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
