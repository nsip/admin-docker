package main

// Use api and Test
import (
	"fmt"

	// "encoding/xml"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	admindocker "github.com/nsip/admin-docker"
)

func main() {
	fmt.Println("Starting ADMIN::DOCKER")

	doc := admindocker.New()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	})) // allow cors requests during testing

	e.Static("/", "dashboard")

	e.GET("/api/running", func(c echo.Context) error {
		doc.UpdateRunning()
		return c.JSON(http.StatusOK, doc.Containers)
	})

	e.Logger.Fatal(e.Start(":8097"))
}
