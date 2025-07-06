package cmd

import (
	"fmt"
	"os"
	"todo/internal/storage"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all your todo items",
	Long: `The list command shows all the todo items currently saved.

It displays each item with its ID, content, and possibly status (e.g., done or not).

Examples:
  todoapp list         # Show all todos
  todoapp list --done  # Show only completed items (if filtering is supported)

This is helpful for reviewing your tasks and tracking progress.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := storage.LoadTodos()

		if err != nil {
			fmt.Println(fmt.Errorf("Failed to load todos : %w", err))
			return
		}

		table := tablewriter.NewTable(os.Stdout)
		table.Header("ID", "Description", "Status", "Created At")
		for _, value := range todos {
			table.Append(value.Id, value.Description, value.Status, value.CreatedAt.Format("02 Jan 15:04"))
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
