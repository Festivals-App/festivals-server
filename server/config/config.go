package config

import (
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/pelletier/go-toml"

	"github.com/rs/zerolog/log"
)

type Config struct {
	ServiceBindHost  string
	ServicePort      int
	ServiceKey       string
	TLSRootCert      string
	TLSCert          string
	TLSKey           string
	LoversEar        string
	Interval         int
	IdentityEndpoint string
	DB               *DBConfig
	ReadOnly         bool
	InfoLog          string
	TraceLog         string
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

func ParseConfig(cfgFile string) *Config {

	content, err := toml.LoadFile(cfgFile)
	if err != nil {
		log.Fatal().Err(err).Msg("server initialize: could not read config file at '" + cfgFile + "'. Error: " + err.Error())
	}

	serviceBindHost := content.Get("service.bind-host").(string)
	servicePort := content.Get("service.port").(int64)
	serviceKey := content.Get("service.key").(string)

	tlsrootcert := content.Get("tls.festivalsapp-root-ca").(string)
	tlscert := content.Get("tls.cert").(string)
	tlskey := content.Get("tls.key").(string)

	loversear := content.Get("heartbeat.endpoint").(string)
	interval := content.Get("heartbeat.interval").(int64)

	identity := content.Get("authentication.endpoint").(string)

	dbHost := content.Get("database.host").(string)
	dbPort := content.Get("database.port").(int64)
	dbUsername := content.Get("database.username").(string)
	dbPassword := content.Get("database.password").(string)
	dbClientCA := content.Get("database.festivalsapp-root-ca").(string)
	dbClientCert := content.Get("database.cert").(string)
	dbClientKey := content.Get("database.key").(string)
	readonly := content.Get("service.read-only").(bool)

	infoLogPath := content.Get("log.info").(string)
	traceLogPath := content.Get("log.trace").(string)

	tlsrootcert = servertools.ExpandTilde(tlsrootcert)
	tlscert = servertools.ExpandTilde(tlscert)
	tlskey = servertools.ExpandTilde(tlskey)
	dbClientCA = servertools.ExpandTilde(dbClientCA)
	dbClientCert = servertools.ExpandTilde(dbClientCert)
	dbClientKey = servertools.ExpandTilde(dbClientKey)
	infoLogPath = servertools.ExpandTilde(infoLogPath)
	traceLogPath = servertools.ExpandTilde(traceLogPath)

	return &Config{
		ServiceBindHost:  serviceBindHost,
		ServicePort:      int(servicePort),
		ServiceKey:       serviceKey,
		TLSRootCert:      tlsrootcert,
		TLSCert:          tlscert,
		TLSKey:           tlskey,
		LoversEar:        loversear,
		Interval:         int(interval),
		IdentityEndpoint: identity,
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
		InfoLog:  infoLogPath,
		TraceLog: traceLogPath,
	}
}
