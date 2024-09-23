package main

import (
	"os"

	db "nongki-test/database"
	"nongki-test/database/migration"
	"nongki-test/internal/factory"
	"nongki-test/internal/http"
	"nongki-test/internal/middleware"
	"nongki-test/pkg/util/env"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title nongki-test
// @version 0.0.1
// @description This is a doc for nongki-test.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /

func main() {
	var PORT = os.Getenv("PORT")

	// make sure database should the first init
	db.Init()

	// enabled migration, disabled if doesn't need
	migration.Init()

	// enable elasticsearch if needed
	// elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
