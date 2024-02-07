package endpoints

import (
	"context"
	"errors"
	"log"

	helper "github.com/antoha2/task"
	service "github.com/antoha2/task/service"
	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	Text string `json:"text"`
}

type CreateResponse struct {
	TaskId int `json:"task_id"`
}

func MakeCreateEndpoint(s service.TodolistService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreateRequest)
		serTask := new(service.SerTask)
		serTask.Text = req.Text

		var ok bool
		if serTask.UserId, ok = ctx.Value(helper.USER_ID).(int); !ok {
			newErr := "UserId не найден"
			log.Println(newErr)
			return nil, errors.New(newErr)
		}
		id, err := s.Create(ctx, serTask)
		if err != nil {
			return nil, err
		}

		/* 	task := &etodo.Task{
			Id:     serTask.Id,
			UserId: serTask.UserId,
			Text:   serTask.Text,
			IsDone: serTask.IsDone,
		}
		*/

		//log.Println("!!!!!!!!!!!!!!!!! - ", id)
		//t := fmt.Sprintf("создана запись - %v", task)
		return CreateResponse{TaskId: id}, nil
	}
}
