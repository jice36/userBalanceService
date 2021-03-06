package main

import (
	"flag"
	"log"

	"github.com/userBalanceService/internal/conversion"
	"github.com/userBalanceService/internal/server"
	"github.com/userBalanceService/internal/service"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.yml", "path to config file")
}

func main() {
	flag.Parse()

	config, cErr := server.NewConfig(configPath)
	if cErr != nil {
		log.Fatal(cErr)
	}

	db := config.Database

	service, dErr := service.NewService(db.Dbname, db.Dbpassword, db.Dbuser)
	if dErr != nil {
		log.Fatal(dErr)
	}

	conv := conversion.NewService()

	s := server.NerServer(config, service, conv)

	sErr := s.StartServer()
	if sErr != nil {
		s.Log.Fatal(sErr)
	}
}
