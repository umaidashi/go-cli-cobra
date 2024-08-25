package dao

import (
	"fmt"
	"testing"

	goJson "encoding/json"

	"github.com/stretchr/testify/suite"
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/json"
)

type TaskDaoSuite struct {
	suite.Suite
	json       json.JSON
	repository repository.TaskRepository
}

func TestTaskDaoSuite(t *testing.T) {
	suite.Run(t, new(TaskDaoSuite))
}

func (s *TaskDaoSuite) SetupTest() {
}

func (s *TaskDaoSuite) BeforeTest(suiteName, testName string) {
	before, _ := json.NewJSON()
	task, _ := model.NewTask("task1", "content1", &model.TaskStatusTodo.Name)
	emptyTask := json.JSON{Tasks: []model.Task{task}}
	emptyTaskJSON, _ := goJson.Marshal(emptyTask)
	before.Write(emptyTaskJSON)
	s.json, _ = json.NewJSON()
	fmt.Println(s.json)
	s.repository = NewTaskDao(s.json)
}

func (s *TaskDaoSuite) AfterTest(suiteName, testName string) {
	s.json.Close()
}

func (s *TaskDaoSuite) TestOne() {
	task, err := s.repository.One(1)
	s.NoError(err)
	s.Equal(1, task.Id)
	s.Equal("task1", task.Title)
	s.Equal("TODO", task.Status.Name)
}

func (s *TaskDaoSuite) TestList() {
	tasks, err := s.repository.List()
	s.NoError(err)
	s.Len(tasks, 1)
	s.Equal(1, tasks[0].Id)
	s.Equal("task1", tasks[0].Title)
	s.Equal("TODO", tasks[0].Status.Name)
}

func (s *TaskDaoSuite) TestStatuses() {
	statuses, err := s.repository.Statuses()
	s.NoError(err)
	s.Len(statuses, len(model.TASK_STATUSES))
}

func (s *TaskDaoSuite) TestCreate() {
	newTask, _ := model.NewTask("task2", "content2", &model.TaskStatusTodo.Name)
	task, err := s.repository.Create(newTask)
	s.NoError(err)
	s.Equal(2, task.Id)
	s.Equal("task2", task.Title)
	s.Equal(model.TaskStatusTodo, task.Status)
}

func (s *TaskDaoSuite) TestUpdate() {
	task, _ := s.repository.One(1)
	start, _ := task.StartTask()
	updated, err := s.repository.Update(start)
	s.NoError(err)
	s.Equal(1, updated.Id)
	s.Equal("task2", updated.Title)
	s.Equal(model.TaskStatusProgress, updated.Status)
}
