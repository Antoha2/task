package endpoints

import (
	service "github.com/antoha2/task/service"

	"github.com/go-kit/kit/endpoint"
)

type TaskEndpoints struct {
	Create endpoint.Endpoint
	Read   endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
}

func MakeTaskEndpoints(s service.TodolistService) *TaskEndpoints {
	return &TaskEndpoints{
		Create: MakeCreateEndpoint(s),
		Read:   MakeReadEndpoint(s),
		Update: MakeUpdateEndpoint(s),
		Delete: MakeDeleteEndpoint(s),
	}
}
