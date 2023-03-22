package main

import (
	"log"

	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/server"
	"github.com/antonpodkur/Emstaht/pkg/db/postgres"
	"github.com/antonpodkur/Emstaht/pkg/utils"
)

// @title Emstaht
// @version         1.0
// @description    This is my project
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
