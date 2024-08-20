/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/fancurson/CLI-Shell/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mars task as completed",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("argument \"%s\" is invalid\n", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ViewTasks()
		if err != nil {
			fmt.Println(err)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid id:", id)
				continue
			}

			err := db.DeleteTask(tasks[id-1].Key)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Mark %d as completed", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
