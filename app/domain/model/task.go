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

	now := time.Now()

	return Task{
		Title:     title,
		Content:   content,
		Status:    taskStatus,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
