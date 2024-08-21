package model

import (
	"database/sql/driver"

	"github.com/samber/lo"
)

type TaskStatus struct {
	Label            string
	Name             string
	ColorFg          HexColor
	ColorBg          HexColor
	BeforeStatusName *string
}

var TaskStatusPending = TaskStatus{
	Label:            "保留中",
	Name:             "PENDING",
	ColorFg:          "#ffffff",
	ColorBg:          "#000000",
	BeforeStatusName: &TaskStatusTodo.Name,
}

var TaskStatusTodo = TaskStatus{
	Label:            "未着手",
	Name:             "TODO",
	ColorFg:          "#ffffff",
	ColorBg:          "#000000",
	BeforeStatusName: nil,
}

var TaskStatusProgress = TaskStatus{
	Label:            "進行中",
	Name:             "PROGRESS",
	ColorFg:          "#ffffff",
	ColorBg:          "#000000",
	BeforeStatusName: &TaskStatusTodo.Name,
}

var TaskStatusComplete = TaskStatus{
	Label:            "完了",
	Name:             "COMPLETE",
	ColorFg:          "#ffffff",
	ColorBg:          "#000000",
	BeforeStatusName: &TaskStatusProgress.Name,
}

var TASK_STATUSES = []TaskStatus{
	TaskStatusPending,
	TaskStatusTodo,
	TaskStatusProgress,
	TaskStatusComplete,
}

var taskStatusMap = lo.KeyBy(TASK_STATUSES, func(t TaskStatus) string {
	return t.Label
})

var TaskStatusDefault = TaskStatusTodo

func (t TaskStatus) Value() (driver.Value, error) {
	return t.Name, nil
}

func (t *TaskStatus) Scan(input interface{}) error {
	*t = TaskStatusOf(input.(string))
	return nil
}

func TaskStatusOf(name string) TaskStatus {
	v, ok := taskStatusMap[name]
	if ok {
		return v
	}
	panic("status is invalid. name = " + name)
}
