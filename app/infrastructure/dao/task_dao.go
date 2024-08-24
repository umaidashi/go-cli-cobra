package dao

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"github.com/samber/lo"
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
)

type TaskDao struct {
	file     *os.File
	taskJSON TaskJSON
}

type TaskJSON struct {
	Tasks []model.Task `json:"tasks"`
}

func (j TaskJSON) getMaxTaskId() int {
	taskIds := lo.Map(j.Tasks, func(t model.Task, _ int) int {
		return t.Id
	})
	maxId := lo.MaxBy(taskIds, func(id int, max int) bool {
		return id > max
	})
	return maxId
}

func NewTaskDao(file *os.File) repository.TaskRepository {
	return &TaskDao{file, TaskJSON{}}
}

func (d *TaskDao) One(id int) (model.Task, error) {
	err := d.setTaskJSON()
	if err != nil {
		return model.Task{}, err
	}

	targetTask, ok := lo.Find(d.taskJSON.Tasks, func(t model.Task) bool {
		return t.Id == id
	})
	if !ok {
		return model.Task{}, errors.New("task not found.")
	}
	return targetTask, nil
}

func (d *TaskDao) List() ([]model.Task, error) {
	err := d.setTaskJSON()
	if err != nil {
		return []model.Task{}, err
	}
	return d.taskJSON.Tasks, nil
}

func (d *TaskDao) Statuses() ([]model.TaskStatus, error) {
	return model.TASK_STATUSES, nil
}

func (d *TaskDao) Search(condition model.TaskSearchCondition) ([]model.Task, error) {
	return []model.Task{}, nil
}

func (d *TaskDao) Create(task model.Task) (model.Task, error) {
	err := d.setTaskJSON()
	if err != nil {
		return model.Task{}, err
	}

	task.Id = d.taskJSON.getMaxTaskId() + 1

	d.taskJSON.Tasks = append(d.taskJSON.Tasks, task)
	json, err := json.Marshal(d.taskJSON)
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
	err := d.setTaskJSON()
	if err != nil {
		return model.Task{}, err
	}

	targetTask, ok := lo.Find(d.taskJSON.Tasks, func(t model.Task) bool {
		return t.Id == task.Id
	})
	if !ok {
		return model.Task{}, errors.New("task not found.")
	}

	targetTask.Title = task.Title
	targetTask.Content = task.Content
	targetTask.Status = task.Status
	targetTask.UpdatedAt = time.Now()

	d.taskJSON.Tasks = lo.Map(d.taskJSON.Tasks, func(t model.Task, _ int) model.Task {
		if t.Id == task.Id {
			return targetTask
		}
		return t
	})

	json, err := json.Marshal(d.taskJSON)
	if err != nil {
		return model.Task{}, err
	}
	err = d.truncateAndWrite(json)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (d *TaskDao) Delete(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}

func (d *TaskDao) setTaskJSON() error {
	if d.taskJSON.Tasks != nil {
		return nil
	}
	var taskJSON TaskJSON
	buf, err := io.ReadAll(d.file)
	err = json.Unmarshal(buf, &taskJSON)
	if err != nil {
		return err
	}
	d.taskJSON = taskJSON
	return nil
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
