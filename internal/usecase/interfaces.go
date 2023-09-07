// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Task
	Task interface {
		CreateTask(context.Context, *CreateTaskRequest) error
		// List(context.Context) ([]*entity.Task, error)
		// Get(context.Context, string) (*entity.Task, error)
		// Update(context.Context, string, *entity.Task) error
		// Delete(context.Context, string) error
	}

	// TaskRepo -.
	TaskRepo interface {
		CreateTask(context.Context, *entity.Task) error
		// List(context.Context) ([]*entity.Task, error)
		// Get(context.Context, string) (*entity.Task, error)
		// Update(context.Context, string, *entity.Task) error
		// Delete(context.Context, string) error
	}
)
