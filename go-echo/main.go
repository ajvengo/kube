package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
        port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

// package main

// import (
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func main() {

// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/", func(c echo.Context) error {
// 		return c.HTML(http.StatusOK, "Hello, Docker! <3")
// 	})

// 	e.GET("/ping", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
// 	})

// 	httpPort := os.Getenv("HTTP_PORT")
// 	if httpPort == "" {
// 		httpPort = "8080"
// 	}

// 	e.Logger.Fatal(e.Start(":" + httpPort))
// }