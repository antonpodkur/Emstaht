package server

import (
	"github.com/antonpodkur/Emstaht/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	cfg *config.Config
	db *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{
		cfg: cfg,
		db: db,
	}
}

func (s *Server) Run() {
	r := gin.Default()

	s.MapHandlers(r)

	r.Run(":8000")
}
