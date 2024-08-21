package model

import (
	"errors"
	"time"
)

type Task struct {
	Id          int
	Title       string
	Content     string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt time.Time
}

type TaskSearchCondition struct {
	Title  string
	Status TaskStatus
}

func NewTask(title string, content string, statusName *string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("title is required.")
	}

	var taskStatus TaskStatus
	if statusName == nil {
		taskStatus = TaskStatusTodo
	} else {
		t, ok := taskStatusMap[*statusName]
		if !ok {
			return Task{}, errors.New("status is invalid.")
		}
		taskStatus = t
	}

	return Task{
		Title:   title,
		Content: content,
		Status:  taskStatus,
	}, nil
}

func (t Task) StartTask() (Task, error) {
	if t.Status.Name != *TaskStatusProgress.BeforeStatusName {
		return Task{}, errors.New("Cannot start except Todo status.")
	}

	startedTask := t
	startedTask.Status = TaskStatusProgress

	return startedTask, nil
}
