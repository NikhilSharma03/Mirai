package main

import (
	"fmt"
	"net/http"

	"github.com/NikhilSharma03/Mirai/config"
	"github.com/NikhilSharma03/Mirai/internal/app"
	"github.com/NikhilSharma03/Mirai/internal/logger"
	"github.com/NikhilSharma03/Mirai/internal/repository"
	"github.com/NikhilSharma03/Mirai/internal/router"
)

func main() {
	log := logger.NewLogger(logger.GetConfiguredLogger())

	config, err := config.NewConfig()
	if err != nil {
		log.Panic(err)
	}

	log.Info("Connecting Database")

	dbConn, err := repository.ConnectDB(config)
	if err != nil {
		log.Panic(err)
	}

	_ = repository.NewDAO(dbConn)

	app := app.NewApp()

	router := router.NewRouter(app)

	log.Info(fmt.Sprintf("Starting server on PORT %v", config.PORT))

	if err = http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), router); err != nil {
		log.Panic(err)
	}
}
