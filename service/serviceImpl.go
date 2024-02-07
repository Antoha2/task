package service

import (
	"context"
	"fmt"
	"log"

	helper "github.com/antoha2/task"
	taskRepository "github.com/antoha2/task/service/taskRepository"
)

//Create
func (s *serviceImpl) Create(ctx context.Context, task *SerTask) (int, error) {

	repTask := new(taskRepository.RepTask)
	repTask.Text = task.Text
	//repTask.CreateAt = time.Now()
	repTask.IsDone = false
	repTask.UserId = task.UserId

	err := s.repository.Create(repTask)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	task.Id = repTask.TaskId
	task.IsDone = repTask.IsDone

	return task.Id, nil
}

//Read
func (s *serviceImpl) Read(ctx context.Context, task *SerFilter) []*SerTask {

	userRoles, ok := ctx.Value(helper.USER_ROLE).([]string)
	if !ok {
		newErr := "read.userRoles не найден"
		log.Println(newErr)
		return nil
	}

	//log.Println("ser -", task.UserId)
	readFilter := new(taskRepository.RepFilter)
	readFilter.Ids = task.Ids
	readFilter.UserId = task.UserId
	for _, role := range userRoles {
		if role == "admin" {
			readFilter.UserId = 0
			break
		}
	}

	tasks := s.repository.Read(readFilter)
	sliceTask := make([]*SerTask, len(tasks))

	for index, task := range tasks {
		t := &SerTask{
			Id:     task.TaskId,
			UserId: task.UserId,
			Text:   task.Text,
			IsDone: task.IsDone,
		}
		sliceTask[index] = t
	}
	return sliceTask

}

//Delete
func (s *serviceImpl) Delete(ctx context.Context, task *SerTask) error {

	userRoles, ok := ctx.Value(helper.USER_ROLE).([]string)
	if !ok {
		newErr := "delete.userRoles не найден"
		log.Println(newErr)
		return nil
	}

	delFilter := new(taskRepository.RepTask)
	delFilter.TaskId = task.Id
	delFilter.UserId = task.UserId

	for _, role := range userRoles {
		if role == "admin" {
			delFilter.UserId = 0
			break
		}
	}

	err := s.repository.Delete(delFilter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//Update
func (s *serviceImpl) Update(ctx context.Context, task *SerTask) error {

	userRoles, ok := ctx.Value(helper.USER_ROLE).([]string)
	if !ok {
		newErr := "update.userRoles не найден"
		log.Println(newErr)
		return nil
	}

	upFilter := new(taskRepository.RepTask)
	upFilter.TaskId = task.Id
	upFilter.IsDone = task.IsDone
	upFilter.Text = task.Text
	upFilter.UserId = task.UserId

	for _, role := range userRoles {
		if role == "admin" {
			upFilter.UserId = 0
			break
		}
	}

	err := s.repository.Update(upFilter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	task.UserId = upFilter.UserId

	return nil
}
