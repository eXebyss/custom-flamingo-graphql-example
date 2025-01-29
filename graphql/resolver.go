// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

//go:build !graphql
// +build !graphql

package graphql

import (
	"FlamingoGQLExample/todo"
	"FlamingoGQLExample/todo/domain"
	domain1 "FlamingoGQLExample/user/domain"
	graphql1 "FlamingoGQLExample/user/interfaces/graphql"
	"context"

	graphql2 "flamingo.me/graphql"
)

var _ ResolverRoot = new(rootResolver)

type rootResolver struct {
	rootResolverMutation *rootResolverMutation
	rootResolverQuery    *rootResolverQuery
	rootResolverUser     *rootResolverUser

	userAttributeFilterResolver *graphql1.UserQueryResolver
}

func (r *rootResolver) Inject(
	rootResolverMutation *rootResolverMutation,
	rootResolverQuery *rootResolverQuery,
	rootResolverUser *rootResolverUser,

	userAttributeFilterResolver *graphql1.UserQueryResolver,
) {
	r.rootResolverMutation = rootResolverMutation
	r.rootResolverQuery = rootResolverQuery
	r.rootResolverUser = rootResolverUser

	r.userAttributeFilterResolver = userAttributeFilterResolver
}

func (r *rootResolver) directives() DirectiveRoot {
	return DirectiveRoot{
		UserAttributeFilter: r.userAttributeFilterResolver.UserAttributeFilter,
	}
}

func (r *rootResolver) Mutation() MutationResolver {
	return r.rootResolverMutation
}
func (r *rootResolver) Query() QueryResolver {
	return r.rootResolverQuery
}
func (r *rootResolver) User() UserResolver {
	return r.rootResolverUser
}

type rootResolverMutation struct {
	resolveFlamingo   func(ctx context.Context) (*string, error)
	resolveTodoAdd    func(ctx context.Context, user string, task string) (*domain.Todo, error)
	resolveTodoDone   func(ctx context.Context, todo string, done bool) (*domain.Todo, error)
	resolveTodoDelete func(ctx context.Context, todo string) (bool, error)
	resolveTodoEdit   func(ctx context.Context, todo string, task string) (*domain.Todo, error)
}

func (r *rootResolverMutation) Inject(
	mutationFlamingo *graphql2.FlamingoQueryResolver,
	mutationTodoAdd *todo.MutationResolver,
	mutationTodoDone *todo.MutationResolver,
	mutationTodoDelete *todo.MutationResolver,
	mutationTodoEdit *todo.MutationResolver,
) {
	r.resolveFlamingo = mutationFlamingo.Flamingo
	r.resolveTodoAdd = mutationTodoAdd.TodoAdd
	r.resolveTodoDone = mutationTodoDone.TodoDone
	r.resolveTodoDelete = mutationTodoDelete.TodoDelete
	r.resolveTodoEdit = mutationTodoEdit.TodoEdit
}

func (r *rootResolverMutation) Flamingo(ctx context.Context) (*string, error) {
	return r.resolveFlamingo(ctx)
}
func (r *rootResolverMutation) TodoAdd(ctx context.Context, user string, task string) (*domain.Todo, error) {
	return r.resolveTodoAdd(ctx, user, task)
}
func (r *rootResolverMutation) TodoDone(ctx context.Context, todo string, done bool) (*domain.Todo, error) {
	return r.resolveTodoDone(ctx, todo, done)
}
func (r *rootResolverMutation) TodoDelete(ctx context.Context, todo string) (bool, error) {
	return r.resolveTodoDelete(ctx, todo)
}
func (r *rootResolverMutation) TodoEdit(ctx context.Context, todo string, task string) (*domain.Todo, error) {
	return r.resolveTodoEdit(ctx, todo, task)
}

type rootResolverQuery struct {
	resolveFlamingo func(ctx context.Context) (*string, error)
	resolveUser     func(ctx context.Context, id string) (*domain1.User, error)
}

func (r *rootResolverQuery) Inject(
	queryFlamingo *graphql2.FlamingoQueryResolver,
	queryUser *graphql1.UserQueryResolver,
) {
	r.resolveFlamingo = queryFlamingo.Flamingo
	r.resolveUser = queryUser.User
}

func (r *rootResolverQuery) Flamingo(ctx context.Context) (*string, error) {
	return r.resolveFlamingo(ctx)
}
func (r *rootResolverQuery) User(ctx context.Context, id string) (*domain1.User, error) {
	return r.resolveUser(ctx, id)
}

type rootResolverUser struct {
	resolveTodos func(ctx context.Context, obj *domain1.User) ([]*domain.Todo, error)
}

func (r *rootResolverUser) Inject(
	userTodos *todo.UserResolver,
) {
	r.resolveTodos = userTodos.Todos
}

func (r *rootResolverUser) Todos(ctx context.Context, obj *domain1.User) ([]*domain.Todo, error) {
	return r.resolveTodos(ctx, obj)
}

func direct(root *rootResolver) map[string]interface{} {
	return map[string]interface{}{
		"Mutation.Flamingo":   root.Mutation().Flamingo,
		"Mutation.TodoAdd":    root.Mutation().TodoAdd,
		"Mutation.TodoDone":   root.Mutation().TodoDone,
		"Mutation.TodoDelete": root.Mutation().TodoDelete,
		"Mutation.TodoEdit":   root.Mutation().TodoEdit,
		"Query.Flamingo":      root.Query().Flamingo,
		"Query.User":          root.Query().User,
		"User.Todos":          root.User().Todos,
	}
}
