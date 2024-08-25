package dao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/domain/repository"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/json"
	strings "github.com/umaidashi/go-cli-cobra/app/utils"
)

type TaskDaoSuite struct {
	suite.Suite
	repository repository.TaskRepository
}

func TestTaskDaoSuite(t *testing.T) {
	suite.Run(t, new(TaskDaoSuite))
}

func (s *TaskDaoSuite) SetupTest() {
}

func (s *TaskDaoSuite) BeforeTest(suiteName, testName string) {
	j, _ := json.NewJSON()
	taskJson := json.JSON{Tasks: []model.Task{
		{Id: 1, Title: "task1", Content: strings.EmptyToNil("content1"), Status: model.TaskStatusTodo},
	}}
	j.Tasks = taskJson.Tasks
	fmt.Printf("j: %v\n", j)
	s.repository = NewTaskDao(j)
}

func (s *TaskDaoSuite) AfterTest(suiteName, testName string) {
}

func (s *TaskDaoSuite) TestOne() {
	task, err := s.repository.One(1)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, task.Id)
	assert.Equal(s.T(), "task1", task.Title)
	assert.Equal(s.T(), "content1", *task.Content)
	assert.Equal(s.T(), "TODO", task.Status.Name)
}

func (s *TaskDaoSuite) TestList() {
	tasks, err := s.repository.List()
	assert.Nil(s.T(), err)
	assert.Len(s.T(), tasks, 1)
	assert.Equal(s.T(), 1, tasks[0].Id)
	assert.Equal(s.T(), "task1", tasks[0].Title)
	assert.Equal(s.T(), "content1", *tasks[0].Content)
	assert.Equal(s.T(), "TODO", tasks[0].Status.Name)
}

func (s *TaskDaoSuite) TestStatuses() {
	statuses, err := s.repository.Statuses()
	assert.Nil(s.T(), err)
	assert.Len(s.T(), statuses, len(model.TASK_STATUSES))
}

func (s *TaskDaoSuite) TestCreate() {
	newTask, _ := model.NewTask("task2", "content2", &model.TaskStatusTodo.Name)
	task, err := s.repository.Create(newTask)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, task.Id)
	assert.Equal(s.T(), "task2", task.Title)
	assert.Equal(s.T(), "content2", *task.Content)
	assert.Equal(s.T(), "TODO", task.Status.Name)
}

func (s *TaskDaoSuite) TestUpdate() {
	task, _ := s.repository.One(1)
	start, _ := task.StartTask()
	updated, err := s.repository.Update(start)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, updated.Id)
	assert.Equal(s.T(), "task1", updated.Title)
	assert.Equal(s.T(), "content1", *updated.Content)
	assert.Equal(s.T(), "PROGRESS", updated.Status.Name)
}
