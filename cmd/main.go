package main

import (
	"fmt"
	"net/http"

	"github.com/NikhilSharma03/Mirai/config"
	"github.com/NikhilSharma03/Mirai/internal/app"
	"github.com/NikhilSharma03/Mirai/internal/logger"
	"github.com/NikhilSharma03/Mirai/internal/router"
)

func main() {
	log := logger.NewLogger(logger.GetConfiguredLogger())

	config, err := config.NewConfig()
	if err != nil {
		log.Panic(err)
	}

	app := app.NewApp()

	router := router.NewRouter(app)

	log.Info(fmt.Sprintf("Starting server on PORT %v", config.PORT))

	port := fmt.Sprintf(":%v", config.PORT)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Panic(err)
	}
}
