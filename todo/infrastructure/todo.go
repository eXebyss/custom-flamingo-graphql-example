package infrastructure

import (
	"context"
	"errors"
	"strconv"

	"FlamingoGQLExample/todo/domain"
)

// TodoService service definition
type TodoService struct{}

var todos = []*domain.Todo{
	{ID: "task-0", Task: "task a"},
	{ID: "task-1", Task: "task b"},
	{ID: "task-2", Task: "task c"},
}

var (
	ErrNoTodoGiven  = errors.New("no todo given")
	ErrTodoNotFound = errors.New("todo not found")
)

// Todos returns a list of mocked todos
func (ts *TodoService) Todos(_ context.Context, _ string) ([]*domain.Todo, error) {
	return todos, nil
}

// AddTodo mutation adds an entry to the list
func (ts *TodoService) AddTodo(_ context.Context, _ string, task string) (*domain.Todo, error) {
	if task == "" {
		return nil, ErrNoTodoGiven
	}

	todo := &domain.Todo{
		ID:   "task-" + strconv.Itoa(len(todos)),
		Task: task,
	}
	todos = append(todos, todo)

	return todo, nil
}

// TodoDone marks a task as finished
func (ts *TodoService) TodoDone(_ context.Context, todoID string, done bool) (*domain.Todo, error) {
	for i, todo := range todos {
		if todo.ID == todoID {
			todos[i].Done = done
			return todos[i], nil
		}
	}

	return nil, ErrTodoNotFound
}

// DeleteTodo deletes a todo
func (ts *TodoService) DeleteTodo(_ context.Context, todoID string) (bool, error) {
	for i, todo := range todos {
		if todo.ID == todoID {
			todos = append(todos[:i], todos[i+1:]...)
			return true, nil
		}
	}
	return false, ErrTodoNotFound
}

// TodoEdit updates the task text of a todo
func (ts *TodoService) TodoEdit(_ context.Context, todoID string, task string) (*domain.Todo, error) {
	if task == "" {
		return nil, ErrNoTodoGiven
	}

	for i, todo := range todos {
		if todo.ID == todoID {
			todos[i].Task = task
			return todos[i], nil
		}
	}

	return nil, ErrTodoNotFound
}
