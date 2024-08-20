/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your schedule",

	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Println(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
