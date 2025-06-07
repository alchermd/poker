package server

import (
	"context"
	"github.com/alchermd/poker/application"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server encapsulates the running web server that handles HTTP requests.
type Server struct {
	srv             *http.Server
	l               *log.Logger
	shutdownTimeout time.Duration
	templates       templateCache
	app             *application.App
}

// Run starts the web server and handles graceful shutdown.
func (s *Server) Run() {
	s.l.Print("Starting poker server...")
	s.l.Printf("Listening on address %q", s.srv.Addr)
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			s.l.Printf("Server error: %s", err)
		}
	}()

	// Trap and handle exit signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, syscall.SIGTERM)
	sig := <-signalChan
	s.l.Println("Received terminate, gracefully shutting down: ", sig)

	//	Graceful shutdown
	tc, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	s.srv.Shutdown(tc)
	cancel()
}

// NewServer initializes a new server instance.
func NewServer(app *application.App) *Server {
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	l := log.Default()

	s := &Server{
		srv:             srv,
		l:               l,
		shutdownTimeout: 30 * time.Second,
		templates:       make(templateCache),
		app:             app,
	}
	s.RegisterRoutes(r)
	s.ParseTemplates()
	return s
}
