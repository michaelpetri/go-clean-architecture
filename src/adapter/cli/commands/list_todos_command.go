package commands

import (
	"com.michael-petri.todo/application/usecase"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewListTodosCommand(create *usecase.ListTodosCase) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists all todos",
		Long:  `Lists all todos with their id`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := create.Invoke()

			if err != nil {
				fmt.Printf("Failed to load todos: %s", err.Error())
				os.Exit(1)
			}

			if len(todos) == 0 {
				fmt.Println("No todos found")
			} else {
				fmt.Printf("%-5s %-10s\n", "ID", "Description")
				for _, todo := range todos {
					fmt.Printf("%-5d %-10s\n", todo.Id.Value, todo.Description)
				}
			}
		},
	}
}
