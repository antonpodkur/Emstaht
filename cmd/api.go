package main

import (
	"log"

	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/server"
	"github.com/antonpodkur/Emstaht/pkg/db/postgres"
	"github.com/antonpodkur/Emstaht/pkg/utils"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath("local")

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err.Error())
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err.Error())
	}

	psqlDb, err := postgres.NewPsqlDb(cfg)
	if err != nil {
		log.Fatalf("Database connection error: %v", err.Error())
	}
	server := server.NewServer(cfg, psqlDb)
	server.Run()
}
