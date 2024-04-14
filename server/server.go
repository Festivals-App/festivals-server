package server

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"

	token "github.com/Festivals-App/festivals-identity-server/jwt"
	festivalspki "github.com/Festivals-App/festivals-pki"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/Festivals-App/festivals-server/server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

// Server has router and db instances
type Server struct {
	Router    *chi.Mux
	DB        *sql.DB
	Config    *config.Config
	TLSConfig *tls.Config
	Validator *token.ValidationService
}

func NewServer(config *config.Config) *Server {
	server := &Server{}
	server.initialize(config)
	return server
}

// Initialize the server with predefined configuration
func (s *Server) initialize(config *config.Config) {

	s.Config = config
	s.Router = chi.NewRouter()

	s.setIdentityService()
	s.setDatabase()
	s.setTLSHandling()
	s.setMiddleware()
	s.setRoutes(config)
}

func (s *Server) setIdentityService() {

	config := s.Config

	val := token.NewValidationService(config.IdentityEndpoint, config.TLSCert, config.TLSKey, config.ServiceKey, false)
	if val == nil {
		log.Fatal().Msg("Failed to create validator.")
	}
	s.Validator = val
}

var mysqlTLSConfigKey string = "org.festivalsapp.mysql.tls"

func (s *Server) setDatabase() {

	config := s.Config

	rootCertPool, err := festivalspki.LoadCertificatePool(config.DB.ClientCA)
	if err != nil {
		log.Fatal().Err(err).Msg("Faile to create pool with root CA file.")
	}

	certs, err := tls.LoadX509KeyPair(config.DB.ClientCert, config.DB.ClientKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Faile to load database client certificate.")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		RootCAs:      rootCertPool,
		Certificates: []tls.Certificate{certs},
	}
	mysql.RegisterTLSConfig(mysqlTLSConfigKey, tlsConfig)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&tls=%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset,
		mysqlTLSConfigKey,
	)
	db, err := sql.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open database handle.")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database.")
	}

	db.SetConnMaxIdleTime(time.Minute * 1)
	db.SetConnMaxLifetime(time.Minute * 5)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)

	s.DB = db
}

func (s *Server) setTLSHandling() {
	tlsConfig := &tls.Config{
		ClientAuth:     tls.RequireAndVerifyClientCert,
		GetCertificate: festivalspki.LoadServerCertificateHandler(s.Config.TLSCert, s.Config.TLSKey, s.Config.TLSRootCert),
	}
	s.TLSConfig = tlsConfig
}

func (s *Server) setMiddleware() {

	// tell the router which middleware to use
	s.Router.Use(
		// used to log the request to the console
		servertools.Middleware(servertools.TraceLogger("/var/log/festivals-server/trace.log")),
		// tries to recover after panics (?)
		middleware.Recoverer,
	)
}

