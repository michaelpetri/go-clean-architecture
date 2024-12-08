package commands

import (
	"bufio"
	"com.michael-petri/todo/application/usecase"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewInteractiveCommand(
	createCase *usecase.CreateTodoCase,
	listCase *usecase.ListTodosCase,
	resolveCase *usecase.ResolveTodoCase,
) *cobra.Command {
	return &cobra.Command{
		Use:   "interactive",
		Short: "Interactive session",
		Long:  `Starts an interactive session to manage todos`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)

			createCommand := NewCreateTodoCommand(createCase)
			listCommand := NewListTodosCommand(listCase)
			resolveCommand := NewResolveTodoCommand(resolveCase)

			for {
				fmt.Print("Command? (create|list|resolve|exit): ")
				if !scanner.Scan() {
					break
				}
				input := scanner.Text()

				switch input {
				case "create":
					fmt.Print("Description: ")
					if !scanner.Scan() {
						break
					}
					description := scanner.Text()

					createCommand.Run(cmd, []string{description})
					listCommand.Run(cmd, []string{})
					break
				case "resolve":
					fmt.Print("ID: ")
					if !scanner.Scan() {
						break
					}
					id := scanner.Text()

					resolveCommand.Run(cmd, []string{id})
					listCommand.Run(cmd, []string{})
					break
				case "list":
					listCommand.Run(cmd, []string{})
					break
				case "exit":
					fmt.Println("Bye!")
					os.Exit(0)
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Faied to read input: %v\n", err)
			}
		},
	}
}
