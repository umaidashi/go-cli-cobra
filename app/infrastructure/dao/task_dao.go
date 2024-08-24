package dao

import (
	"encoding/json"
	"io"
	"os"

	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
)

type TaskDao struct {
	file *os.File
}

type TaskJSON struct {
	Tasks []model.Task `json:"tasks"`
}

func NewTaskDao(file *os.File) repository.TaskRepository {
	return &TaskDao{file}
}

func (d *TaskDao) One(id int) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) List() ([]model.Task, error) {
	taskJSON, err := d.getTaskJSON()
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
	taskJSON, err := d.getTaskJSON()
	taskJSON.Tasks = append(taskJSON.Tasks, task)
	json, err := json.Marshal(taskJSON)
	if err != nil {
		return model.Task{}, err
	}

	err = d.truncateAndWrite(json)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (d *TaskDao) Update(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) Delete(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) getTaskJSON() (TaskJSON, error) {
	var taskJSON TaskJSON
	buf, err := io.ReadAll(d.file)
	err = json.Unmarshal(buf, &taskJSON)
	if err != nil {
		return TaskJSON{}, err
	}
	return taskJSON, nil
}

func (d *TaskDao) truncateAndWrite(buf []byte) error {
	err := d.file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = d.file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = d.file.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
