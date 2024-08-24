package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/dao"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/json"
	"github.com/umaidashi/go-cli-cobra/app/usecase"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "task done",
	Long:  `task done`,
	Run: func(cmd *cobra.Command, args []string) {
		json, err := json.NewJSON()
		cobra.CheckErr(err)
		defer json.Close()

		taskRepository := dao.NewTaskDao(json)
		taskUsecase, err := usecase.NewTaskUsecase(taskRepository)
		cobra.CheckErr(err)

		tasks, err := taskUsecase.DoneTask(id)
		cobra.CheckErr(err)

		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	doneCmd.Flags().IntVarP(&id, "id", "i", 0, "Help message for toggle")
	doneCmd.MarkFlagRequired("id")
}
