package server

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Festivals-App/festivals-gateway/server/logger"
	"github.com/Festivals-App/festivals-identity-server/authentication"
	"github.com/Festivals-App/festivals-identity-server/festivalspki"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/Festivals-App/festivals-server/server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/acme/autocert"
)

// Server has router and db instances
type Server struct {
	Router      *chi.Mux
	DB          *sql.DB
	Config      *config.Config
	CertManager *autocert.Manager
	TLSConfig   *tls.Config
}

// Initialize the server with predefined configuration
func (s *Server) Initialize(config *config.Config) {

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)
	db, err := sql.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to database")
	}

	s.DB = db
	s.Router = chi.NewRouter()
	s.Config = config

	s.setTLSHandling()
	s.setMiddleware()
	s.setRoutes(config)
}

func (s *Server) setTLSHandling() {

	base := s.Config.ServiceBindAddress
	hosts := []string{base}

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(hosts...),
		Cache:      autocert.DirCache("/etc/letsencrypt/live/" + base),
	}

	tlsConfig := certManager.TLSConfig()
	tlsConfig.GetCertificate = festivalspki.LoadServerCertificates(s.Config.TLSCert, s.Config.TLSKey, s.Config.TLSRootCert, &certManager)
	s.CertManager = &certManager
	s.TLSConfig = tlsConfig
}

func (s *Server) setMiddleware() {

	// tell the router which middleware to use
	s.Router.Use(
		// used to log the request to the console
		logger.Middleware(logger.TraceLogger("/var/log/festivals-server/trace.log")),
		// tries to recover after panics (?)
		middleware.Recoverer,
	)
}

