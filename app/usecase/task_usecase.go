package usecase

import (
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
)

type TaskUsecase struct {
	repository repository.TaskRepository
}

func NewTaskUsecase(repository repository.TaskRepository) (TaskUsecase, error) {
	return TaskUsecase{
		repository: repository,
	}, nil
}

func (u TaskUsecase) One(id int) (model.Task, error) {
	task, err := u.repository.One(id)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (u TaskUsecase) List() ([]model.Task, error) {
	tasks, err := u.repository.List()
	if err != nil {
		return []model.Task{}, err
	}

	return tasks, nil
}

func (u TaskUsecase) AddTask(title string, content string) (model.Task, error) {
	newTask, err := model.NewTask(
		title,
		content,
		&model.TaskStatusDefault.Name,
	)
	if err != nil {
		return model.Task{}, err
	}
	createdTask, err := u.repository.Create(newTask)
	if err != nil {
		return model.Task{}, err
	}

	return createdTask, nil
}

func (u TaskUsecase) StartTask(id int) (model.Task, error) {
	task, err := u.repository.One(id)
	if err != nil {
		return model.Task{}, err
	}
	startedTask, err := task.StartTask()
	if err != nil {
		return model.Task{}, err
	}
	updatedTask, err := u.repository.Update(startedTask)
	if err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}
