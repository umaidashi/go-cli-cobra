package json

import (
	"io"
	"os"

	goJson "encoding/json"

	"github.com/samber/lo"
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
)

type JSON struct {
	Tasks []model.Task `json:"tasks"`
	file  *os.File
}

func NewJSON() (JSON, error) {
	file, err := os.OpenFile("/tmp/tasks.json", os.O_RDWR, 0666)
	if err != nil {
		return JSON{}, err
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		return JSON{}, err
	}
	var json JSON
	err = goJson.Unmarshal(buf, &json)
	if err != nil {
		return JSON{}, err
	}
	json.file = file
	return json, nil
}

func (j JSON) GetMaxTaskId() int {
	taskIds := lo.Map(j.Tasks, func(t model.Task, _ int) int {
		return t.Id
	})
	maxId := lo.MaxBy(taskIds, func(id int, max int) bool {
		return id > max
	})
	return maxId
}

func (j JSON) Write(buf []byte) error {
	err := j.file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = j.file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = j.file.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func (j JSON) Close() {
	j.file.Close()
}
