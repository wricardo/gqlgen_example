package resolvers

import (
	"context"

	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/dataloader"
)

// ConstructLoaders constructs dataloaders. More info on dataloaders: https://gqlgen.com/reference/dataloaders
func (r *Resolver) ConstructLoaders(ctx context.Context) *dataloader.Loaders {
	return &dataloader.Loaders{
		// UserByID: dataloadgen.NewLoader(
		// 	func(ids []string) (map[string]model.User, error) {
		// 		authContext := auth.FromContext(ctx)
		// 		if authContext == nil {
		// 			return nil, errors2.New("auth context is nil")
		// 		}
		// 		res, err := r.AccountClient.GetManyUsers(ctx, &pbaccount.GetManyUsersRequest{
		// 			Ids: ids,
		// 		})
		// 		if err != nil {
		// 			return nil, errors2.Wrap(err, 0)
		// 		}
		// 		return ToMapWithID(model.UsersFromPb(res.Users)), nil
		// 	},
		// 	dataloadgen.WithBatchCapacity(100),
		// 	dataloadgen.WithWait(1*time.Millisecond),
		// ),
	}
}

type IDer[V comparable] interface {
	GetID() V
}

// ToMapWithID returns a map of ValueT where the key of the map is KeyT
// ValueT must have a pointer receiver method GetID() KeyT
func ToMapWithID[ValueT IDer[KeyT], KeyT string](slice []ValueT) map[KeyT]ValueT {
	res := make(map[KeyT]ValueT)
	for k, v := range slice {
		res[v.GetID()] = slice[k]
	}
	return res
}
