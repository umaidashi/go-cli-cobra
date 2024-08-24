package dao

import (
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
	strings "github.com/umaidashi/go-cli-cobra/app/utils"
)

type TaskDaoMock struct{}

func NewTaskDaoMock() repository.TaskRepository {
	return &TaskDaoMock{}
}

func (d *TaskDaoMock) One(id int) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: strings.EmptyToNil("content1"),
		Status:  model.TaskStatusDefault,
	}, nil
}

func (d *TaskDaoMock) List() ([]model.Task, error) {
	return []model.Task{
		{
			Id:      1,
			Title:   "title1",
			Content: strings.EmptyToNil("content1"),
			Status:  model.TaskStatusDefault,
		},
		{
			Id:      2,
			Title:   "title2",
			Content: strings.EmptyToNil("content1"),
			Status:  model.TaskStatusComplete,
		},
	}, nil
}

func (d *TaskDaoMock) Statuses() ([]model.TaskStatus, error) {
	return model.TASK_STATUSES, nil
}

func (d *TaskDaoMock) Search(condition model.TaskSearchCondition) ([]model.Task, error) {
	return []model.Task{
		{
			Id:      1,
			Title:   "title1",
			Content: strings.EmptyToNil("content1"),
			Status:  model.TaskStatusDefault,
		},
		{
			Id:      2,
			Title:   "title2",
			Content: strings.EmptyToNil("content1"),
			Status:  model.TaskStatusComplete,
		},
	}, nil
}

func (d *TaskDaoMock) Create(task model.Task) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: strings.EmptyToNil("content1"),
		Status:  model.TaskStatusDefault,
	}, nil
}

func (d *TaskDaoMock) Update(task model.Task) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: strings.EmptyToNil("content1"),
		Status:  model.TaskStatusDefault,
	}, nil
}

func (d *TaskDaoMock) Delete(task model.Task) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: strings.EmptyToNil("content1"),
		Status:  model.TaskStatusDefault,
	}, nil
}
