package repository

import (
	"github.com/antoha2/task/pkg/logger"
	"gorm.io/gorm"
)

type repositoryImplDB struct {
	rep    *gorm.DB
	logger logger.Logger
	TodolistRep
}

func NewTaskRepository(dbx *gorm.DB, logger logger.Logger) *repositoryImplDB {

	return &repositoryImplDB{
		rep:    dbx,
		logger: logger,
	}
}

type TodolistRep interface {
	Create(*RepTask) error
	Read(*RepFilter) []RepTask
	Update(*RepTask) error
	Delete(*RepTask) error //Delete(*RepFilter) error
}

type RepTask struct {
	TaskId int `gorm:"primaryKey;"` //  index:unique"`
	UserId int
	Text   string
	IsDone bool `gorm:"column:isdone"`
	//CreateAt time.Time
}

func (RepTask) TableName() string {
	return "todolist"
}

type RepFilter struct {
	TaskId    int
	UserId    int
	UserRoles []string
	Ids       []int
	Text      string
	IsDone    bool
	//Tasks  []RepTask
}

func NewDB(dbx *gorm.DB) *repositoryImplDB {

	return &repositoryImplDB{
		rep: dbx,
	}
}
