package app

import (
	"log"

	"github.com/ddritzenhoff/dinny/pkg/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router             *gin.Engine
	cookService        api.CookService
	participantService api.ParticipantService
	// TODO (ddritzenhoff) should I put the signing secret here? Is there a better way of doing this?
	signingSecret	string
}

func NewServer(router *gin.Engine, cookService api.CookService, participantService api.ParticipantService, signingSecret string) *Server {
	return &Server{
		router:             router,
		cookService:        cookService,
		participantService: participantService,
		signingSecret: signingSecret,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
