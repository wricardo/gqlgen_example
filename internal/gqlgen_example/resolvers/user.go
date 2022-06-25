package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/generated"
	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (bool, error) {
	return true, nil
}

func (r *organizationResolver) CreatedBy(ctx context.Context, obj *model.Organization) (*model.User, error) {
	return &model.User{
		ID:        "hardcoded123",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
		Organization: &model.Organization{
			ID:        "abc",
			Name:      "my org",
			Deleted:   false,
			CreatedBy: &model.User{},
			UpdatedBy: &model.User{},
			CreatedAt: time.Time{},
			UpdatedAt: &time.Time{},
		},
		Deleted:   false,
		CreatedBy: &model.User{},
		UpdatedBy: &model.User{},
		CreatedAt: time.Time{},
		UpdatedAt: &time.Time{},
	}, nil
}

func (r *organizationResolver) UpdatedBy(ctx context.Context, obj *model.Organization) (*model.User, error) {
	return &model.User{
		ID:        "hardcoded123",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
		Organization: &model.Organization{
			ID:        "abc",
			Name:      "my org",
			Deleted:   false,
			CreatedBy: &model.User{},
			UpdatedBy: &model.User{},
			CreatedAt: time.Time{},
			UpdatedAt: &time.Time{},
		},
		Deleted:   false,
		CreatedBy: &model.User{},
		UpdatedBy: &model.User{},
		CreatedAt: time.Time{},
		UpdatedAt: &time.Time{},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context, page *int, size *int) (*model.UsersPage, error) {
	return &model.UsersPage{
		Users: []model.User{
			model.User{
				ID:        "hardcoded123",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@test.com",
				Organization: &model.Organization{
					ID:        "abc",
					Name:      "my org",
					Deleted:   false,
					CreatedBy: &model.User{},
					UpdatedBy: &model.User{},
					CreatedAt: time.Time{},
					UpdatedAt: &time.Time{},
				},
				Deleted:   false,
				CreatedBy: &model.User{},
				UpdatedBy: &model.User{},
				CreatedAt: time.Time{},
				UpdatedAt: &time.Time{},
			},
		},
		Pagination: &model.Pagination{
			Page:          1,
			Size:          1,
			TotalElements: 1,
			TotalPages:    1,
			First:         true,
			Last:          true,
		},
	}, nil
}

func (r *userResolver) Organization(ctx context.Context, obj *model.User) (*model.Organization, error) {
	return &model.Organization{
		ID:        "abc",
		Name:      "my org",
		Deleted:   false,
		CreatedBy: &model.User{},
		UpdatedBy: &model.User{},
		CreatedAt: time.Time{},
		UpdatedAt: &time.Time{},
	}, nil
}

func (r *userResolver) CreatedBy(ctx context.Context, obj *model.User) (*model.User, error) {
	return &model.User{
		ID:        "hardcoded123",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
		Organization: &model.Organization{
			ID:        "abc",
			Name:      "my org",
			Deleted:   false,
			CreatedBy: &model.User{},
			UpdatedBy: &model.User{},
			CreatedAt: time.Time{},
			UpdatedAt: &time.Time{},
		},
		Deleted:   false,
		CreatedBy: &model.User{},
		UpdatedBy: &model.User{},
		CreatedAt: time.Time{},
		UpdatedAt: &time.Time{},
	}, nil
}

func (r *userResolver) UpdatedBy(ctx context.Context, obj *model.User) (*model.User, error) {
	return &model.User{
		ID:        "hardcoded123",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
		Organization: &model.Organization{
			ID:        "abc",
			Name:      "my org",
			Deleted:   false,
			CreatedBy: &model.User{},
			UpdatedBy: &model.User{},
			CreatedAt: time.Time{},
			UpdatedAt: &time.Time{},
		},
		Deleted:   false,
		CreatedBy: &model.User{},
		UpdatedBy: &model.User{},
		CreatedAt: time.Time{},
		UpdatedAt: &time.Time{},
	}, nil
}

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type organizationResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
