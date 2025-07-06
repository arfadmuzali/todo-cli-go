/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"todo/internal/storage"

	"github.com/spf13/cobra"
)

var All bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a todo item or all items",
	Long: `Delete command allows you to remove one or more todo items.
You can delete a specific item by ID, or all items at once.
Examples:
  app delete 1      # Delete item with ID 1
  app delete --all  # Delete all items`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if All {
			if err = storage.DeleteAll(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("ok")
			}
			return
		}

		if len(args) == 0 {
			fmt.Println(fmt.Errorf("Failed to delete a todo: Argument ID not found"))
			return

		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to delete a todo: Argument ID must be a number"))
			return
		}

		if err = storage.DeleteTodo(id); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVarP(&All, "all", "a", false, "Delete all todos")
}
