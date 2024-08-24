package dao

import (
	"encoding/json"
	"errors"

	myJson "github.com/umaidashi/go-cli-cobra/app/infrastructure/json"

	"github.com/samber/lo"
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
)

type TaskDao struct {
	json myJson.JSON
}

func NewTaskDao(json myJson.JSON) repository.TaskRepository {
	return &TaskDao{json}
}

func (d *TaskDao) One(id int) (model.Task, error) {
	targetTask, ok := lo.Find(d.json.Tasks, func(t model.Task) bool {
		return t.Id == id
	})
	if !ok {
		return model.Task{}, errors.New("task not found.")
	}
	return targetTask, nil
}

func (d *TaskDao) List() ([]model.Task, error) {
	return d.json.Tasks, nil
}

func (d *TaskDao) Statuses() ([]model.TaskStatus, error) {
	return model.TASK_STATUSES, nil
}

func (d *TaskDao) Search(condition model.TaskSearchCondition) ([]model.Task, error) {
	return []model.Task{}, nil
}

func (d *TaskDao) Create(task model.Task) (model.Task, error) {
	task.Id = d.json.GetMaxTaskId() + 1

	d.json.Tasks = append(d.json.Tasks, task)
	json, err := json.Marshal(d.json)
	if err != nil {
		return model.Task{}, err
	}
	err = d.json.Write(json)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (d *TaskDao) Update(task model.Task) (model.Task, error) {
	targetTask, ok := lo.Find(d.json.Tasks, func(t model.Task) bool {
		return t.Id == task.Id
	})
	if !ok {
		return model.Task{}, errors.New("task not found.")
	}

	targetTask.Title = task.Title
	targetTask.Content = task.Content
	targetTask.Status = task.Status
	targetTask.UpdatedAt = task.UpdatedAt
	targetTask.CompletedAt = task.CompletedAt

	d.json.Tasks = lo.Map(d.json.Tasks, func(t model.Task, _ int) model.Task {
		if t.Id == task.Id {
			return targetTask
		}
		return t
	})

	json, err := json.Marshal(d.json)
	if err != nil {
		return model.Task{}, err
	}
	err = d.json.Write(json)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (d *TaskDao) Delete(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}
