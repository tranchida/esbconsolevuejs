package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed dist
var staticFiles embed.FS

type Environment struct {
	Env string `json:"env"`
	Title string `json:"title"`
}

func main() {

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
		return c.JSON(http.StatusOK, []Environment{
			{
				Env: "de",
				Title: "Développement",
			},
			{
				Env: "in",
				Title: "Intégration",
			},
			{
				Env: "va",
				Title: "Validation",
			},
			{
				Env: "pp",
				Title: "Pré-production",
			},
			{
				Env: "pr",
				Title: "Production",
			},
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