// setRouters sets the all required routers
func (s *Server) setRoutes(config *config.Config) {

	s.Router.Get("/version", s.handleRequestWithoutValidation(handler.GetVersion))
	s.Router.Get("/info", s.handleRequestWithoutValidation(handler.GetInfo))
	s.Router.Get("/health", s.handleRequestWithoutValidation(handler.GetHealth))

	s.Router.Post("/update", s.handleAdminRequest(handler.MakeUpdate))
	s.Router.Get("/log", s.handleAdminRequest(handler.GetLog))
	s.Router.Get("/log/trace", s.handleAdminRequest(handler.GetTraceLog))

	s.Router.Get("/festivals", s.handleRequest(handler.GetFestivals))
	s.Router.Get("/festivals/{objectID}", s.handleRequest(handler.GetFestival))
	s.Router.Get("/festivals/{objectID}/events", s.handleRequest(handler.GetFestivalEvents))
	s.Router.Get("/festivals/{objectID}/image", s.handleRequest(handler.GetFestivalImage))
	s.Router.Get("/festivals/{objectID}/links", s.handleRequest(handler.GetFestivalLinks))
	s.Router.Get("/festivals/{objectID}/place", s.handleRequest(handler.GetFestivalPlace))
	s.Router.Get("/festivals/{objectID}/tags", s.handleRequest(handler.GetFestivalTags))

	s.Router.Get("/artists", s.handleRequest(handler.GetArtists))
	s.Router.Get("/artists/{objectID}", s.handleRequest(handler.GetArtist))
	s.Router.Get("/artists/{objectID}/image", s.handleRequest(handler.GetArtistImage))
	s.Router.Get("/artists/{objectID}/links", s.handleRequest(handler.GetArtistLinks))
	s.Router.Get("/artists/{objectID}/tags", s.handleRequest(handler.GetArtistTags))

	s.Router.Get("/locations", s.handleRequest(handler.GetLocations))
	s.Router.Get("/locations/{objectID}", s.handleRequest(handler.GetLocation))
	s.Router.Get("/locations/{objectID}/image", s.handleRequest(handler.GetLocationImage))
	s.Router.Get("/locations/{objectID}/links", s.handleRequest(handler.GetLocationLinks))
	s.Router.Get("/locations/{objectID}/place", s.handleRequest(handler.GetLocationPlace))

	s.Router.Get("/events", s.handleRequest(handler.GetEvents))
	s.Router.Get("/events/{objectID}", s.handleRequest(handler.GetEvent))
	s.Router.Get("/events/{objectID}/festival", s.handleRequest(handler.GetEventFestival))
	s.Router.Get("/events/{objectID}/image", s.handleRequest(handler.GetEventImage))
	s.Router.Get("/events/{objectID}/artist", s.handleRequest(handler.GetEventArtist))
	s.Router.Get("/events/{objectID}/location", s.handleRequest(handler.GetEventLocation))

	s.Router.Get("/images", s.handleRequest(handler.GetImages))
	s.Router.Get("/images/{objectID}", s.handleRequest(handler.GetImage))

	s.Router.Get("/links", s.handleRequest(handler.GetLinks))
	s.Router.Get("/links/{objectID}", s.handleRequest(handler.GetLink))

	s.Router.Get("/places", s.handleRequest(handler.GetPlaces))
	s.Router.Get("/places/{objectID}", s.handleRequest(handler.GetPlace))

	s.Router.Get("/tags", s.handleRequest(handler.GetTags))
	s.Router.Get("/tags/{objectID}", s.handleRequest(handler.GetTag))
	s.Router.Get("/tags/{objectID}/festivals", s.handleRequest(handler.GetTagFestivals))

	if !config.ReadOnly {

		s.Router.Post("/festivals", s.handleAdminRequest(handler.CreateFestival))
		s.Router.Patch("/festivals/{objectID}", s.handleAdminRequest(handler.UpdateFestival))
		s.Router.Delete("/festivals/{objectID}", s.handleAdminRequest(handler.DeleteFestival))
		s.Router.Post("/festivals/{objectID}/events/{resourceID}", s.handleAdminRequest(handler.SetEventForFestival))
		s.Router.Post("/festivals/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.SetImageForFestival))
		s.Router.Post("/festivals/{objectID}/links/{resourceID}", s.handleAdminRequest(handler.SetLinkForFestival))
		s.Router.Post("/festivals/{objectID}/place/{resourceID}", s.handleAdminRequest(handler.SetPlaceForFestival))
		s.Router.Post("/festivals/{objectID}/tags/{resourceID}", s.handleAdminRequest(handler.SetTagForFestival))
		s.Router.Delete("/festivals/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.RemoveImageForFestival))
		s.Router.Delete("/festivals/{objectID}/links/{resourceID}", s.handleAdminRequest(handler.RemoveLinkForFestival))
		s.Router.Delete("/festivals/{objectID}/place/{resourceID}", s.handleAdminRequest(handler.RemovePlaceForFestival))
		s.Router.Delete("/festivals/{objectID}/tags/{resourceID}", s.handleAdminRequest(handler.RemoveTagForFestival))

		s.Router.Post("/artists", s.handleAdminRequest(handler.CreateArtist))
		s.Router.Patch("/artists/{objectID}", s.handleAdminRequest(handler.UpdateArtist))
		s.Router.Delete("/artists/{objectID}", s.handleAdminRequest(handler.DeleteArtist))
		s.Router.Post("/artists/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.SetImageForArtist))
		s.Router.Post("/artists/{objectID}/links/{resourceID}", s.handleAdminRequest(handler.SetLinkForArtist))
		s.Router.Post("/artists/{objectID}/tags/{resourceID}", s.handleAdminRequest(handler.SetTagForArtist))
		s.Router.Delete("/artists/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.RemoveImageForArtist))
		s.Router.Delete("/artists/{objectID}/links/{resourceID}", s.handleAdminRequest(handler.RemoveLinkForArtist))
		s.Router.Delete("/artists/{objectID}/tags/{resourceID}", s.handleAdminRequest(handler.RemoveTagForArtist))

		s.Router.Post("/locations", s.handleAdminRequest(handler.CreateLocation))
		s.Router.Patch("/locations/{objectID}", s.handleAdminRequest(handler.UpdateLocation))
		s.Router.Delete("/locations/{objectID}", s.handleAdminRequest(handler.DeleteLocation))
		s.Router.Post("/locations/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.SetImageForLocation))
		s.Router.Post("/locations/{objectID}/links/{resourceID}", s.handleAdminRequest(handler.SetLinkForLocation))
		s.Router.Post("/locations/{objectID}/place/{resourceID}", s.handleAdminRequest(handler.SetPlaceForLocation))
		s.Router.Delete("/locations/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.RemoveImageForLocation))
		s.Router.Delete("/locations/{objectID}/links/{resourceID}", s.handleAdminRequest(handler.RemoveLinkForLocation))
		s.Router.Delete("/locations/{objectID}/place/{resourceID}", s.handleAdminRequest(handler.RemovePlaceForLocation))

		s.Router.Post("/events", s.handleAdminRequest(handler.CreateEvent))
		s.Router.Patch("/events/{objectID}", s.handleAdminRequest(handler.UpdateEvent))
		s.Router.Delete("/events/{objectID}", s.handleAdminRequest(handler.DeleteEvent))
		s.Router.Post("/events/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.SetImageForEvent))
		s.Router.Post("/events/{objectID}/artist/{resourceID}", s.handleAdminRequest(handler.SetArtistForEvent))
		s.Router.Post("/events/{objectID}/location/{resourceID}", s.handleAdminRequest(handler.SetLocationForEvent))
		s.Router.Delete("/events/{objectID}/image/{resourceID}", s.handleAdminRequest(handler.RemoveImageForEvent))
		s.Router.Delete("/events/{objectID}/artist/{resourceID}", s.handleAdminRequest(handler.RemoveArtistForEvent))
		s.Router.Delete("/events/{objectID}/location/{resourceID}", s.handleAdminRequest(handler.RemoveLocationForEvent))

		s.Router.Post("/images", s.handleAdminRequest(handler.CreateImage))
		s.Router.Patch("/images/{objectID}", s.handleAdminRequest(handler.UpdateImage))
		s.Router.Delete("/images/{objectID}", s.handleAdminRequest(handler.DeleteImage))

		s.Router.Post("/links", s.handleAdminRequest(handler.CreateLink))
		s.Router.Patch("/links/{objectID}", s.handleAdminRequest(handler.UpdateLink))
		s.Router.Delete("/links/{objectID}", s.handleAdminRequest(handler.DeleteLink))

		s.Router.Post("/places", s.handleAdminRequest(handler.CreatePlace))
		s.Router.Patch("/places/{objectID}", s.handleAdminRequest(handler.UpdatePlace))
		s.Router.Delete("/places/{objectID}", s.handleAdminRequest(handler.DeletePlace))

		s.Router.Post("/tags", s.handleAdminRequest(handler.CreateTag))
		s.Router.Patch("/tags/{objectID}", s.handleAdminRequest(handler.UpdateTag))
		s.Router.Delete("/tags/{objectID}", s.handleAdminRequest(handler.DeleteTag))
	}
}

func (s *Server) Run(host string) {

	server := http.Server{
		Addr:      host,
		Handler:   s.Router,
		TLSConfig: s.TLSConfig,
	}

	specifiedInTLSConfig := ""
	if err := server.ListenAndServeTLS(specifiedInTLSConfig, specifiedInTLSConfig); err != nil {
		log.Fatal().Err(err).Str("type", "server").Msg("Failed to run server")
	}
}

// function prototype to inject DB instance in handleRequest()
type RequestHandlerFunction func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func (s *Server) handleRequest(requestHandler RequestHandlerFunction) http.HandlerFunc {

	return authentication.IsEntitled(s.Config.APIKeys, func(w http.ResponseWriter, r *http.Request) {
		requestHandler(s.DB, w, r)
	})
}

func (s *Server) handleAdminRequest(requestHandler RequestHandlerFunction) http.HandlerFunc {

	return authentication.IsEntitled(s.Config.AdminKeys, func(w http.ResponseWriter, r *http.Request) {
		requestHandler(s.DB, w, r)
	})
}

func (s *Server) handleRequestWithoutValidation(requestHandler RequestHandlerFunction) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHandler(s.DB, w, r)
	})
}
