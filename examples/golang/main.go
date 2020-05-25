package main

// Use api and Test
import (
	"fmt"

	// "encoding/xml"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Starting EXAMPLE::DOCKER")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	})) // allow cors requests during testing

	e.GET("/test", func(c echo.Context) error {
		encodedJSON := []byte("<test>a</test>")
		return c.JSONBlob(http.StatusOK, encodedJSON)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
