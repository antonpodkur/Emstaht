package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	r := gin.Default()

	s.MapHandlers(r)

	r.Run(":8000")
}
