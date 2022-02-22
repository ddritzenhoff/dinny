package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// group all routes under /v1/api
	v1 := router.Group("/v1/")
	{
		v1.GET("/ping", s.Pong())
		v1.POST("/event", s.HandleEvent())
	}

	return router
}
