package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/dao"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/json"
	"github.com/umaidashi/go-cli-cobra/app/usecase"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new task.",
	Long:  `add new task.`,
	Run: func(cmd *cobra.Command, args []string) {
		json, err := json.NewJSON()
		cobra.CheckErr(err)
		defer json.Close()

		taskRepository := dao.NewTaskDao(json)
		taskUsecase, err := usecase.NewTaskUsecase(taskRepository)
		cobra.CheckErr(err)

		tasks, err := taskUsecase.AddTask(title, content)
		cobra.CheckErr(err)

		fmt.Println(tasks)
	},
}

var title, content string

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&title, "title", "t", "", "title of task")
	addCmd.Flags().StringVarP(&content, "content", "c", "", "content of task")

	addCmd.MarkFlagRequired("title")
	addCmd.MarkFlagRequired("content")
}
