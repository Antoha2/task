package service

import (
	"context"

	taskRepository "github.com/antoha2/task/service/taskRepository"
)

type TodolistService interface {
	Create(ctx context.Context, task *SerTask) (int, error)
	Read(ctx context.Context, task *SerFilter) []*SerTask
	Delete(ctx context.Context, task *SerTask) error
	Update(ctx context.Context, task *SerTask) error
}

type serviceImpl struct {
	repository taskRepository.TodolistRep
}

func NewTaskService(rep taskRepository.TodolistRep) *serviceImpl {
	return &serviceImpl{
		repository: rep,
	}
}

type SerTask struct {
	Id     int
	UserId int
	Text   string
	IsDone bool
}

type SerFilter struct {
	UserId int
	Ids    []int
	Text   string
	IsDone bool
}
