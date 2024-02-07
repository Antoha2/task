package endpoints

import (
	"context"
	"errors"
	"log"

	helper "github.com/antoha2/task"
	service "github.com/antoha2/task/service"
	"github.com/go-kit/kit/endpoint"
)

type ReadRequest struct {
	Ids []int `json:"ids"`
}

type ReadResponse struct {
	Tasks []*Task
}

type Task struct {
	TaskId int    `json:"id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

func MakeReadEndpoint(s service.TodolistService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var ok bool
		req := request.(ReadRequest)
		serTask := new(service.SerFilter)

		if serTask.UserId, ok = ctx.Value(helper.USER_ID).(int); !ok {
			newErr := "UserId не найден"
			log.Println(newErr)
			return nil, errors.New(newErr)
		}

		serTask.Ids = make([]int, len(req.Ids))

		for index, id := range req.Ids {
			serTask.Ids[index] = id
		}

		tasks := s.Read(ctx, serTask)

		respTasks := make([]*Task, len(tasks))
		for index, task := range tasks {
			t := &Task{
				TaskId: task.Id,
				UserId: task.UserId,
				Text:   task.Text,
				IsDone: task.IsDone,
			}
			respTasks[index] = t
		}

		return ReadResponse{Tasks: respTasks}, nil
	}
}
