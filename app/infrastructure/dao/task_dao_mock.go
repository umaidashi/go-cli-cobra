package dao

import (
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
)

type TaskDaoMock struct{}

func NewTaskDaoMock() repository.TaskRepository {
	return &TaskDaoMock{}
}

func (d *TaskDaoMock) One(id int) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: "content1",
		Status:  model.TaskStatusDefault,
	}, nil
}

func (d *TaskDaoMock) List() ([]model.Task, error) {
	return []model.Task{
		{
			Id:      1,
			Title:   "title1",
			Content: "content1",
			Status:  model.TaskStatusDefault,
		},
		{
			Id:      2,
			Title:   "title2",
			Content: "content2",
			Status:  model.TaskStatusComplete,
		},
	}, nil
}

func (d *TaskDaoMock) Search(condition model.TaskSearchCondition) ([]model.Task, error) {
	return []model.Task{
		{
			Id:      1,
			Title:   "title1",
			Content: "content1",
			Status:  model.TaskStatusDefault,
		},
		{
			Id:      2,
			Title:   "title2",
			Content: "content2",
			Status:  model.TaskStatusComplete,
		},
	}, nil
}

func (d *TaskDaoMock) Create(task model.Task) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: "content1",
		Status:  model.TaskStatusDefault,
	}, nil
}

func (d *TaskDaoMock) Update(task model.Task) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: "content1",
		Status:  model.TaskStatusDefault,
	}, nil
}

func (d *TaskDaoMock) Delete(task model.Task) (model.Task, error) {
	return model.Task{
		Id:      1,
		Title:   "title1",
		Content: "content1",
		Status:  model.TaskStatusDefault,
	}, nil
}
