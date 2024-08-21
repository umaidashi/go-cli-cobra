package model

import (
	"database/sql/driver"

	"github.com/samber/lo"
)

type TaskStatus struct {
	Label   string
	Name    string
	ColorFg HexColor
	ColorBg HexColor
}

var TaskStatusPending = TaskStatus{
	Label: "保留中",
	Name:  "PENDING",
}

var TaskStatusTodo = TaskStatus{
	Label: "未着手",
	Name:  "TODO",
}

var TaskStatusProcessing = TaskStatus{
	Label: "進行中",
	Name:  "PROCESSING",
}

var TaskStatusComplete = TaskStatus{
	Label: "完了",
	Name:  "COMPLETE",
}

var TASK_STATUSES = []TaskStatus{
	TaskStatusPending,
	TaskStatusTodo,
	TaskStatusProcessing,
	TaskStatusComplete,
}

var taskStatusMap = lo.KeyBy(TASK_STATUSES, func(t TaskStatus) string {
	return t.Label
})

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
