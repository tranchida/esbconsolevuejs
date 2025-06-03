package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"esbconsolevuejs/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed dist
var staticFiles embed.FS

func main() {

	conf, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	var config models.Config
	err = json.Unmarshal(conf, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config file: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.StaticWithConfig(
		middleware.StaticConfig{
			HTML5:      true,
			Root:       "dist",
			Filesystem: http.FS(staticFiles),
		},
	))

	e.GET("/api/environments", func(c echo.Context) error {

			environments := []models.FrontendEnvironment{}
		for _, env := range config.Environments {
			environments = append(environments, models.FrontendEnvironment{
				Env: env.Env,
				Title: env.Title,
			})
		}
		return c.JSON(http.StatusOK, environments)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
