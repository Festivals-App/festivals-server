package server

import (
	"database/sql"
	"fmt"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/Festivals-App/festivals-server/server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// Server has router and db instances
type Server struct {
	Router *chi.Mux
	DB     *sql.DB
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
		log.Fatal("server initialize: could not connect to database")
	}

	s.DB = db
	s.Router = chi.NewRouter()

	s.setMiddleware()
	s.setWalker()
	s.setRoutes(config)
}

func (s *Server) setMiddleware() {
	// tell the router which middleware to use
	s.Router.Use(
		// used to log the request to the console | development
		middleware.Logger,
		// helps to redirect wrong requests (why do one want that?)
		//middleware.RedirectSlashes,
		// tries to recover after panics (?)
		middleware.Recoverer,
	)
}

func (s *Server) setWalker() {

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s \n", method, route)
		return nil
	}
	if err := chi.Walk(s.Router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
}

// setRouters sets the all required routers
func (s *Server) setRoutes(config *config.Config) {

	if config.ReadOnly == true {

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
		s.Router.Get("/events/{objectID}/artist", s.handleRequest(handler.GetEventArtist))
		s.Router.Get("/events/{objectID}/location", s.handleRequest(handler.GetEventLocation))

	} else {
		// Festival routes
		s.Router.Get("/festivals", s.handleRequest(handler.GetFestivals))
		s.Router.Get("/festivals/{objectID}", s.handleRequest(handler.GetFestival))
		s.Router.Post("/festivals", s.handleRequest(handler.CreateFestival))
		s.Router.Patch("/festivals/{objectID}", s.handleRequest(handler.UpdateFestival))
		s.Router.Delete("/festivals/{objectID}", s.handleRequest(handler.DeleteFestival))
		s.Router.Get("/festivals/{objectID}/events", s.handleRequest(handler.GetFestivalEvents))
		s.Router.Get("/festivals/{objectID}/image", s.handleRequest(handler.GetFestivalImage))
		s.Router.Get("/festivals/{objectID}/links", s.handleRequest(handler.GetFestivalLinks))
		s.Router.Get("/festivals/{objectID}/place", s.handleRequest(handler.GetFestivalPlace))
		s.Router.Get("/festivals/{objectID}/tags", s.handleRequest(handler.GetFestivalTags))
		s.Router.Post("/festivals/{objectID}/events/{resourceID}", s.handleRequest(handler.SetEventForFestival))
		s.Router.Post("/festivals/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForFestival))
		s.Router.Post("/festivals/{objectID}/links/{resourceID}", s.handleRequest(handler.SetLinkForFestival))
		s.Router.Post("/festivals/{objectID}/place/{resourceID}", s.handleRequest(handler.SetPlaceForFestival))
		s.Router.Post("/festivals/{objectID}/tags/{resourceID}", s.handleRequest(handler.SetTagForFestival))
		s.Router.Delete("/festivals/{objectID}/image/{resourceID}", s.handleRequest(handler.RemoveImageForFestival))
		s.Router.Delete("/festivals/{objectID}/links/{resourceID}", s.handleRequest(handler.RemoveLinkForFestival))
		s.Router.Delete("/festivals/{objectID}/place/{resourceID}", s.handleRequest(handler.RemovePlaceForFestival))
		s.Router.Delete("/festivals/{objectID}/tags/{resourceID}", s.handleRequest(handler.RemoveTagForFestival))

		// Artist routes
		s.Router.Get("/artists", s.handleRequest(handler.GetArtists))
		s.Router.Get("/artists/{objectID}", s.handleRequest(handler.GetArtist))
		s.Router.Post("/artists", s.handleRequest(handler.CreateArtist))
		s.Router.Patch("/artists/{objectID}", s.handleRequest(handler.UpdateArtist))
		s.Router.Delete("/artists/{objectID}", s.handleRequest(handler.DeleteArtist))
		s.Router.Get("/artists/{objectID}/image", s.handleRequest(handler.GetArtistImage))
		s.Router.Get("/artists/{objectID}/links", s.handleRequest(handler.GetArtistLinks))
		s.Router.Get("/artists/{objectID}/tags", s.handleRequest(handler.GetArtistTags))
		s.Router.Post("/artists/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForArtist))
		s.Router.Post("/artists/{objectID}/links/{resourceID}", s.handleRequest(handler.SetLinkForArtist))
		s.Router.Post("/artists/{objectID}/tags/{resourceID}", s.handleRequest(handler.SetTagForArtist))
		s.Router.Delete("/artists/{objectID}/image/{resourceID}", s.handleRequest(handler.RemoveImageForArtist))
		s.Router.Delete("/artists/{objectID}/links/{resourceID}", s.handleRequest(handler.RemoveLinkForArtist))
		s.Router.Delete("/artists/{objectID}/tags/{resourceID}", s.handleRequest(handler.RemoveTagForArtist))

		// Location routes
		s.Router.Get("/locations", s.handleRequest(handler.GetLocations))
		s.Router.Get("/locations/{objectID}", s.handleRequest(handler.GetLocation))
		s.Router.Post("/locations", s.handleRequest(handler.CreateLocation))
		s.Router.Patch("/locations/{objectID}", s.handleRequest(handler.UpdateLocation))
		s.Router.Delete("/locations/{objectID}", s.handleRequest(handler.DeleteLocation))
		s.Router.Get("/locations/{objectID}/image", s.handleRequest(handler.GetLocationImage))
		s.Router.Get("/locations/{objectID}/links", s.handleRequest(handler.GetLocationLinks))
		s.Router.Get("/locations/{objectID}/place", s.handleRequest(handler.GetLocationPlace))
		s.Router.Post("/locations/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForLocation))
		s.Router.Post("/locations/{objectID}/links/{resourceID}", s.handleRequest(handler.SetLinkForLocation))
		s.Router.Post("/locations/{objectID}/place/{resourceID}", s.handleRequest(handler.SetPlaceForLocation))
		s.Router.Delete("/locations/{objectID}/image/{resourceID}", s.handleRequest(handler.RemoveImageForLocation))
		s.Router.Delete("/locations/{objectID}/links/{resourceID}", s.handleRequest(handler.RemoveLinkForLocation))
		s.Router.Delete("/locations/{objectID}/place/{resourceID}", s.handleRequest(handler.RemovePlaceForLocation))

		// Event routes
		s.Router.Get("/events", s.handleRequest(handler.GetEvents))
		s.Router.Get("/events/{objectID}", s.handleRequest(handler.GetEvent))
		s.Router.Post("/events", s.handleRequest(handler.CreateEvent))
		s.Router.Patch("/events/{objectID}", s.handleRequest(handler.UpdateEvent))
		s.Router.Delete("/events/{objectID}", s.handleRequest(handler.DeleteEvent))
		s.Router.Get("/events/{objectID}/festival", s.handleRequest(handler.GetEventFestival))
		s.Router.Get("/events/{objectID}/artist", s.handleRequest(handler.GetEventArtist))
		s.Router.Get("/events/{objectID}/location", s.handleRequest(handler.GetEventLocation))
		s.Router.Post("/events/{objectID}/artist/{resourceID}", s.handleRequest(handler.SetArtistForEvent))
		s.Router.Post("/events/{objectID}/location/{resourceID}", s.handleRequest(handler.SetLocationForEvent))
		s.Router.Delete("/events/{objectID}/artist/{resourceID}", s.handleRequest(handler.RemoveArtistForEvent))
		s.Router.Delete("/events/{objectID}/location/{resourceID}", s.handleRequest(handler.RemoveLocationForEvent))

		// Image routes
		s.Router.Get("/images", s.handleRequest(handler.GetImages))
		s.Router.Get("/images/{objectID}", s.handleRequest(handler.GetImage))
		s.Router.Post("/images", s.handleRequest(handler.CreateImage))
		s.Router.Patch("/images/{objectID}", s.handleRequest(handler.UpdateImage))
		s.Router.Delete("/images/{objectID}", s.handleRequest(handler.DeleteImage))

		// Link routes
		s.Router.Get("/links", s.handleRequest(handler.GetLinks))
		s.Router.Get("/links/{objectID}", s.handleRequest(handler.GetLink))
		s.Router.Post("/links", s.handleRequest(handler.CreateLink))
		s.Router.Patch("/links/{objectID}", s.handleRequest(handler.UpdateLink))
		s.Router.Delete("/links/{objectID}", s.handleRequest(handler.DeleteLink))

		// Place routes
		s.Router.Get("/places", s.handleRequest(handler.GetPlaces))
		s.Router.Get("/places/{objectID}", s.handleRequest(handler.GetPlace))
		s.Router.Post("/places", s.handleRequest(handler.CreatePlace))
		s.Router.Patch("/places/{objectID}", s.handleRequest(handler.UpdatePlace))
		s.Router.Delete("/places/{objectID}", s.handleRequest(handler.DeletePlace))

		// Tag routes
		s.Router.Get("/tags", s.handleRequest(handler.GetTags))
		s.Router.Get("/tags/{objectID}", s.handleRequest(handler.GetTag))
		s.Router.Get("/tags/{objectID}/festivals", s.handleRequest(handler.GetTagFestivals))
		s.Router.Post("/tags", s.handleRequest(handler.CreateTag))
		s.Router.Patch("/tags/{objectID}", s.handleRequest(handler.UpdateTag))
		s.Router.Delete("/tags/{objectID}", s.handleRequest(handler.DeleteTag))
	}
}

// Run the server on it's router
func (s *Server) Run(host string) {
	log.Fatal(http.ListenAndServe(host, s.Router))
}

// function prototype to inject DB instance in handleRequest()
type RequestHandlerFunction func(db *sql.DB, w http.ResponseWriter, r *http.Request)

// inject DB in handler functions
func (s *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(s.DB, w, r)
	}
}
