/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/arfadmuzali/todo-cli-go/internal/storage"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item to your list",
	Args:  cobra.ExactArgs(1),
	Long: `The add command lets you create a new todo item and save it to your list.

Usage examples:
  todoapp add "Learn English vocabulary"
  todoapp add "Team meeting at 10 AM"

You can add any task or reminder. Once added, it will appear in your todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := storage.AddTodo(args[0]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
