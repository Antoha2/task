package endpoints

import (
	"context"
	"errors"
	"log"

	helper "github.com/antoha2/task"
	service "github.com/antoha2/task/service"
	"github.com/go-kit/kit/endpoint"
)

type UpdateRequest struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

type UpdateResponse struct {
	TaskId int `json:"task_id"`
}

func MakeUpdateEndpoint(s service.TodolistService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var ok bool
		req := request.(UpdateRequest)
		serTask := new(service.SerTask)

		if serTask.UserId, ok = ctx.Value(helper.USER_ID).(int); !ok {
			newErr := "UserId не найден"
			log.Println(newErr)
			return nil, errors.New(newErr)
		}

		serTask.Id = req.Id
		serTask.Text = req.Text
		serTask.IsDone = req.IsDone

		if err := s.Update(ctx, serTask); err != nil {
			return nil, err
		}
		return UpdateResponse{TaskId: serTask.Id}, nil
	}
}