// setRouters sets the all required routers
func (s *Server) setRoutes(config *config.Config) {

	s.Router.Get("/version", s.handleRequest(handler.GetVersion))
	s.Router.Get("/info", s.handleRequest(handler.GetInfo))
	s.Router.Get("/health", s.handleRequest(handler.GetHealth))

	s.Router.Post("/update", s.handleRequest(handler.MakeUpdate))
	s.Router.Get("/log", s.handleRequest(handler.GetLog))
	s.Router.Get("/log/trace", s.handleRequest(handler.GetTraceLog))

	s.Router.Get("/festivals", s.handleAPIRequest(handler.GetFestivals))
	s.Router.Get("/festivals/{objectID}", s.handleAPIRequest(handler.GetFestival))
	s.Router.Get("/festivals/{objectID}/events", s.handleAPIRequest(handler.GetFestivalEvents))
	s.Router.Get("/festivals/{objectID}/image", s.handleAPIRequest(handler.GetFestivalImage))
	s.Router.Get("/festivals/{objectID}/links", s.handleAPIRequest(handler.GetFestivalLinks))
	s.Router.Get("/festivals/{objectID}/place", s.handleAPIRequest(handler.GetFestivalPlace))
	s.Router.Get("/festivals/{objectID}/tags", s.handleAPIRequest(handler.GetFestivalTags))

	s.Router.Get("/artists", s.handleAPIRequest(handler.GetArtists))
	s.Router.Get("/artists/{objectID}", s.handleAPIRequest(handler.GetArtist))
	s.Router.Get("/artists/{objectID}/image", s.handleAPIRequest(handler.GetArtistImage))
	s.Router.Get("/artists/{objectID}/links", s.handleAPIRequest(handler.GetArtistLinks))
	s.Router.Get("/artists/{objectID}/tags", s.handleAPIRequest(handler.GetArtistTags))

	s.Router.Get("/locations", s.handleAPIRequest(handler.GetLocations))
	s.Router.Get("/locations/{objectID}", s.handleAPIRequest(handler.GetLocation))
	s.Router.Get("/locations/{objectID}/image", s.handleAPIRequest(handler.GetLocationImage))
	s.Router.Get("/locations/{objectID}/links", s.handleAPIRequest(handler.GetLocationLinks))
	s.Router.Get("/locations/{objectID}/place", s.handleAPIRequest(handler.GetLocationPlace))

	s.Router.Get("/events", s.handleAPIRequest(handler.GetEvents))
	s.Router.Get("/events/{objectID}", s.handleAPIRequest(handler.GetEvent))
	s.Router.Get("/events/{objectID}/festival", s.handleAPIRequest(handler.GetEventFestival))
	s.Router.Get("/events/{objectID}/image", s.handleAPIRequest(handler.GetEventImage))
	s.Router.Get("/events/{objectID}/artist", s.handleAPIRequest(handler.GetEventArtist))
	s.Router.Get("/events/{objectID}/location", s.handleAPIRequest(handler.GetEventLocation))

	s.Router.Get("/images", s.handleAPIRequest(handler.GetImages))
	s.Router.Get("/images/{objectID}", s.handleAPIRequest(handler.GetImage))

	s.Router.Get("/links", s.handleAPIRequest(handler.GetLinks))
	s.Router.Get("/links/{objectID}", s.handleAPIRequest(handler.GetLink))

	s.Router.Get("/places", s.handleAPIRequest(handler.GetPlaces))
	s.Router.Get("/places/{objectID}", s.handleAPIRequest(handler.GetPlace))

	s.Router.Get("/tags", s.handleAPIRequest(handler.GetTags))
	s.Router.Get("/tags/{objectID}", s.handleAPIRequest(handler.GetTag))
	s.Router.Get("/tags/{objectID}/festivals", s.handleAPIRequest(handler.GetTagFestivals))

	if !config.ReadOnly {

		s.Router.Post("/festivals", s.handleRequest(handler.CreateFestival))
		s.Router.Patch("/festivals/{objectID}", s.handleRequest(handler.UpdateFestival))
		s.Router.Delete("/festivals/{objectID}", s.handleRequest(handler.DeleteFestival))
		s.Router.Post("/festivals/{objectID}/events/{resourceID}", s.handleRequest(handler.SetEventForFestival))
		s.Router.Post("/festivals/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForFestival))
		s.Router.Post("/festivals/{objectID}/links/{resourceID}", s.handleRequest(handler.SetLinkForFestival))
		s.Router.Post("/festivals/{objectID}/place/{resourceID}", s.handleRequest(handler.SetPlaceForFestival))
		s.Router.Post("/festivals/{objectID}/tags/{resourceID}", s.handleRequest(handler.SetTagForFestival))
		s.Router.Delete("/festivals/{objectID}/image/{resourceID}", s.handleRequest(handler.RemoveImageForFestival))
		s.Router.Delete("/festivals/{objectID}/links/{resourceID}", s.handleRequest(handler.RemoveLinkForFestival))
		s.Router.Delete("/festivals/{objectID}/place/{resourceID}", s.handleRequest(handler.RemovePlaceForFestival))
		s.Router.Delete("/festivals/{objectID}/tags/{resourceID}", s.handleRequest(handler.RemoveTagForFestival))

		s.Router.Post("/artists", s.handleRequest(handler.CreateArtist))
		s.Router.Patch("/artists/{objectID}", s.handleRequest(handler.UpdateArtist))
		s.Router.Delete("/artists/{objectID}", s.handleRequest(handler.DeleteArtist))
		s.Router.Post("/artists/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForArtist))
		s.Router.Post("/artists/{objectID}/links/{resourceID}", s.handleRequest(handler.SetLinkForArtist))
		s.Router.Post("/artists/{objectID}/tags/{resourceID}", s.handleRequest(handler.SetTagForArtist))
		s.Router.Delete("/artists/{objectID}/image/{resourceID}", s.handleRequest(handler.RemoveImageForArtist))
		s.Router.Delete("/artists/{objectID}/links/{resourceID}", s.handleRequest(handler.RemoveLinkForArtist))
		s.Router.Delete("/artists/{objectID}/tags/{resourceID}", s.handleRequest(handler.RemoveTagForArtist))

		s.Router.Post("/locations", s.handleRequest(handler.CreateLocation))
		s.Router.Patch("/locations/{objectID}", s.handleRequest(handler.UpdateLocation))
		s.Router.Delete("/locations/{objectID}", s.handleRequest(handler.DeleteLocation))
		s.Router.Post("/locations/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForLocation))
		s.Router.Post("/locations/{objectID}/links/{resourceID}", s.handleRequest(handler.SetLinkForLocation))
		s.Router.Post("/locations/{objectID}/place/{resourceID}", s.handleRequest(handler.SetPlaceForLocation))
		s.Router.Delete("/locations/{objectID}/image/{resourceID}", s.handleRequest(handler.RemoveImageForLocation))
		s.Router.Delete("/locations/{objectID}/links/{resourceID}", s.handleRequest(handler.RemoveLinkForLocation))
		s.Router.Delete("/locations/{objectID}/place/{resourceID}", s.handleRequest(handler.RemovePlaceForLocation))

		s.Router.Post("/events", s.handleRequest(handler.CreateEvent))
		s.Router.Patch("/events/{objectID}", s.handleRequest(handler.UpdateEvent))
		s.Router.Delete("/events/{objectID}", s.handleRequest(handler.DeleteEvent))
		s.Router.Post("/events/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForEvent))
		s.Router.Post("/events/{objectID}/artist/{resourceID}", s.handleRequest(handler.SetArtistForEvent))
		s.Router.Post("/events/{objectID}/location/{resourceID}", s.handleRequest(handler.SetLocationForEvent))
		s.Router.Delete("/events/{objectID}/image/{resourceID}", s.handleRequest(handler.RemoveImageForEvent))
		s.Router.Delete("/events/{objectID}/artist/{resourceID}", s.handleRequest(handler.RemoveArtistForEvent))
		s.Router.Delete("/events/{objectID}/location/{resourceID}", s.handleRequest(handler.RemoveLocationForEvent))

		s.Router.Post("/images", s.handleRequest(handler.CreateImage))
		s.Router.Patch("/images/{objectID}", s.handleRequest(handler.UpdateImage))
		s.Router.Delete("/images/{objectID}", s.handleRequest(handler.DeleteImage))

		s.Router.Post("/links", s.handleRequest(handler.CreateLink))
		s.Router.Patch("/links/{objectID}", s.handleRequest(handler.UpdateLink))
		s.Router.Delete("/links/{objectID}", s.handleRequest(handler.DeleteLink))

		s.Router.Post("/places", s.handleRequest(handler.CreatePlace))
		s.Router.Patch("/places/{objectID}", s.handleRequest(handler.UpdatePlace))
		s.Router.Delete("/places/{objectID}", s.handleRequest(handler.DeletePlace))

		s.Router.Post("/tags", s.handleRequest(handler.CreateTag))
		s.Router.Patch("/tags/{objectID}", s.handleRequest(handler.UpdateTag))
		s.Router.Delete("/tags/{objectID}", s.handleRequest(handler.DeleteTag))
	}
}

func (s *Server) Run(conf *config.Config) {

	server := http.Server{
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,

		Addr:      conf.ServiceBindHost + ":" + strconv.Itoa(conf.ServicePort),
		Handler:   s.Router,
		TLSConfig: s.TLSConfig,
	}

	//server.SetKeepAlivesEnabled(false)

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal().Err(err).Str("type", "server").Msg("Failed to run server")
	}
}

type APIKeyAuthenticatedHandlerFunction func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func (s *Server) handleAPIRequest(requestHandler APIKeyAuthenticatedHandlerFunction) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apikey := token.GetAPIToken(r)
		if !slices.Contains((*s.Validator.APIKeys), apikey) {
			claims := token.GetValidClaims(r, s.Validator)
			if claims == nil {
				servertools.UnauthorizedResponse(w)
				return
			}
		}
		requestHandler(s.DB, w, r)
	})
}

type JWTAuthenticatedHandlerFunction func(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request)

func (s *Server) handleRequest(requestHandler JWTAuthenticatedHandlerFunction) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims := token.GetValidClaims(r, s.Validator)
		if claims == nil {
			servertools.UnauthorizedResponse(w)
			return
		}
		requestHandler(s.Validator, claims, s.Config, s.DB, w, r)
	})
}
