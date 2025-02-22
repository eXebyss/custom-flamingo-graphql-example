package graphql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"FlamingoGQLExample/user/domain"
	"github.com/99designs/gqlgen/graphql"
)

var errNoAttributeReturned = errors.New("no attributes returned")

// UserQueryResolver resolver for the user service
type UserQueryResolver struct {
	userService domain.UserService
}

// Inject dependencies
func (r *UserQueryResolver) Inject(userService domain.UserService) *UserQueryResolver {
	r.userService = userService
	return r
}

// User getter
func (r *UserQueryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	user, err := r.userService.UserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can not load user: %w", err)
	}

	return user, nil
}

// UserAttributeFilter directive
func (r *UserQueryResolver) UserAttributeFilter(ctx context.Context, obj interface{}, next graphql.Resolver, prefix string) (res interface{}, err error) {
	rawAttributes, err := next(ctx)
	if err != nil {
		return nil, err
	}

	attributes, ok := rawAttributes.(domain.Attributes)
	if !ok {
		return nil, errNoAttributeReturned
	}

	for k := range attributes {
		if strings.HasPrefix(k, prefix) {
			delete(attributes, k)
		}
	}

	return attributes, nil
}
