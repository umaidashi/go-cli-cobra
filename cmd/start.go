package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/dao"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/json"
	"github.com/umaidashi/go-cli-cobra/app/usecase"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start task",
	Long:  `start task`,
	Run: func(cmd *cobra.Command, args []string) {
		json, err := json.NewJSON()
		cobra.CheckErr(err)
		defer json.Close()

		taskRepository := dao.NewTaskDao(json)
		taskUsecase, err := usecase.NewTaskUsecase(taskRepository)
		cobra.CheckErr(err)

		tasks, err := taskUsecase.StartTask(id)
		cobra.CheckErr(err)

		fmt.Println(tasks)
	},
}

var id int

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntVarP(&id, "id", "i", 0, "Help message for toggle")
	startCmd.MarkFlagRequired("id")
}
