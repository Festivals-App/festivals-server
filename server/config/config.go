package config

import (
	"os"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/pelletier/go-toml"

	"github.com/rs/zerolog/log"
)

type Config struct {
	ServiceBindAddress string
	ServiceBindHost    string
	ServicePort        int
	ServiceKey         string
	TLSRootCert        string
	TLSCert            string
	TLSKey             string
	LoversEar          string
	Interval           int
	APIKeys            []string
	AdminKeys          []string
	DB                 *DBConfig
	ReadOnly           bool
}

type DBConfig struct {
	Dialect    string
	Host       string
	Port       int
	Username   string
	Password   string
	ClientCA   string
	ClientCert string
	ClientKey  string
	Name       string
	Charset    string
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
	if servertools.FileExists("/etc/festivals-server.conf") {
		config := ParseConfig("/etc/festivals-server.conf")
		if config != nil {
			return config
		}
	}

	// if there is no global configuration check the current folder for the template config file
	// this is mostly so the application will run in development environment
	path, err := os.Getwd()
	if err != nil {
		log.Fatal().Msg("server initialize: could not read default config file with error:" + err.Error())
	}
	path = path + "/config_template.toml"
	return ParseConfig(path)
}

func ParseConfig(cfgFile string) *Config {

	content, err := toml.LoadFile(cfgFile)
	if err != nil {
		log.Fatal().Err(err).Msg("server initialize: could not read config file at '" + cfgFile + "'")
	}

	serviceBindAdress := content.Get("service.bind-address").(string)
	serviceBindHost := content.Get("service.bind-host").(string)
	servicePort := content.Get("service.port").(int64)
	serviceKey := content.Get("service.key").(string)

	tlsrootcert := content.Get("tls.festivalsapp-root-ca").(string)
	tlscert := content.Get("tls.cert").(string)
	tlskey := content.Get("tls.key").(string)

	loversear := content.Get("heartbeat.endpoint").(string)
	interval := content.Get("heartbeat.interval").(int64)

	keyValues := content.Get("authentication.api-keys").([]interface{})
	keys := make([]string, len(keyValues))
	for i, v := range keyValues {
		keys[i] = v.(string)
	}
	adminKeyValues := content.Get("authentication.admin-keys").([]interface{})
	adminKeys := make([]string, len(adminKeyValues))
	for i, v := range adminKeyValues {
		adminKeys[i] = v.(string)
	}

	dbHost := content.Get("database.host").(string)
	dbPort := content.Get("database.port").(int64)
	dbUsername := content.Get("database.username").(string)
	dbPassword := content.Get("database.password").(string)
	dbClientCA := content.Get("database.festivalsapp-root-ca").(string)
	dbClientCert := content.Get("database.cert").(string)
	dbClientKey := content.Get("database.key").(string)
	readonly := content.Get("service.read-only").(bool)

	return &Config{
		ServiceBindAddress: serviceBindAdress,
		ServiceBindHost:    serviceBindHost,
		ServicePort:        int(servicePort),
		ServiceKey:         serviceKey,
		TLSRootCert:        tlsrootcert,
		TLSCert:            tlscert,
		TLSKey:             tlskey,
		LoversEar:          loversear,
		Interval:           int(interval),
		APIKeys:            keys,
		AdminKeys:          adminKeys,
		DB: &DBConfig{
			Dialect:    "mysql",
			Host:       dbHost,
			Port:       int(dbPort),
			Username:   dbUsername,
			Password:   dbPassword,
			ClientCA:   dbClientCA,
			ClientCert: dbClientCert,
			ClientKey:  dbClientKey,
			Name:       "festivals_api_database",
			Charset:    "utf8",
		},
		ReadOnly: readonly,
	}
}
