package infrastructure

import (
	"context"

	"FlamingoGQLExample/user/domain"
)

// UserServiceImpl default implementation
type UserServiceImpl struct{}

// UserByID mocked user getter
func (us *UserServiceImpl) UserByID(_ context.Context, id string) (*domain.User, error) {
	return &domain.User{
		Name:      "User " + id,
		Nicknames: []string{"nick", id},
		Attributes: map[string]interface{}{
			"movie":        "starwars",
			"secret_crush": "r2d2",
		},
	}, nil
}
