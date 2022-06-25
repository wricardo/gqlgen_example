package gqlgen_example

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/generated"
	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/resolvers"
)

func GetGraphqlHTTPHandler(resolver *resolvers.Resolver, directives generated.DirectiveRoot) http.Handler {
	es := generated.NewExecutableSchema(generated.Config{
		Resolvers:  resolver,
		Directives: directives,
	})
	srv := handler.New(es)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Println("panic:", err)
		debug.PrintStack()

		return gqlerror.Errorf("Internal server error!")
	})

	srv.SetQueryCache(lru.New(lruCacheSize))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(lruCacheSize),
	})
	return srv
}
