package dao

import (
	"encoding/json"

	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
)

type TaskDao struct {
	buf []byte
}

type TaskJSON struct {
	Tasks []model.Task `json:"tasks"`
}

func NewTaskDao(buf []byte) repository.TaskRepository {
	return &TaskDao{buf}
}

func (d *TaskDao) One(id int) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) List() ([]model.Task, error) {
	var taskJSON TaskJSON
	err := json.Unmarshal(d.buf, &taskJSON)
	if err != nil {
		return []model.Task{}, err
	}
	return taskJSON.Tasks, nil
}

func (d *TaskDao) Statuses() ([]model.TaskStatus, error) {
	return model.TASK_STATUSES, nil
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
