package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/Arismonx/Golang-GraphQl/graph"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	server := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	// Configure HTTP transport
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.POST("/query", func(c echo.Context) error {
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
