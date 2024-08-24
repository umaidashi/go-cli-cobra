package cmd

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/dao"
	"github.com/umaidashi/go-cli-cobra/app/usecase"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of tasks.",
	Long:  `display list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("/tmp/tasks.json", os.O_RDONLY, 0666)
		cobra.CheckErr(err)
		defer file.Close()

		taskRepository := dao.NewTaskDao(file)
		taskUsecase, err := usecase.NewTaskUsecase(taskRepository)
		cobra.CheckErr(err)

		tasks, err := taskUsecase.List()
		cobra.CheckErr(err)

		for _, task := range tasks {
			fmt.Println(task.Id, task.Title, task.Status.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
