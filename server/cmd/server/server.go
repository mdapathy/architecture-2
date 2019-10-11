package main

import (
	"context"
	"fmt"
	"github.com/mdapathy/architecture-2/server/forums"
	"github.com/mdapathy/architecture-2/server/users"
	"net/http"
)

type HttpPortNumber int

// ChatApiServer configures necessary handlers and starts listening on a configured port.
type ChatApiServer struct {
	Port HttpPortNumber

	ForumsHandler forums.HttpHandlerFunc
	UsersHandler  users.HttpHandlerFunc

	server *http.Server
}

func (s *ChatApiServer) Start() error {

	if s.ForumsHandler == nil || s.UsersHandler == nil {
		return fmt.Errorf(" HTTP handlers are not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)

	handler.HandleFunc("/forums", s.ForumsHandler)
	handler.HandleFunc("/user", s.UsersHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down the previously started HTTP server.
func (s *ChatApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
