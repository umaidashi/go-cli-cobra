package model

import (
	"errors"
	"time"
)

type Task struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Content     *string    `json:"content"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at"`
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
		Content:   &content,
		Status:    taskStatus,
		CreatedAt: now,
	}, nil
}

func (t Task) StartTask() (Task, error) {
	if t.Status.Name != *TaskStatusProgress.BeforeStatusName {
		return Task{}, errors.New("Cannot start except Todo status.")
	}

	startedTask := t
	startedTask.Status = TaskStatusProgress
	startedTask.UpdatedAt = time.Now()

	return startedTask, nil
}

func (t Task) DoneTask() (Task, error) {
	if t.Status.Name != *TaskStatusComplete.BeforeStatusName {
		return Task{}, errors.New("Cannot done except Progress status.")
	}

	doneTask := t
	doneTask.Status = TaskStatusComplete
	doneTask.CompletedAt = &time.Time{}

	return doneTask, nil
}
