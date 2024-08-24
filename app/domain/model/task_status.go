package model

import (
	"encoding/json"

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

func (t *TaskStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = TaskStatusOf(name)
	return nil
}

func (t TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name)
}

var taskStatusMap = lo.KeyBy(TASK_STATUSES, func(t TaskStatus) string {
	return t.Name
})

var TaskStatusDefault = TaskStatusTodo

func TaskStatusOf(name string) TaskStatus {
	v, ok := taskStatusMap[name]
	if ok {
		return v
	}
	panic("status is invalid. name = " + name)
}
