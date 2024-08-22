package dao

import (
	"database/sql"

	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
)

type TaskDao struct {
	db *sql.DB
}

func NewTaskDao(db *sql.DB) repository.TaskRepository {
	return &TaskDao{db}
}

func (d *TaskDao) One(id int) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) List() ([]model.Task, error) {
	return []model.Task{}, nil
}

func (d *TaskDao) Search(condition model.TaskSearchCondition) ([]model.Task, error) {
	return []model.Task{}, nil
}

func (d *TaskDao) Create(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) Update(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) Delete(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}
