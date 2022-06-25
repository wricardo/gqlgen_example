package dataloader

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mshaeon/dataloadgen"
	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/model"
)

type loadersKeyType string

const loadersKey loadersKeyType = "dataloaders"

type DataLoaderProvider interface {
	ConstructLoaders(ctx context.Context) *Loaders
}

type Loaders struct {
	UserByID *dataloadgen.Loader[string, model.User]
}

func Middleware(dataloaderProvider DataLoaderProvider) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		ctx := context.WithValue(gctx.Request.Context(), loadersKey, dataloaderProvider.ConstructLoaders(gctx.Request.Context()))
		gctx.Request = gctx.Request.WithContext(ctx)
		gctx.Next()
	}
}

func For(ctx context.Context) *Loaders {
	val, ok := ctx.Value(loadersKey).(*Loaders)
	if !ok {
		panic("ctx value is not a Loader")
	}
	return val
}
