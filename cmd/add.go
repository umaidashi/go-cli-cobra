package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/dao"
	"github.com/umaidashi/go-cli-cobra/app/usecase"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new task.",
	Long:  `add new task.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("/tmp/tasks.json", os.O_RDWR, 0666)
		cobra.CheckErr(err)
		defer file.Close()

		taskRepository := dao.NewTaskDao(file)
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
}
