package commands

import (
	"com.michael-petri/todo/application/usecase"
	"com.michael-petri/todo/domain/model"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewCreateTodoCommand(create *usecase.CreateTodoCase) *cobra.Command {
	return &cobra.Command{
		Use:   "create <description>",
		Short: "Creates a new todo",
		Long:  `Creates a new todo with the given description.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			todo := model.NewTodo(
				args[0],
			)

			_, err := create.Invoke(todo)

			if err != nil {
				fmt.Printf("Failed to create todo: %s", err.Error())
				os.Exit(1)

			}

			fmt.Printf("Created Todo %d: %s\n", todo.Id.Value, todo.Description)
		},
	}
}
