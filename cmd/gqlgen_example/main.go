package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wricardo/gqlgen_example/internal/gqlgen_example"
	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/dataloader"
	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/resolvers"
)

func main() {
	spew.Config.Indent = "\t"
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	_ = godotenv.Load()

	// cfg := gqlgen_example.ServiceConfig{}

	resolver := resolvers.NewResolver()

	// use the gin framework to setup routes and middlewares
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthcheck := router.Group("/healthcheck")
	healthcheck.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "healthy")
	})

	srv := gqlgen_example.GetGraphqlHTTPHandler(resolver, gqlgen_example.DefaultDirectives())

	router.GET("/playground", gin.WrapF(playground.Handler("GraphQL playground", "/graphql")))
	router.Any("/graphql", dataloader.Middleware(resolver), gin.WrapH(srv))

	// router.Run will listen on a port for http requests, this call blocks.
	return router.Run(":" + os.Getenv("LISTEN_PORT"))
}
