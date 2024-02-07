package web

import (
	"net/http"

	authService "github.com/antoha2/task/pkg/auth"
	taskService "github.com/antoha2/task/service"
)

type Transport interface {
}

type webImpl struct {
	taskService taskService.TodolistService
	authService authService.AuthService
	server      *http.Server
}

func NewWeb(taskService taskService.TodolistService, authService authService.AuthService) *webImpl {
	return &webImpl{
		taskService: taskService,
		authService: authService,
	}
}
