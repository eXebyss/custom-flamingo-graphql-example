package todo

import (
	"context"
	"fmt"

	"FlamingoGQLExample/todo/domain"
	"FlamingoGQLExample/todo/infrastructure"
	userDomain "FlamingoGQLExample/user/domain"
)

// UserResolver mapper between graphql and the todos backend
type UserResolver struct {
	todosBackend *infrastructure.TodoService
}

// Inject dependencies
func (r *UserResolver) Inject(todosBackend *infrastructure.TodoService) *UserResolver {
	r.todosBackend = todosBackend
	return r
}

// Todos getter
func (r *UserResolver) Todos(ctx context.Context, obj *userDomain.User) ([]*domain.Todo, error) {
	todos, err := r.todosBackend.Todos(ctx, obj.Name)
	if err != nil {
		return nil, fmt.Errorf("can not load todos: %w", err)
	}

	return todos, nil
}

// MutationResolver maps mutations
type MutationResolver struct {
	resolver *UserResolver
}

// Inject dependencies
func (r *MutationResolver) Inject(resolver *UserResolver) *MutationResolver {
	r.resolver = resolver
	return r
}

// TodoAdd mutation
func (r *MutationResolver) TodoAdd(ctx context.Context, user string, task string) (*domain.Todo, error) {
	todo, err := r.resolver.todosBackend.AddTodo(ctx, user, task)
	if err != nil {
		return nil, fmt.Errorf("can not add todo: %w", err)
	}

	return todo, nil
}

// TodoDone mutation
func (r *MutationResolver) TodoDone(ctx context.Context, todo string, done bool) (*domain.Todo, error) {
	todoDone, err := r.resolver.todosBackend.TodoDone(ctx, todo, done)
	if err != nil {
		return nil, fmt.Errorf("can not update todo: %w", err)
	}

	return todoDone, nil
}

// TodoEdit mutation
func (r *MutationResolver) TodoDelete(ctx context.Context, todo string) (bool, error) {
	success, err := r.resolver.todosBackend.DeleteTodo(ctx, todo)
	if err != nil {
		return false, fmt.Errorf("can not delete todo: %w", err)
	}
	return success, nil
}

// TodoEdit mutation
func (r *MutationResolver) TodoEdit(ctx context.Context, todo string, task string) (*domain.Todo, error) {
	editedTodo, err := r.resolver.todosBackend.TodoEdit(ctx, todo, task)
	if err != nil {
		return nil, fmt.Errorf("can not edit todo: %w", err)
	}

	return editedTodo, nil
}
