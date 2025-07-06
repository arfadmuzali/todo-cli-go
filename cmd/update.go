/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	// "strings"
	"todo/internal/storage"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [new description]",
	Args:  cobra.ExactArgs(2),
	Short: "Update the description of an existing todo item",
	Long: `The update command allows you to change the text or details of a specific todo item.

You need to provide the ID of the todo item you want to update, followed by the new content.

Examples:
  todoapp update 3 "Review English grammar notes"
  todoapp update 7 "Go jogging at 6 AM"

This command is useful for correcting or modifying your existing tasks.`,
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to update a todo: Argument ID must be a number"))
			return
		}

		if err = storage.UpdateTodoDescription(id, args[1]); err != nil {
			fmt.Println(fmt.Errorf("Failed to update a todo: %w", err))
			return
		}

		fmt.Println("ok")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
