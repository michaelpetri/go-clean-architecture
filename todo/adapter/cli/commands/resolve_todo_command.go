package commands

import (
	"com.michael-petri/todo/application/usecase"
	"com.michael-petri/todo/domain/value"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func NewResolveTodoCommand(create *usecase.ResolveTodoCase) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve <id>",
		Short: "Resolves a todo",
		Long:  `Resolves a todo by its id`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			idValue, parseErr := strconv.ParseUint(args[0], 10, 64)

			if parseErr != nil {
				fmt.Printf("Failed to parse todo id: %s", parseErr.Error())
			}

			id := value.NewTodoId(idValue)

			err := create.Invoke(id)

			if err != nil {
				fmt.Printf("Failed to resolve todo %d: %s", id.Value, err.Error())
				os.Exit(1)
			}

			fmt.Printf("Todo %d resolved\n", id.Value)
		},
	}
}
