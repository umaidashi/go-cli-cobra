/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/dao"
	"github.com/umaidashi/go-cli-cobra/app/usecase"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start task",
	Long:  `start task`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("/tmp/tasks.json", os.O_RDWR, 0666)
		cobra.CheckErr(err)
		defer file.Close()

		taskRepository := dao.NewTaskDao(file)
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
}
