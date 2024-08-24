package repository

import "github.com/umaidashi/go-cli-cobra/app/domain/model"

type TaskRepository interface {
	One(id int) (model.Task, error)
	List() ([]model.Task, error)
	Statuses() ([]model.TaskStatus, error)
	Search(condition model.TaskSearchCondition) ([]model.Task, error)
	Create(task model.Task) (model.Task, error)
	Update(task model.Task) (model.Task, error)
	Delete(task model.Task) (model.Task, error)
}
