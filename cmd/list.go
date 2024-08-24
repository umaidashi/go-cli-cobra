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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		buf, err := os.ReadFile("/tmp/tasks.json")
		cobra.CheckErr(err)

		taskRepository := dao.NewTaskDao(buf)
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
