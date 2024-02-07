package endpoints

import (
	"context"
	"errors"
	"log"

	helper "github.com/antoha2/task"
	service "github.com/antoha2/task/service"
	"github.com/go-kit/kit/endpoint"
)

type DeleteRequest struct {
	Id int `json:"task_id"`
}

type DeleteResponse struct {
	TaskId int `json:"task_id"`
}

func MakeDeleteEndpoint(s service.TodolistService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var ok bool
		req := request.(DeleteRequest)
		serTask := new(service.SerTask)

		serTask.Id = req.Id
		if serTask.UserId, ok = ctx.Value(helper.USER_ID).(int); !ok {
			newErr := "UserId не найден"
			log.Println(newErr)
			return nil, errors.New(newErr)
		}

		if err := s.Delete(ctx, serTask); err != nil {
			return nil, err
		}
		return DeleteResponse{TaskId: serTask.Id}, nil
	}
}
