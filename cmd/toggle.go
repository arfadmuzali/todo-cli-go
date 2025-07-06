/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"todo/internal/storage"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle [id1] [id2] ...",
	Args:  cobra.MinimumNArgs(1),
	Short: "Toggle the status of one or more todo items",
	Long: `The toggle command switches the status of one or more todo items between done and not done.

You can pass one or multiple IDs separated by spaces. Each item's status will be flipped.

Examples:
  todoapp toggle 1           # Toggle a single item
  todoapp toggle 2 3 5 8     # Toggle multiple items

This is useful for marking tasks as completed or incomplete quickly, especially in bulk.`,
	Run: func(cmd *cobra.Command, args []string) {
		var failedUpdate []string
		for _, value := range args {
			id, err := strconv.Atoi(value)
			if err != nil {
				failedUpdate = append(failedUpdate, value)
				continue
			}
			if err = storage.UpdateTodosStatus(id); err != nil {
				failedUpdate = append(failedUpdate, value)
				continue
			}
		}
		if len(failedUpdate) != 0 {
			fmt.Printf("Failed to update status IDs %v\n", strings.Join(failedUpdate, ", "))
			return
		}

		fmt.Println("ok")
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
