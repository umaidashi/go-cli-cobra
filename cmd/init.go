package cmd

import (
	goJson "encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/umaidashi/go-cli-cobra/app/domain/model"
	"github.com/umaidashi/go-cli-cobra/app/infrastructure/json"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "task app init",
	Long:  `task app init`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Create("/tmp/tasks.json")
		cobra.CheckErr(err)
		defer file.Close()

		emptyTask := json.JSON{Tasks: []model.Task{}}
		emptyTaskJSON, err := goJson.Marshal(emptyTask)
		cobra.CheckErr(err)
		file.Write(emptyTaskJSON)

		fmt.Println("init task app")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
